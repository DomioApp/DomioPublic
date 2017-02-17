package subscription_page_handler

func GetSubscriptionConfigMainViewTemplate() string {
    return `{{ define "subscription_config_mainview_template"}}

                {{with .PageData.Subscription}}

                    <form>
                        <div>
                            <label>CNAME:<label>
                            <input />
                        <div>
                    </form>

                {{end}}

            {{end}}
        `
}