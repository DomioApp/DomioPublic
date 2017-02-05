package user_domains_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
)

type PageData struct {
    PageTitle             string
    UserDomainsTopBarData ProfileTopBarData
    UserDomains           []api.DomainJson
}

var userDomainsPageTemplate *template.Template

func init() {
    userDomainsPageTemplate = templater.BuildTemplate(GetUserDomainsPageTemplate)
}

func UserDomainsPageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, userDomainsPageTemplate, GetPageName(), GetPageData())

}

func GetUrl() string {
    return "/profile/domains"
}

func GetPageName() string {
    return "UserDomainsPage"
}

func GetPageData() PageData {
    return PageData{
        PageTitle: "Domio - My Domains",
        UserDomainsTopBarData: GetUserDomainsTopBarData(),
        UserDomains: GetUserDomains(),
    }

}

func GetUserDomainsTopBarData() ProfileTopBarData {
    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"/profile/domains", Label:"My Domains"},
            {Url:"/profile/domains/rented", Label:"Domains I rent"},
            {Url:"/profile/stats", Label:"Stats"},
        },
    }
}

func GetUserDomains() []api.DomainJson {
    return api.GetUserDomains()
}
