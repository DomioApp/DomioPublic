package user_payment_sources_page_handler

import (
    "html/template"
    "log"
)

func GetUserPaymentsSourcesPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "user_domains_topbar_template" .}}
                                                            {{range .PageData.UserPaymentSources}}
                                                                <div>
                                                                    <hr>
                                                                    <div><span>ID:</span> <span>{{.Id}}</span></div>
                                                                    <div><span>Name on card:</span> <span>{{.Name}}</span></div>
                                                                    <div><span>Expiration:</span> <span>{{.ExpMonth}}/{{.ExpYear}}</span></div>
                                                                    <div><span>Last 4 Digits:</span> <span>{{.LastFour}}</span></div>
                                                                    <div><span>Brand:</span> <span>{{.Brand}}</span></div>
                                                                    <div><span>Currency:</span> <span>{{.Currency}}</span></div>
                                                                    <div><span>Country:</span> <span>{{.Country}}</span></div>
                                                                </div>
                                                            {{end}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err2 := parsedTemplate.New("user_domains_topbar_template").Parse(GetUserDomainsTopBarTemplate())

    if (err2 != nil) {
        log.Print(err2)
    }
}