package home_page_handler

import (
    "html/template"
    "log"
    "domio_public/templater"
)

func GetHomeTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "sidebar_template" .}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err2 := parsedTemplate.New("sidebar_template").Parse(templater.GetSideBarTemplate())

    if (err2 != nil) {
        log.Print(err2)
    }
}
