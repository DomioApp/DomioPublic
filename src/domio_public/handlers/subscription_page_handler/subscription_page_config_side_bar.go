package subscription_page_handler

func GetSubscriptionConfigSideBarTemplate() string {
    return `{{ define "subscription_config_sidebar_template"}}

                {{with .PageData.Subscription}}
                    <div>
                        <div><span>Domain:</span> <span>{{.Metadata.Domain}}</span></div>

                        <hr>

                        <div><span>ID:</span> <span>{{.Id}}</span></div>
                        <div><span>Status:</span> <span>{{.Status}}</span></div>
                        <div><span>Created:</span> <span>{{.Created}}</span></div>

                        <hr>

                        <div><span>Plan:</span> <span>{{.Plan.Id}}</span></div>
                        <div><span>Amount:</span> <span>{{.Plan.Amount}}</span></div>
                        <div><span>Created:</span> <span>{{.Plan.Created}}</span></div>
                        <div><span>Currency:</span> <span>{{.Plan.Currency}}</span></div>
                        <div><span>Interval:</span> <span>{{.Plan.Interval}}</span></div>

                        <hr>
                        <button value="{{.Id}}" class="delete-subscription-button">Cancel subscription</button>
                    </div>
                {{end}}

            {{end}}
        `
}