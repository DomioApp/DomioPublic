package subscription_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
    "github.com/gorilla/mux"
    "log"
)

type PageData struct {
    PageTitle              string
    SubscriptionTopBarData ProfileTopBarData
    Subscription           api.Subscription
}

var subscriptionPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    subscriptionPageTemplate = templater.BuildTemplate(GetSubscriptionPageTemplate)
}

func SubscriptionPageHandler(w http.ResponseWriter, req *http.Request) {
    var err error

    tokenCookie, err = req.Cookie("token")

    if (err != nil) {
        http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)

    } else {

        vars := mux.Vars(req)
        id := vars["id"]

        templater.WriteTemplate(w, req, subscriptionPageTemplate, GetPageName(), GetPageData(id))

    }

}

func GetUrl() string {
    return "/profile/subscriptions/{id}"
}

func GetPageName() string {
    return "SubscriptionPage"
}

func GetPageData(id string) PageData {

    log.Print(api.GetSubscription(id, tokenCookie.Value))

    pageData := PageData{
        PageTitle: "Domio - Subscription Details",
        SubscriptionTopBarData: GetSubscriptionTopBarData(),
        Subscription: api.GetSubscription(id, tokenCookie.Value),
    }

    return pageData
}

func GetSubscriptionTopBarData() ProfileTopBarData {
    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"/profile/payments", Label:"My Payment Sources"},
            {Url:"/profile/payments/add", Label:"Add a Payment Source"},
            {Url:"/profile/payments/history", Label:"Payments History"},
        },
    }
}