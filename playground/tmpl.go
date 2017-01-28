//http://stackoverflow.com/questions/41856021/how-to-parse-multiple-strings-into-a-template-with-go
//http://golang-examples.tumblr.com/post/87553422434/template-and-associated-templates
//https://hackernoon.com/golang-template-1-bcb690165663#.yyz27j1br

package main

import (
	"html/template"
	"io"
	"os"
)

type homeTemplateFn func(*template.Template)

type Data struct {
	Title          string
	SidebarContent string
}

func main() {
	data := Data{
		Title:          "Hello there",
		SidebarContent: "123123123",
	}

	homePageTemplate := BuildTemplate(GetHomeTemplate)

	writeTemplate(os.Stdout, homePageTemplate, data)
}

func writeTemplate(w io.Writer, tmpl *template.Template, data interface{}) {
	tmpl.Execute(w, data)
}

func BuildTemplate(cb homeTemplateFn) *template.Template {

	var basetemplate_name = "base_template"
	parsedTemplate, _ := template.New(basetemplate_name).Parse(`{{define "` + basetemplate_name + `"}}
                                                                    {{template "main_template" .}}
                                                                {{end}}`)
	cb(parsedTemplate)
	return parsedTemplate
}

func GetHomeTemplate(parsedTemplate *template.Template) {
	parsedTemplate.New("main_template").Parse(`{{define "main_template"}}I'm inner {{.Title}}. {{template "sidebar_template" .}}{{end}}`)

	parsedTemplate.New("sidebar_template").Parse(`{{define "sidebar_template"}}
                                                        Sidebar here:
                                                        {{.SidebarContent}}
                                                    {{end}}`)
}
