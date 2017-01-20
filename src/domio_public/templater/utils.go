package templater

import (
    "html/template"
    "bytes"
    "log"
    "net/http"
)

type BasePageData struct {
    Title string
    Body  template.HTML
}

func GetParsedTemplate(templateName string, title string) (*template.Template, error) {
    t := template.New(templateName)

    parsed, err := t.Parse(getTemplateAsString(templateName, title))
    if (err != nil) {
        log.Print(err)
    }
    return parsed, err
}

func getTemplateAsString(templateName string, title string) string {
    if (templateName == "IndexPage") {
        t := template.New("temp_tmpl")

        pageData := BasePageData{
            Title:title,
            Body: template.HTML("{{.TopBar}} <br> {{.UserName}}"),
        }

        baseTemplate, err := t.Parse(getBaseTemplate())

        if (err != nil) {
            log.Print(err)
        }

        var doc bytes.Buffer
        baseTemplate.Execute(&doc, pageData)
        return doc.String()
    }

    return "error"
}

func getBaseTemplate() string {
    return `
                <!DOCTYPE html>
                <html lang="en">
                <head>
                    <meta charset="UTF-8">
                    <title>{{.Title}}</title>
                    <link rel="stylesheet" type="text/css" href="/style.css">
                    <script src="/client.public.bundle.js"></script>
                </head>
                <body>
                    {{.Body}}
                </body>
                </html>
    `
}

func WritePage(w http.ResponseWriter, parsedTemplate *template.Template, pageData interface{}) {
    err := parsedTemplate.Execute(w, pageData)
    if (err != nil) {
        log.Print(err.Error())
    }

}