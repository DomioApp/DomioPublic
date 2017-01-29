package profile_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
)

type PageData struct {
    PageTitle   string

    SideBarData templater.SideBarData
}

var profilePageTemplate *template.Template

func init() {
    profilePageTemplate = templater.BuildTemplate(GetHomeTemplate)
}

func ProfilePageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, profilePageTemplate, GetPageName(), GetPageData())

}

func GetUrl() string {
    return "/profile"
}

func GetPageName() string {
    return "ProfilePage"
}

func GetPageData() PageData {
    return PageData{
        PageTitle: "Domio - Marketplace for domains",
        SideBarData:templater.GetSideBarData(),
    }

}
