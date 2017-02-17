package subscription_page_handler

func GetSubscriptionConfigMainViewTemplate() string {
    return `{{ define "subscription_config_mainview_template"}}

                {{with .PageData.Subscription}}

                    <div class="b-subscription-config-mainview">
                        <form>
                            <div>
                                <label>CNAME:<label>
                                <input />
                            <div>
                        </form>
                    </div>

                {{end}}

            {{end}}
        `
}