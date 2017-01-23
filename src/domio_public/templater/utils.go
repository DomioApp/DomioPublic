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
    Title    string
    PageName string
    Body     template.HTML
}

type HomePageData struct {
    Title    string
    PageName string
    Header   template.HTML
    Body     template.HTML
}

const IndexPageTemplate = `
                            {{.LeftColumn}}
                            {{.RightColumn}}
                            {{.AppStatusInfoBar}}
                          `
const LoginPageTemplate = "{{.LoginForm}}"
const SignupPageTemplate = "{{.SignupForm}}"

func getTemplateAsString(templateName string, title string) string {
    t := template.New("temp_tmpl")

    if (templateName == "IndexPage") {
        pageData := HomePageData{
            PageName: templateName,
            Title: title,
            Header: template.HTML(`{{.TopBar}}`),
            Body: template.HTML(IndexPageTemplate),
        }

        baseTemplate, err := t.Parse(GetBaseTemplateAsString())

        if (err != nil) {
            log.Print(err)
        }

        var doc bytes.Buffer
        baseTemplate.Execute(&doc, pageData)
        return doc.String()
    } else if (templateName == "LoginPage") {
        pageData := BasePageData{
            PageName: templateName,
            Title: title,
            Body: template.HTML(LoginPageTemplate),
        }

        baseTemplate, err := t.Parse(GetBaseTemplateAsString())

        if (err != nil) {
            log.Print(err)
        }

        var doc bytes.Buffer
        baseTemplate.Execute(&doc, pageData)
        return doc.String()
    } else if (templateName == "SignupPage") {
        pageData := BasePageData{
            PageName: templateName,
            Title: title,
            Body: template.HTML(SignupPageTemplate),
        }

        baseTemplate, err := t.Parse(GetBaseTemplateAsString())

        if (err != nil) {
            log.Print(err)
        }

        var doc bytes.Buffer
        baseTemplate.Execute(&doc, pageData)
        return doc.String()
    }

    return "error"
}


func WritePage(w http.ResponseWriter, parsedTemplate *template.Template, pageData interface{}) {
    err := parsedTemplate.Execute(w, pageData)
    if (err != nil) {
        log.Print(err.Error())
    }

}