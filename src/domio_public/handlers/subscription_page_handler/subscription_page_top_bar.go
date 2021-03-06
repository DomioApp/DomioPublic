package subscription_page_handler

import "domio_public/templater"

type ProfileTopBarData struct {
    Links []templater.Link
}

func GetSubscriptionTopBarTemplate() string {
    return `{{ define "subscriptions_topbar_template"}}

                {{with .PageData.SubscriptionTopBarData}}

                    <div class="b-profile-top-bar-container">
                        {{range .Links}}
                             <a href="{{.Url}}">{{.Label}}</a>
                        {{end}}
                    </div>

                {{end}}

            {{end}}
        `
}