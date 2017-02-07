package add_domain_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
)

type PageData struct {
    PageTitle string

    FormData  FormData
}

var addDomainPageTemplate *template.Template

func init() {
    addDomainPageTemplate = templater.BuildTemplate(GetAddDomainPageTemplate)
}

func AddDomainPageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, req, addDomainPageTemplate, GetPageName(), GetPageData())

}

func GetUrl() string {
    return "/profile/domains/add"
}

func GetPageName() string {
    return "AddDomainPage"
}

func GetPageData() PageData {
    return PageData{
        PageTitle: "Domio - Marketplace for domains",
        FormData:GetFormData(),
    }

}
