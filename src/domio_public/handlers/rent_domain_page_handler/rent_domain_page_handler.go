package rent_domain_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
)

type PageData struct {
    PageTitle string

    FormData  FormData
}

var rentDomainPageTemplate *template.Template

func init() {
    rentDomainPageTemplate = templater.BuildTemplate(GetRentDomainPageTemplate)
}

func RentDomainPageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, req, rentDomainPageTemplate, GetPageName(), GetPageData())

}

func GetUrl() string {
    return "/domain/{domainName}/rent"
}

func GetPageName() string {
    return "RentDomainPage"
}

func GetPageData() PageData {
    return PageData{
        PageTitle: "Domio - Marketplace for domains. Rent domain",
        FormData:GetFormData(),
    }

}
