package home_page_handler

import (
    "html/template"
    "net/http"
    "github.com/tdewolff/minify"
    "github.com/tdewolff/minify/html"
    "domio_public/templater"
)

type HomePageData struct {
    PageTitle string
    PageName  string

    Header    template.HTML
    TopBar    template.HTML
    MainArea  template.HTML
    SideBar   template.HTML
}

var homePageTemplate *template.Template

func init() {
    homePageTemplate = templater.BuildTemplate(GetHomeTemplate)
}

func HomePageHandler(w http.ResponseWriter, req *http.Request) {

    data := templater.Data{
        PageName: "HomePage",
        PageTitle: "Domio - Marketplace for domains",
        SidebarContent: "sidebar here",
        SideBarTitle: "Categories",
        SideBarLinks:[]templater.Link{
            {Url:"/home", Label:"Home"},
        },
    }

    templater.WriteTemplate(w, homePageTemplate, data)

}

func initMinifier(w http.ResponseWriter, req *http.Request) {
    m := minify.New()
    m.AddFunc("text/html", html.Minify)

    mw := m.ResponseWriter(w, req)
    defer mw.Close()

    w = mw
}
func GetUrl() string {
    return "/"
}