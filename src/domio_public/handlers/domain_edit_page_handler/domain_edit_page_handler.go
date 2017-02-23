package domain_edit_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/components/api"
    "github.com/gorilla/mux"
    "log"
    "domio_public/errors"
)

type PageData struct {
    PageTitle  string
    DomainInfo *api.DomainJson
}

var domainEditPageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    domainEditPageTemplate = templater.BuildTemplate(GetDomainEditPageTemplate)
}

func DomainEditPageHandler(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    domainName := vars["domainName"]

    var tokenErr error
    log.Print(tokenErr)

    tokenCookie, tokenErr = req.Cookie("token")

    domainInfo, domainInfoError := api.GetDomainInfo(domainName, tokenCookie.Value)

    if (domainInfoError != nil) {
        log.Print(domainInfoError)
        if (domainInfoError == &errors.DomainNotFound) {
            templater.WriteTemplate(w, req, domainEditPageTemplate, "Domain not found", GetPageData("Domain not found", domainInfo))
            return
        }
    } else {
        templater.WriteTemplate(w, req, domainEditPageTemplate, GetPageName(), GetPageData(domainName, domainInfo))
    }

}

func GetUrl() string {
    return "/profile/domain/{domainName}"
}

func GetPageName() string {
    return "DomainEditPage"
}

func GetPageData(domainName string, domainInfo *api.DomainJson) PageData {

    pageData := PageData{
        PageTitle: "Domio - Edit Domain " + domainName,
        DomainInfo: domainInfo,
    }

    return pageData
}
