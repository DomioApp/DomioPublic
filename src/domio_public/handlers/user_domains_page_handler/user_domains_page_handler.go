package user_domains_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
)

type UserDomain struct {
    Name string
}

type PageData struct {
    PageTitle             string
    UserDomainsTopBarData ProfileTopBarData
    UserDomains           []UserDomain
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
            {Url:"/profile/domains", Label:"Domains"},
            {Url:"/profile/domains/rented", Label:"Domains I rent"},
            {Url:"/profile/stats", Label:"Stats"},
        },
    }
}

func GetUserDomains() []UserDomain {
    return []UserDomain{
        {Name:"john.com"},
        {Name:"jack.com"},
    }
}
