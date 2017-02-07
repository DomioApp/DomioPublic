package domains_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
)

type PageData struct {
    PageTitle   string

    SideBarData templater.SideBarData
}

var domainsPageTemplate *template.Template

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
    return PageData{
        PageTitle: "Domio - Marketplace for domains. Domains Listings",
        SideBarData:templater.GetSideBarData(),
    }

}
