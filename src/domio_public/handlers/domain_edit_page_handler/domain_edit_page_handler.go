package domain_edit_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
    "github.com/gorilla/mux"
    "log"
)

type PageData struct {
    PageTitle  string
    DomainInfo api.DomainJson
}

var domainEditPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    domainEditPageTemplate = templater.BuildTemplate(GetDomainEditPageTemplate)
}

func DomainEditPageHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    domainName := vars["domainName"]

    var tokenErr error

    tokenCookie, tokenErr = req.Cookie("token")

    log.Print(tokenErr)

    templater.WriteTemplate(w, req, domainEditPageTemplate, GetPageName(), GetPageData(domainName))
}

func GetUrl() string {
    return "/profile/domain/{domainName}"
}

func GetPageName() string {
    return "DomainEditPage"
}

func GetPageData(domainName string) PageData {
    domainInfo := api.GetDomainInfo(domainName, tokenCookie.Value)

    log.Print(domainInfo)

    pageData := PageData{
        PageTitle: "Domio - Edit Domain " + domainName,
        DomainInfo: domainInfo,
    }

    return pageData
}
