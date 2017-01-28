package templater

import (
    "html/template"
    "net/http"
    "log"
)

type Link struct {
    Url       string
    Label     string
    ClassName string
}

type homeTemplateFn func(*template.Template)

type Data struct {
    SideBarTitle   string
    PageName       string
    PageTitle      string
    SidebarContent string
    SideBarLinks   []Link
}

func BuildTemplate(cb homeTemplateFn) *template.Template {

    parsedTemplate, parseErr := template.New("base_template").Parse(getBaseTemplateContent())

    if (parseErr != nil) {
        log.Fatalln(parseErr)
    }

    cb(parsedTemplate)
    return parsedTemplate
}

func WriteTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
    w.Header().Set("Content-Type", "text/html")

    //wr := gohtml.NewWriter(os.Stdout)

    execErr := tmpl.Execute(w, data)

    if (execErr != nil) {
        log.Print(execErr)
    }
}