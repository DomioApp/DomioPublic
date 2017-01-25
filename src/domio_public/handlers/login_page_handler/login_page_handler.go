package login_page_handler

import (
    "html/template"
    "net/http"
    "log"
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

var homeTemplate *template.Template

func init() {
    var err error

    homeTemplate, err = GetTemplate();

    if (err != nil) {
        log.Print(err)
    }
}

func LoginPageHandler(w http.ResponseWriter, req *http.Request) {

    data := HomePageData{
        PageName: "LoginPage",
        PageTitle: "Domio Login",

        TopBar: GetTopBar(),
        MainArea: GetMainArea(),
        SideBar: GetSideBar(),
    }

    templater.WritePage(w, homeTemplate, data)

}

func initMinifier(w http.ResponseWriter, req *http.Request) {
    m := minify.New()
    m.AddFunc("text/html", html.Minify)

    mw := m.ResponseWriter(w, req)
    defer mw.Close()

    w = mw
}
func GetUrl() string {
    return "/login"
}