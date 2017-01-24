package signup_page_handler

import (
    "net/http"
    "html/template"
    "domio_public/templater"
    "log"
)

type Person struct {
    UserName string
}

type SignupPageData struct {
    Title      string
    SignupForm template.HTML
}

var parsedTemplate *template.Template

func init() {
    var err error
    parsedTemplate, err = templater.GetParsedTemplate("SignupPage", "Domio Signup")

    if (err != nil) {
        log.Print(err)
    }

}

func SignupPageHandler(w http.ResponseWriter, req *http.Request) {

    w.Header().Set("Content-Type", "text/html")

    //pageData := SignupPageData{
    //    Title: "Domio",
    //    SignupForm:templater.GetLoginForm(),
    //}
    //
    //templater.WritePage(w, parsedTemplate, pageData)
}
