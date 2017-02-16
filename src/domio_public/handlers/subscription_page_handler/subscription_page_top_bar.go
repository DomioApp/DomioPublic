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
func GetSubscriptionConfigTopBarTemplate() string {
    return `{{ define "subscription_config_topbar_template"}}

                {{with .PageData.SubscriptionConfigTopBarData}}

                    <div class="b-profile-config-top-bar-container">
                        {{range .Links}}
                             <a href="{{.Url}}">{{.Label}}</a>
                        {{end}}
                    </div>

                {{end}}

            {{end}}
        `
}