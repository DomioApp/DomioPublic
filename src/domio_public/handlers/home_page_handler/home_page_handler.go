package home_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
)

type PageData struct {
    PageTitle      string

    SideBarTitle   string
    SidebarContent string
    SideBarLinks   []templater.Link
}

var homePageTemplate *template.Template

func init() {
    homePageTemplate = templater.BuildTemplate(GetHomeTemplate)
}

func HomePageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, homePageTemplate, GetPageName(), GetPageData())

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
        SidebarContent: "sidebar here",
        SideBarTitle: "Categories",
        SideBarLinks:[]templater.Link{
            {Url:"/home", Label:"Home"},
        },
    }

}
