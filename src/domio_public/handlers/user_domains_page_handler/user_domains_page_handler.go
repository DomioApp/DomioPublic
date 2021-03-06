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
        http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)

    } else {
        templater.WriteTemplate(w, req, userDomainsPageTemplate, GetPageName(), GetPageData())

    }

}

func GetUrl() string {
    return "/profile/domains"
}

func GetPageName() string {
    return "UserDomainsPage"
}

func GetPageData() PageData {
    pageData := PageData{
        PageTitle: "Domio - My Domains",
        UserDomainsTopBarData: GetUserDomainsTopBarData(),
    }

    log.Print("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
    log.Print(tokenCookie.Value)
    log.Print("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

    if (tokenCookie != nil) {
        pageData.UserDomains = api.GetUserDomains(tokenCookie.Value)
    }

    return pageData
}

func GetUserDomainsTopBarData() ProfileTopBarData {
    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"/profile/domains", Label:"My Domains", ClassName:"active"},
            {Url:"/profile/subscriptions", Label:"My Subscriptions"},
        },
    }
}