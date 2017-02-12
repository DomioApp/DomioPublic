package subscriptions_page_handler

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
    Subscriptions         []api.Subscription
}

var rentedDomainsPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    rentedDomainsPageTemplate = templater.BuildTemplate(GetUserDomainsPageTemplate)
}

func SubscriptionsPageHandler(w http.ResponseWriter, req *http.Request) {
    var err error

    tokenCookie, err = req.Cookie("token")

    if (err != nil) {
        http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)

    } else {
        templater.WriteTemplate(w, req, rentedDomainsPageTemplate, GetPageName(), GetPageData())

    }

}

func GetUrl() string {
    return "/profile/subscriptions"
}

func GetPageName() string {
    return "RentedDomainsPage"
}

func GetPageData() PageData {
    pageData := PageData{
        PageTitle: "Domio - My Subscriptions",
        UserDomainsTopBarData: GetRentedDomainsTopBarData(),
    }

    if (tokenCookie != nil) {
        pageData.Subscriptions = api.GetSubscriptions(tokenCookie.Value)
    }

    log.Print(pageData.Subscriptions)

    return pageData
}

func GetRentedDomainsTopBarData() ProfileTopBarData {
    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"/profile/domains", Label:"My Domains"},
            {Url:"/profile/subscriptions", Label:"My Subscriptions", ClassName: "active"},
        },
    }
}
