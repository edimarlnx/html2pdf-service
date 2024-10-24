package render

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/edimarlnx/html2pdf-service/internal/utils"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

var (
	DebugMode = utils.GetEnv("HTML_2_PDF_DEBUG", "") != ""
)

func PDFFromContent(content []byte, waitForSelector string) ([]byte, error) {
	if len(content) == 0 {
		return nil, errors.New("no content to process")
	}

	tmpFile := fmt.Sprintf("%s/%s.html", os.TempDir(), uuid.New().String())
	if err := os.WriteFile(tmpFile, content, 0o644); err != nil {
		return nil, err
	}
	defer func() {
		err := os.Remove(tmpFile)
		if err != nil {
			fmt.Println("Error on remove tmp file", err)
		} else {
			fmt.Printf("Tmp file [%s] removed\n", tmpFile)
		}
	}()

	return PDF(fmt.Sprintf("file://%s", tmpFile), nil, waitForSelector)
}

func PDF(url string, headers map[string]interface{}, waitForSelector string) ([]byte, error) {
	fmt.Printf("Create PDF for %s\n", url)
	var opts []chromedp.ContextOption
	if DebugMode {
		opts = append(opts, chromedp.WithDebugf(log.Printf))
	}
	ctx, cancel := chromedp.NewContext(context.Background(), opts...)
	defer cancel()

	var pdfBuffer []byte
	var tasks []chromedp.Action
	if headers != nil && len(headers) > 0 {
		tasks = append(tasks, network.Enable())
		tasks = append(tasks, network.SetExtraHTTPHeaders(headers))
	}
	tasks = append(tasks, chromedp.Navigate(url))
	tasks = append(tasks, chromedp.WaitVisible(`body`, chromedp.ByQuery))
	if waitForSelector != "" {
		tasks = append(tasks, chromedp.WaitVisible(waitForSelector, chromedp.ByQuery))

	}
	tasks = append(tasks, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		pdfBuffer, _, err = page.PrintToPDF().
			WithPrintBackground(true).
			WithDisplayHeaderFooter(false).
			Do(ctx)
		return err
	}))

	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 60*time.Second)
	defer timeoutCancel()
	err := chromedp.Run(timeoutCtx, tasks...)

	if err != nil {
		fmt.Println("Error on generate PDF", err)
		return nil, err
	}

	return pdfBuffer, nil
}
