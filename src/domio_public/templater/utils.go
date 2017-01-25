package templater

import (
    "html/template"
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

func WritePage(w http.ResponseWriter, template *template.Template, data interface{}) {
    w.Header().Set("Content-Type", "text/html")
    template.ExecuteTemplate(w, "BaseTemplate", data)
}