package subscription_page_handler

func GetSubscriptionConfigMainViewTemplate() string {
    return `{{ define "subscription_config_mainview_template"}}

                {{with .PageData}}

                    <div class="b-subscription-config-mainview">

                        <div>

                            {{range .Entries}}

                                <div>
                                    <label>{{.Name}}:<label>
                                </div>

                            {{end}}

                        </div>

                        <form class="b-subscription-record-form">
                            <input class="value-input"></input>

                            <input type="submit" class="update-subscription-button" type="submit" value="Save"></input>
                        </form>

                    </div>

                {{end}}

            {{end}}
        `
}