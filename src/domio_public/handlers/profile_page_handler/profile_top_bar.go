package profile_page_handler

import "domio_public/templater"

type ProfileMainAreaData struct {
    Links []templater.Link
}

func GetProfileMainAreaTemplate() string {
    return `{{ define "profile_main_area_template"}}

                {{with .PageData.ProfileTopBarData}}

                    <div class="b-profile-top-bar-container">
                        {{range .Links}}
                             <a href="{{.Url}}">{{.Label}}</a>
                        {{end}}
                    </div>

                {{end}}

            {{end}}
        `
}