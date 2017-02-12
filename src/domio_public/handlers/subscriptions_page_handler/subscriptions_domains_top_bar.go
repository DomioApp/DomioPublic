package subscriptions_page_handler

import "domio_public/templater"

type ProfileTopBarData struct {
    Links []templater.Link
}

func GetRentedDomainsTopBarTemplate() string {
    return `{{ define "user_domains_topbar_template"}}

                {{with .PageData.UserDomainsTopBarData}}

                    <div class="b-profile-top-bar-container">
                        {{range .Links}}
                             <a class="{{.ClassName}}" href="{{.Url}}">{{.Label}}</a>
                        {{end}}
                    </div>

                {{end}}

            {{end}}
        `
}