package index_handler

import (
    "net/http"
    "html/template"
    "domio_public/templater"
    "log"
)

type PageData struct {
    Title          string
    TopBar template.HTML
    UserName       string
}

var parsedTemplate *template.Template

func init() {
    var err error
    parsedTemplate, err = templater.GetParsedTemplate("IndexPage", "Domio")

    if (err != nil) {
        log.Print(err)
    }
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {

    w.Header().Set("Content-Type", "text/html")

    pageData := PageData{
        Title: "Domio",
        UserName: "Domio",
        TopBar: templater.GetHeader(),
    }

    templater.WritePage(w, parsedTemplate, pageData)
}