package domain_info_page_handler

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
    DomainInfo *api.DomainJson
}

var domainInfoPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    domainInfoPageTemplate = templater.BuildTemplate(GetDomainInfoPageTemplate)
}

func DomainInfoPageHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    domainName := vars["domainName"]

    domainInfo, domainInfoError := api.GetDomainInfo(domainName, "")

    if (domainInfoError != nil) {
        log.Print(domainInfoError)
    }

    templater.WriteTemplate(w, req, domainInfoPageTemplate, GetPageName(), GetPageData(domainName, domainInfo))
}

func GetUrl() string {
    return "/domain/{domainName}"
}

func GetPageName() string {
    return "DomainInfoPage"
}

func GetPageData(domainName string, domainInfo *api.DomainJson) PageData {
    pageData := PageData{
        PageTitle: "Domio - Domain information " + domainName,
        DomainInfo: domainInfo,
    }

    return pageData
}
