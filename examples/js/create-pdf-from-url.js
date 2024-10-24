import fs from 'node:fs'

// const apiUrl = 'http://localhost:8080'
const apiUrl = 'API_URL'
const apiKey = 'API_KEY'

export const createPdfFromUrl = async ({
                                           url,
                                           headers = {
                                               Authorization: `Basic ${btoa('username:password')}`
                                           },
                                           downloadFileName = 'example.pdf',
                                           waitForSelector,
                                       }) => {

    const uri = new URL(`${apiUrl}/create-pdf`)
    uri.searchParams.set('api-key', apiKey)
    uri.searchParams.set('url', url)
    uri.searchParams.set('downloadFileName', downloadFileName)
    uri.searchParams.set('waitForSelector', waitForSelector)
    uri.searchParams.set('headers', JSON.stringify(headers))
    console.log(uri.toString())
    return fetch(uri.toString(), {
        method: 'GET',
    }).then(req => req.blob().then(data => data.arrayBuffer()))
}

export const createPdfFromHtml = async ({
                                            content,
                                            downloadFileName = 'example.pdf',
                                            waitForSelector,
                                        }) => {

    const uri = new URL(`${apiUrl}/create-pdf`)
    uri.searchParams.set('api-key', apiKey)
    uri.searchParams.set('downloadFileName', downloadFileName)
    uri.searchParams.set('waitForSelector', waitForSelector)
    console.log(uri.toString())
    return fetch(uri.toString(), {
        method: 'POST',
        body: content,
    }).then(req => req.blob().then(data => data.arrayBuffer()))
}

createPdfFromUrl({
    url: 'URL_TO_GENERATE_PDF',
    headers: {"Authorization": "Basic BASIC_TOKEN"},
    waitForSelector: '[placeholder~="Filter"]',
})
    .then(pdfData => {
        fs.writeFileSync('example.pdf', pdfData)
    })

const htmlContent = fs.readFileSync('sample.html', 'utf8')
createPdfFromHtml({
    content: htmlContent,
    waitForSelector: '[alt="Sample Image"]',
})
    .then(pdfData => {
        fs.writeFileSync('example-html.pdf', pdfData)
    })