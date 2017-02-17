package subscription_page_handler

import (
    "html/template"
    "log"
)

func GetSubscriptionPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "subscriptions_topbar_template" .}}
                                                            {{template "subscription_config_topbar_template" .}}
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

                                                                    <form>
                                                                        <div>
                                                                            <label>CNAME:<label>
                                                                            <input />
                                                                        <div>
                                                                    </form>

                                                                    <hr>
                                                                    <button value="{{.Id}}" class="delete-subscription-button">Cancel subscription</button>
                                                                </div>
                                                            {{end}}
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err2 := parsedTemplate.New("subscriptions_topbar_template").Parse(GetSubscriptionTopBarTemplate())

    if (err2 != nil) {
        log.Print(err2)
    }

    _, err3 := parsedTemplate.New("subscription_config_topbar_template").Parse(GetSubscriptionConfigTopBarTemplate())

    if (err3 != nil) {
        log.Print(err3)
    }
}