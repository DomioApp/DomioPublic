package user_domains_page_handler

import (
    "html/template"
    "log"
)

func GetUserDomainsPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "user_domains_topbar_template" .}}
                                                            {{range .PageData.UserDomains}}
                                                                <div class="b-domains-list-container">
                                                                    <a href="/profile/domain/{{.Name}}">{{.Name}}</a>
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