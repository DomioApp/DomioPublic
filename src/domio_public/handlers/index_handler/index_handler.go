package index_handler

import (
    "net/http"
    "html/template"
    "domio_public/templater"
)

type pageData struct {
    LinksContainer template.HTML
    UserName       string
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {

    w.Header().Set("Content-Type", "text/html")

    t := template.New("Index")
    output, _ := t.Parse("{{.LinksContainer}} <br> {{.UserName}}")

    header := templater.GetHeader()
    p := pageData{UserName: "Astaxie", LinksContainer:header}

    output.Execute(w, p)
}