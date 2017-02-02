package user_domains_page_handler

import (
    "html/template"
    "log"
    "domio_public/widgets"
)

func GetProfilePageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "profile_topbar_template" .}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err2 := parsedTemplate.New("profile_topbar_template").Parse(widgets.GetProfileTopBarTemplate())

    if (err2 != nil) {
        log.Print(err2)
    }
}
