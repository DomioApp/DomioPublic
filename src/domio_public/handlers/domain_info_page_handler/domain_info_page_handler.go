package domain_info_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
    "github.com/gorilla/mux"
)

type PageData struct {
    PageTitle  string
    DomainInfo api.DomainJson
}

var domainInfoPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    domainInfoPageTemplate = templater.BuildTemplate(GetDomainInfoPageTemplate)
}

func DomainInfoPageHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    domainName := vars["domainName"]

    templater.WriteTemplate(w, req, domainInfoPageTemplate, GetPageName(), GetPageData(domainName))
}

func GetUrl() string {
    return "/domains/{domainName}"
}

func GetPageName() string {
    return "DomainInfoPage"
}

func GetPageData(domainName string) PageData {
    pageData := PageData{
        PageTitle: "Domio - Available Domains",
        DomainInfo: api.GetDomainInfo(domainName),
    }

    return pageData
}
