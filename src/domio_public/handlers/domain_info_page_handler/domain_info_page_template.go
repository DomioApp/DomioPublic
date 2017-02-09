package domain_info_page_handler

import (
    "html/template"
    "log"
)

func GetDomainInfoPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{with .PageData.DomainInfo}}
                                                                <div data-domain="{{.Name}}">
                                                                    <a href="/domains/{{.Name}}">{{.Name}}</a>
                                                                    <p>Price Per Month: ${{.PricePerMonth}}</p>
                                                                    <a href="/domain/{{.Name}}/rent">Rent {{.Name}} now</a>
                                                                </div>
                                                            {{end}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }
}