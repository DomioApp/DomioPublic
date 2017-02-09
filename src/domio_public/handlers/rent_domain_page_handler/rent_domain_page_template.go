package rent_domain_page_handler

import (
    "html/template"
    "log"
)

func GetRentDomainPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "add_domain_form_template" .}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err2 := parsedTemplate.New("add_domain_form_template").Parse(GetRentDomainFormTemplate())

    if (err2 != nil) {
        log.Print(err2)
    }
}
