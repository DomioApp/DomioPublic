package add_payment_source_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
    "github.com/gorilla/mux"
    "log"
)

type PageData struct {
    PageTitle  string
    DomainInfo *api.DomainJson
}

var domainEditPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    domainEditPageTemplate = templater.BuildTemplate(GetAddPaymentSourcePageTemplate)
}

func AddPaymentSourcePageHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    domainName := vars["domainName"]

    var tokenErr error

    tokenCookie, tokenErr = req.Cookie("token")

    log.Print(tokenErr)

    templater.WriteTemplate(w, req, domainEditPageTemplate, GetPageName(), GetPageData(domainName))
}

func GetUrl() string {
    return "/profile/payments/add"
}

func GetPageName() string {
    return "AddPaymentSourcePage"
}

func GetPageData(domainName string) PageData {
    domainInfo, domainInfoError := api.GetDomainInfo(domainName, tokenCookie.Value)

    if (domainInfoError != nil) {
        log.Print(domainInfoError)
    }

    pageData := PageData{
        PageTitle: "Domio - Add Payment Source ",
        DomainInfo: domainInfo,
    }

    return pageData
}
