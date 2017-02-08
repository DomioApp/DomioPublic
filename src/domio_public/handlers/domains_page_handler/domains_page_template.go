package domains_page_handler

import (
    "html/template"
    "log"
)

func GetDomainsPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{range .PageData.Domains}}
                                                                <div data-domain="{{.Name}}">
                                                                    <a href="/domains/{{.Name}}">{{.Name}}: ${{.PricePerMonth}} per month</a>
                                                                </div>
                                                            {{end}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }
}