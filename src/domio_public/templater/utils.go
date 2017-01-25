package templater

import (
    "html/template"
    //"bytes"
    //"log"
    //"net/http"
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

const LoginPageTemplate = "{{.LoginForm}}"
const SignupPageTemplate = "{{.SignupForm}}"

func getTemplateAsString(templateName string, title string) string {

    //if (templateName == "IndexPage") {
    //
    //
    //} else if (templateName == "LoginPage") {
    //    pageData := BasePageData{
    //        PageName: templateName,
    //        Title: title,
    //        Body: template.HTML(LoginPageTemplate),
    //    }
    //
    //    baseTemplate, err := t.Parse(GetBaseTemplateAsString())
    //
    //    if (err != nil) {
    //        log.Print(err)
    //    }
    //
    //    var doc bytes.Buffer
    //    baseTemplate.Execute(&doc, pageData)
    //    return doc.String()
    //} else if (templateName == "SignupPage") {
    //    pageData := BasePageData{
    //        PageName: templateName,
    //        Title: title,
    //        Body: template.HTML(SignupPageTemplate),
    //    }
    //
    //    baseTemplate, err := t.Parse(GetBaseTemplateAsString())
    //
    //    if (err != nil) {
    //        log.Print(err)
    //    }
    //
    //    var doc bytes.Buffer
    //    baseTemplate.Execute(&doc, pageData)
    //    return doc.String()
    //}

    return "error"
}
func WritePage(w http.ResponseWriter, template *template.Template, data interface{}) {
    w.Header().Set("Content-Type", "text/html")
    template.ExecuteTemplate(w, "BaseTemplate", data)
}