package user_payment_source_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
    "github.com/gorilla/mux"
)

type PageData struct {
    PageTitle                    string
    UserPaymentSourcesTopBarData ProfileTopBarData
    UserPaymentSource            api.PaymentSource
}

var userPaymentSourcePageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    userPaymentSourcePageTemplate = templater.BuildTemplate(GetUserPaymentSourcePageTemplate)
}

func UserPaymentSourcePageHandler(w http.ResponseWriter, req *http.Request) {
    var err error

    tokenCookie, err = req.Cookie("token")

    if (err != nil) {
        http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)

    } else {

        vars := mux.Vars(req)
        id := vars["id"]

        templater.WriteTemplate(w, req, userPaymentSourcePageTemplate, GetPageName(), GetPageData(id))

    }

}

func GetUrl() string {
    return "/profile/payments/{id}"
}

func GetPageName() string {
    return "UserPaymentSourcePage"
}

func GetPageData(id string) PageData {
    pageData := PageData{
        PageTitle: "Domio - My Cards",
        UserPaymentSourcesTopBarData: GetUserDomainsTopBarData(),
        UserPaymentSource: api.GetUserPaymentSource(tokenCookie.Value, id),
    }

    return pageData
}

func GetUserDomainsTopBarData() ProfileTopBarData {
    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"/profile/payments", Label:"My Cards"},
            {Url:"/profile/payments/add", Label:"Add a Card"},
        },
    }
}