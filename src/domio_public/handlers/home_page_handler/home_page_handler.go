package home_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
)

type PageData struct {
    PageTitle   string

    SideBarData templater.SideBarData
}

var homePageTemplate *template.Template

func init() {
    homePageTemplate = templater.BuildTemplate(GetHomeTemplate)
}

func HomePageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, req, homePageTemplate, GetPageName(), GetPageData())

}

func GetUrl() string {
    return "/"
}

func GetPageName() string {
    return "HomePage"
}

func GetPageData() PageData {
    return PageData{
        PageTitle: "Domio - Marketplace for domains",
        SideBarData:templater.GetSideBarData(),
    }

}
