package user_payment_sources_page_handler

import "domio_public/templater"

type ProfileTopBarData struct {
    Links []templater.Link
}

func GetUserDomainsTopBarTemplate() string {
    return `{{ define "user_domains_topbar_template"}}

                {{with .PageData.UserPaymentSourcesTopBarData}}

                    <div class="b-profile-top-bar-container">
                        {{range .Links}}
                             <a href="{{.Url}}">{{.Label}}</a>
                        {{end}}
                    </div>

                {{end}}

            {{end}}
        `
}