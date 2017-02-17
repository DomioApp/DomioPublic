package subscription_page_handler

func GetSubscriptionConfigMainViewTemplate() string {
    return `{{ define "subscription_config_mainview_template"}}

                {{with .PageData}}

                    <div class="b-subscription-config-mainview">
                        <form>

                        {{range .Entries}}

                            <div>
                                <label>{{.Name}}:<label>
                                <input></input>
                            </div>

                        {{end}}

                        <submit class="update-subscription-button" type="submit" value="{{.Subscription.Id}}">Save</submit>

                        </form>
                    </div>

                {{end}}

            {{end}}
        `
}