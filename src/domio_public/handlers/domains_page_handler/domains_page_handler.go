package domains_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
)

type PageData struct {
    PageTitle string
    Domains   []api.DomainJson
}

var domainsPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    domainsPageTemplate = templater.BuildTemplate(GetDomainsPageTemplate)
}

func DomainsPageHandler(w http.ResponseWriter, req *http.Request) {
    templater.WriteTemplate(w, req, domainsPageTemplate, GetPageName(), GetPageData())
}

func GetUrl() string {
    return "/domains"
}

func GetPageName() string {
    return "DomainsPage"
}

func GetPageData() PageData {
    pageData := PageData{
        PageTitle: "Domio - Available Domains",
        Domains: GetDomains(),
    }

    return pageData
}

func GetDomains() []api.DomainJson {
    return api.GetDomains()
}
