package home_handler

import (
    "html/template"
    "domio_public/templater"
    "log"
)

func GetTemplate() (*template.Template, error) {
    baseTemplate := templater.GetBaseTemplate()
    parsedHomeTemplate, parseErr := baseTemplate.Parse(getHomeTemplateContent())

    if (parseErr != nil) {
        log.Print(parseErr)
    }

    return parsedHomeTemplate, nil
}

func getHomeTemplateContent() string {
    homeTemplateContent := `
                            {{define "head"}}{{end}}
                            {{define "body"}}
                                {{.TopBar}}
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