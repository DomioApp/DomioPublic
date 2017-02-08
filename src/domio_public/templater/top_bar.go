package templater

import "log"

type TopBarData struct {
    LeftColumnLinks  []Link
    RightColumnLinks []Link
}

func GetTopBarTemplate() string {
    return `{{ define "top_bar_template"}}

                {{with .TopBarData}}

                    <div class="b-top-bar-container">

                        <div class="left-area">
                            {{range .LeftColumnLinks}}
                                 <a class="{{.ClassName}}" href="{{.Url}}">{{.Label}}</a>
                            {{end}}
                        </div>

                        <div class="right-area">
                            {{range .RightColumnLinks}}
                                 <a class="{{.ClassName}}" href="{{.Url}}">{{.Label}}</a>
                            {{end}}
                        </div>

                    </div>

                {{end}}

            {{end}}
        `
}

func GetTopBarData(pageName string, userName string) TopBarData {
    var customLinks []Link

    if (userName != "") {
        customLinks = []Link{
            {Url:"/profile/domains/add", Label:"Add Domain", ClassName:"b-top-bar-container__domain-add-link"},
            {Url:"/profile/domains", Label:"My Domains"},
            {Url:"/profile", Label:userName},
        }
    } else {
        customLinks = []Link{
            {Url:"/profile/domains/add", Label:"Add Domain", ClassName:"b-top-bar-container__domain-add-link"},
            {Url:"/login", Label:"Login"},
            {Url:"/signup", Label:"Signup"},
        }
    }

    dataset := TopBarData{
        LeftColumnLinks:[]Link{
            {Url:"/", Label:"Home"},
            {Url:"/domains", Label:"Domains"},
        },
        RightColumnLinks:customLinks,
    }

    return dataset
}