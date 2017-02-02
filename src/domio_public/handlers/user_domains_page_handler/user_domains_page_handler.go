package user_domains_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
    "domio_public/widgets"
)

type PageData struct {
    PageTitle             string
    UserDomainsTopBarData widgets.ProfileTopBarData
    SideBarData           templater.SideBarData
}

var profilePageTemplate *template.Template

func init() {
    profilePageTemplate = templater.BuildTemplate(GetProfilePageTemplate)
}

func UserDomainsPageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, profilePageTemplate, GetPageName(), GetPageData())

}

func GetUrl() string {
    return "/profile/domains"
}

func GetPageName() string {
    return "ProfilePage"
}

func GetPageData() PageData {
    return PageData{
        PageTitle: "Domio - Marketplace for domains",
        SideBarData: templater.GetSideBarData(),
        UserDomainsTopBarData: GetUserDomainsTopBarData(),
    }

}

func GetUserDomainsTopBarData() widgets.ProfileTopBarData {

    return widgets.ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"/profile/domains", Label:"Domains"},
            {Url:"/profile/domains/rented", Label:"Domains I rent"},
            {Url:"/profile/stats", Label:"Stats"},
        },
    }
}