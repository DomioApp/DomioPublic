package subscription_page_handler

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