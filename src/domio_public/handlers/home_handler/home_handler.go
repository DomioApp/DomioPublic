package home_handler

import (
    "html/template"
    "net/http"
    "github.com/tdewolff/minify"
    "github.com/tdewolff/minify/html"
)

type HomePageData struct {
    Title    string
    PageName string
    Header   template.HTML
    Body     template.HTML
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
                            <head>{{template "head" .}}</head>
                            <body>{{template "body" .}}</body>
                        {{end}}`

    homeTemplateContent := `
                            {{define "head"}}<title>Domio</title>{{end}}
                            {{define "body"}}index{{end}}
                        `

    parsedBaseTemplate, _ := baseTemplate.Parse(baseTemplateContent)

    parsedHomeTemplate, _ := parsedBaseTemplate.Parse(homeTemplateContent)


    m := minify.New()
    m.AddFunc("text/html", html.Minify)

    mw := m.ResponseWriter(w, req)
    defer mw.Close()

    w = mw

    parsedHomeTemplate.ExecuteTemplate(w, "BaseTemplate", nil)

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
