package index_handler

import (
    "net/http"
    "html/template"
    "domio_public/templater"
    "log"
)

type pageData struct {
    Title          string
    LinksContainer template.HTML
    UserName       string
}

var parsedTemplate *template.Template

func init() {
    var err error
    parsedTemplate, err = templater.GetParsedTemplate("Index")

    if (err != nil) {
        log.Print(err)
    }
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {

    w.Header().Set("Content-Type", "text/html")

    header := templater.GetHeader()

    p := pageData{
        Title: "Domio",
        UserName: "Astaxie",
        LinksContainer:header,
    }

    parsedTemplate.Execute(w, p)
}