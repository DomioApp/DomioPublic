package templater

import (
    "html/template"
    "bytes"
    "log"
    "net/http"
)

type Link struct {
    Url       string
    Label     string
    ClassName string
}

type BasePageData struct {
    Title string
    Body  template.HTML
}

const IndexPageTemplate = `
                            {{.TopBar}}
                                <div class="b-app-status-info">
                                    {{with .AppStatusInfo}}
                                        <div>Version: <span>{{.Version}}</span></div>
                                        <div>Built: <span>{{.BuildAgo}} ago</span></div>
                                        <div>Buildstamp: <span>{{.Buildstamp}}</span></div>
                                        <div>BuildDateTime: <span>{{.BuildDateTime}}</span></div>
                                        <div>Hash: <span>{{.Hash}}</span></div>
                                    {{end}}
                                </div>
                            `

const LoginPageTemplate = "{{.LoginForm}}"

func GetParsedTemplate(templateName string, title string) (*template.Template, error) {
    t := template.New(templateName)

    parsed, err := t.Parse(getTemplateAsString(templateName, title))
    if (err != nil) {
        log.Print(err)
    }
    return parsed, err
}

func getTemplateAsString(templateName string, title string) string {
    t := template.New("temp_tmpl")

    if (templateName == "IndexPage") {
        pageData := BasePageData{
            Title:title,
            Body: template.HTML(IndexPageTemplate),
        }

        baseTemplate, err := t.Parse(getBaseTemplate())

        if (err != nil) {
            log.Print(err)
        }

        var doc bytes.Buffer
        baseTemplate.Execute(&doc, pageData)
        return doc.String()
    } else if (templateName == "LoginPage") {
        pageData := BasePageData{
            Title: title,
            Body: template.HTML(LoginPageTemplate),
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
                    </head>

                    <body>
                        {{.Body}}
                    </body>
                    <script src="/bundle.js"></script>
                </html>
    `
}

func WritePage(w http.ResponseWriter, parsedTemplate *template.Template, pageData interface{}) {
    err := parsedTemplate.Execute(w, pageData)
    if (err != nil) {
        log.Print(err.Error())
    }

}