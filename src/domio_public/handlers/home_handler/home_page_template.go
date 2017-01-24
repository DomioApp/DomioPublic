package home_handler

import (
    "domio_public/templater"
    "html/template"
    "log"
)

const HomePageTemplate = `
                            {{.LeftColumn}}
                            {{.RightColumn}}
                            {{.AppStatusInfoBar}}
                          `
const TopBar = `{{.TopBar}}`

func GetTemplate() (*template.Template, error) {
    baseTemplate := templater.GetBaseTemplateAsString()

    t := template.New("temp_tmpl")

    homeTemplate, err := t.Parse(baseTemplate)

    if (err != nil) {
        log.Print(err)
    }

    return homeTemplate, nil
}