# Create PDF from URL

Send a GET request to `/create-pdf` with the following query parameters:

* `url` - URL to scrape
* `headers` - Headers to send to the URL
* `downloadFileName` - (Optional) Name of the file to download
* `waitForSelector` - (Optional) DOM Selector to wait before create a PDF.

# Create PDF from HTML content

Send a POST request to `/create-pdf` with the following query parameters and a HTML content on body of the request:

* `downloadFileName` - (Optional) Name of the file to download
* `waitForSelector` - (Optional) DOM Selector to wait before create a PDF.

# Authentication

You can use the `x-api-key` header to authenticate your requests or use the query param `api-key`:
