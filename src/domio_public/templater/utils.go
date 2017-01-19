package templater

import (
    "html/template"
    "bytes"
    "log"
)

type BasePageData struct {
    Body template.HTML
}

func GetParsedTemplate(templateName string) (*template.Template, error) {
    t := template.New(templateName)

    parsed, err := t.Parse(getTemplateAsString(templateName))
    if (err != nil) {
        log.Print(err)
    }
    return parsed, err
}

func getTemplateAsString(templateName string) string {
    if (templateName == "Index") {
        t := template.New("temp_tmpl")
        pageData := BasePageData{
            Body: template.HTML("{{.LinksContainer}} <br> {{.UserName}}"),
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
                    <title></title>
                    <link rel="stylesheet" type="text/css" href="/style.css">
                    <script src="/client.public.bundle.js"></script>
                </head>
                <body>
                    {{.Body}}
                </body>
                </html>
    `
}