package home_handler

import (
    "html/template"
    "net/http"
    "github.com/tdewolff/minify"
    "github.com/tdewolff/minify/html"
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

    baseTemplate := template.New("BaseTemplate")
    //homeTemplate := template.New("HomeTemplate")

    baseTemplateContent := `
                        {{define "BaseTemplate"}}
                            <!DOCTYPE html>

                            <html lang="en">

                            <head>
                                <meta charset="UTF-8">
                                <title>{{.PageTitle}}</title>
                                <link rel="stylesheet" href="/style.css" />
                                {{template "head" .}}
                            </head>

                            <body>
                                {{template "body" .}}
                                <script src="/bundle.js"></script>
                            </body>

                            </html>
                        {{end}}`

    homeTemplateContent := `
                            {{define "head"}}{{end}}
                            {{define "body"}}
                                <h1>{{.PageTitle}}</h1>
                                {{.Body}}
                                {{.SideBar}}
                            {{end}}
                        `

    parsedBaseTemplate, _ := baseTemplate.Parse(baseTemplateContent)

    parsedHomeTemplate, _ := parsedBaseTemplate.Parse(homeTemplateContent)

    m := minify.New()
    m.AddFunc("text/html", html.Minify)

    mw := m.ResponseWriter(w, req)
    defer mw.Close()

    w = mw

    data := HomePageData{PageTitle:"Domio Home", Body:GetBody(), SideBar: GetSideBar()}

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

func GetSideBar() template.HTML {
    return template.HTML("<h3>sidebar</h3>")
}

func GetBody() template.HTML {
    return template.HTML("<h3>body</h3>")
}