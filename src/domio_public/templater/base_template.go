package templater

import (
    "log"
    "html/template"
)

func GetParsedTemplate(templateName string, title string) (*template.Template, error) {
    t := template.New(templateName)

    parsed, err := t.Parse(GetBaseTemplateAsString())
    if (err != nil) {
        log.Print(err)
    }
    return parsed, err
}

func GetBaseTemplateAsString() string {
    return `
                <!DOCTYPE html>
                <html lang="en">

                    <head>
                        <meta charset="UTF-8">
                        <meta name="page" content="{{.PageName}}">
                        <title>{{.Title}}</title>
                        <link rel="stylesheet" type="text/css" href="/style.css">
                    </head>

                    <body>
                        {{.Body}}
                    </body>
                    <script src="/bundle.js"></script>
                </html>
           `
}
