package subscription_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
    "github.com/gorilla/mux"
)

type PageData struct {
    PageTitle                    string
    SubscriptionTopBarData       ProfileTopBarData
    SubscriptionConfigTopBarData ProfileTopBarData
    Subscription                 api.Subscription
    Records                      []api.Record
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
        subId := vars["id"]

        templater.WriteTemplate(w, req, subscriptionPageTemplate, GetPageName(), GetPageData(subId))

    }

}

func GetUrl() string {
    return "/profile/subscriptions/{id}"
}

func GetPageName() string {
    return "SubscriptionPage"
}

func GetPageData(subId string) PageData {

    //log.Print(api.GetSubscription(id, tokenCookie.Value))

    pageData := PageData{
        PageTitle: "Domio - Subscription Details",
        SubscriptionTopBarData: GetSubscriptionTopBarData(),
        SubscriptionConfigTopBarData: GetSubscriptionConfigTopBarData(),
        Subscription: api.GetSubscription(subId, tokenCookie.Value),
        Records:api.GetSubscriptionRecords(subId, tokenCookie.Value),
    }

    return pageData
}

func GetSubscriptionTopBarData() ProfileTopBarData {
    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"", Label:"Subscription Details"},
            {Url:"/profile/subscriptions", Label:"My Subscriptions"},
        },
    }
}
func GetSubscriptionConfigTopBarData() ProfileTopBarData {
    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"", Label:"DNS Config"},
        },
    }
}