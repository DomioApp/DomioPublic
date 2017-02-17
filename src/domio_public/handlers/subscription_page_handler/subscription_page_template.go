package subscription_page_handler

import (
    "html/template"
    "log"
)

func GetSubscriptionPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            {{template "subscriptions_topbar_template" .}}
                                                            {{template "subscription_config_topbar_template" .}}

                                                            {{template "subscription_config_sidebar_template" .}}
                                                            {{template "subscription_config_mainview_template" .}}
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
    _, err4 := parsedTemplate.New("subscription_config_sidebar_template").Parse(GetSubscriptionConfigSideBarTemplate())

    if (err4 != nil) {
        log.Print(err4)
    }

    _, err5 := parsedTemplate.New("subscription_config_mainview_template").Parse(GetSubscriptionConfigMainViewTemplate())

    if (err5 != nil) {
        log.Print(err5)
    }
}