package index_handler

import (
    "domio_public/components/api"
    "domio_public/templater"
    "html/template"
    "log"
    "net/http"
)

type PageData struct {
    Title         string
    UserName      string
    TopBar        template.HTML
    AppStatusInfo api.AppStatusInfo
}

var parsedTemplate *template.Template

func init() {
    var err error
    parsedTemplate, err = templater.GetParsedTemplate("IndexPage", "Domio")

    if err != nil {
        log.Print(err)
    }
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
    appStatusInfo := api.GetAPIStatus()

    w.Header().Set("Content-Type", "text/html")

    pageData := PageData{
        Title:    "Domio",
        UserName: "Domio",
        TopBar:   templater.GetTopBar(),
        AppStatusInfo: appStatusInfo,
    }

    templater.WritePage(w, parsedTemplate, pageData)
}
