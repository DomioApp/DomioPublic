package user_domains_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
    "log"
)

type PageData struct {
    PageTitle             string
    UserDomainsTopBarData ProfileTopBarData
    UserDomains           []api.DomainJson
}

var userDomainsPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    userDomainsPageTemplate = templater.BuildTemplate(GetUserDomainsPageTemplate)
}

func UserDomainsPageHandler(w http.ResponseWriter, req *http.Request) {
    var err error
    tokenCookie, err = req.Cookie("token")

    if (err != nil) {
        log.Print(err)
    }

    log.Print(tokenCookie.Value)

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
        UserDomains: GetUserDomains(tokenCookie.Value),
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

func GetUserDomains(token string) []api.DomainJson {
    return api.GetUserDomains(token)
}
