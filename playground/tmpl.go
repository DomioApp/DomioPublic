//http://stackoverflow.com/questions/41856021/how-to-parse-multiple-strings-into-a-template-with-go
//http://golang-examples.tumblr.com/post/87553422434/template-and-associated-templates
//https://hackernoon.com/golang-template-1-bcb690165663#.yyz27j1br

package main

import (
    "html/template"
    "os"
    "io"
)

type Data struct {
    Title          string
    SidebarContent string
}

func main() {
    data := Data{
        Title:"Hello there",
        SidebarContent:"123123123",
    }

    homePageTemplate := buildTemplate()

    writeTemplate(os.Stdout, homePageTemplate, data)
}

func writeTemplate(w io.Writer, tmpl *template.Template, data interface{}) {
    tmpl.Execute(w, data)
}

func buildTemplate() *template.Template {
    parsedTemplate, _ := template.New("base_template").Parse(`
                                                                {{define "base_template"}}
                                                                    I'm base, including inner: {{template "home_template" .}}
                                                 {{end}}`)

    parsedTemplate.New("sidebar_template").Parse(`{{define "sidebar_template"}}
                                                        Sidebar here:
                                                        {{.SidebarContent}}
                                                    {{end}}`)

    parsedTemplate.New("home_template").Parse(`{{define "home_template"}}I'm inner {{.Title}}. {{template "sidebar_template" .}}{{end}}`)

    return parsedTemplate
}