package templater

type TopBarData struct {
    LeftColumnLinks  []Link
    RightColumnLinks []Link
    DomainAddLink    Link
}

func GetTopBarTemplate() string {
    return `{{ define "top_bar_template"}}

                {{with .TopBarData}}

                    <div class="b-top-bar-container">
                        <div class="left-area">
                            {{range .LeftColumnLinks}}
                                 <a href="{{.Url}}">{{.Label}}</a>
                            {{end}}
                        </div>
                        <div class="right-area">

                            {{with .DomainAddLink}}
                                <a href="{{.Url}}" class="{{.ClassName}}">{{.Label}}</a>
                            {{end}}

                            {{range .RightColumnLinks}}
                                 <a href="{{.Url}}">{{.Label}}</a>
                            {{end}}
                        </div>
                    </div>

                {{end}}

            {{end}}
        `
}

func GetTopBarData() TopBarData {
    topBarData := TopBarData{
        LeftColumnLinks:[]Link{
            {Url:"/", Label:"Home"},
            {Url:"/domains", Label:"Domains"},
        },
        RightColumnLinks:[]Link{
            {Url:"/login", Label:"Login"},
            {Url:"/signup", Label:"Signup"},
        },
        DomainAddLink:Link{Url:"/domains/add", Label:"Add Domain", ClassName:"b-top-bar-container__domain-add-link"},
    }

    return topBarData
}