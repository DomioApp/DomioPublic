package subscription_page_handler

func GetSubscriptionConfigMainViewTemplate() string {
    return `{{ define "subscription_config_mainview_template"}}

                {{with .PageData}}

                    <div class="b-subscription-config-mainview">
                        <form>

                        {{range .Entries}}

                            <div>
                                <label>{{.Name}}:<label>
                                <input />
                            <div>

                        {{end}}

                        </form>
                    </div>

                {{end}}

            {{end}}
        `
}