package login_page_handler

import (
    "net/http"
    "html/template"
    "domio_public/templater"
    "log"
)

type Person struct {
    UserName string
}

type LoginPageData struct {
    Title     string
    LoginForm template.HTML
}

var parsedTemplate *template.Template

func init() {
    var err error
    parsedTemplate, err = templater.GetParsedTemplate("LoginPage", "Domio Login")

    if (err != nil) {
        log.Print(err)
    }

}

func LoginPageHandler(w http.ResponseWriter, req *http.Request) {

    w.Header().Set("Content-Type", "text/html")

    //pageData := LoginPageData{
    //    Title: "Domio",
    //    LoginForm:templater.GetLoginForm(),
    //}

    //templater.WritePage(w, parsedTemplate, pageData)
}
