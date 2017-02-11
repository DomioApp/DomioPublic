package user_payment_sources_page_handler

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

var userPaymentSourcesPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    userPaymentSourcesPageTemplate = templater.BuildTemplate(GetUserPaymentsSourcesPageTemplate)
}

func UserPaymentSourcesPageHandler(w http.ResponseWriter, req *http.Request) {
    var err error

    tokenCookie, err = req.Cookie("token")

    if (err != nil) {
        http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)

    } else {
        templater.WriteTemplate(w, req, userPaymentSourcesPageTemplate, GetPageName(), GetPageData())

    }

}

func GetUrl() string {
    return "/profile/payments"
}

func GetPageName() string {
    return "UserPaymentSourcesPage"
}

func GetPageData() PageData {
    pageData := PageData{
        PageTitle: "Domio - My Payment Sources",
        UserDomainsTopBarData: GetUserDomainsTopBarData(),
    }

    if (tokenCookie != nil) {
        pageData.UserDomains = GetUserDomains(tokenCookie.Value)
    }

    return pageData
}

func GetUserDomainsTopBarData() ProfileTopBarData {
    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"/profile/payments", Label:"My Payment Sources"},
            {Url:"/profile/payments/add", Label:"Add a Payment Source"},
            {Url:"/profile/payments/history", Label:"Payments History"},
        },
    }
}

func GetUserDomains(token string) []api.DomainJson {
    return api.GetUserDomains(token)
}
