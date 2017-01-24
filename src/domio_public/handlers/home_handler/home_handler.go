package home_handler

import (
    "html/template"
    "net/http"
    //"github.com/tdewolff/minify"
    //"github.com/tdewolff/minify/html"
    "domio_public/templater"
)

type HomePageData struct {
    PageTitle string
    PageName  string
    Header    template.HTML
    Body      template.HTML
    SideBar   template.HTML
}

var parsedTemplate *template.Template

func init() {
    //parsedTemplate, _ = GetTemplate()
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    baseTemplate := base_template.GetBaseTemplate()
    parsedHomeTemplate, _ := baseTemplate.Parse(getHomeTemplateContent())

    /*
    m := minify.New()
    m.AddFunc("text/html", html.Minify)

    mw := m.ResponseWriter(w, req)
    defer mw.Close()

    w = mw
    */

    data := HomePageData{
        PageTitle: "Domio Home",
        Body: GetBody(),
        SideBar: GetSideBar(),
    }

    parsedHomeTemplate.ExecuteTemplate(w, "BaseTemplate", data)

    //appStatusInfo := api.GetAPIStatus()

    //t := template.New("HomeTemplate")

    /*
    pageData := HomePageData{
        PageName: "home",
        Title: "Domio",
        Header: template.HTML(TopBar),
        Body: template.HTML(HomePageTemplate),
    }
    */

    //baseTemplate, err := t.Parse(templater.GetBaseTemplateAsString())

    //if (err != nil) {
    //    log.Print(err)
    //}

    //var doc bytes.Buffer
    //baseTemplate.Execute(&doc, pageData)

    /*
    homePageData := PageData{
        Title:    "Domio",
        TopBar:   templater.GetTopBar(),
        AppStatusInfoBar: templater.GetAppStatusInfoBar(appStatusInfo),
    }
    */


}

func getHomeTemplateContent() string {
    homeTemplateContent := `
                            {{define "head"}}{{end}}
                            {{define "body"}}
                                <h1>{{.PageTitle}}</h1>
                                {{.Body}}
                                {{.SideBar}}
                            {{end}}
                        `
    return homeTemplateContent
}

func GetSideBar() template.HTML {
    return template.HTML("<h3>sidebar</h3>")
}

func GetBody() template.HTML {
    return template.HTML("<h3>body</h3>")
}