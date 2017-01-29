package home_page_handler

import (
    "html/template"
    "log"
)

func GetHomeTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                    {{template "sidebar_template" .}}
                                                {{end}}`)
    if (err != nil) {
        log.Print(err)
    }

    _, err2 := parsedTemplate.New("sidebar_template").Parse(`{{ define "sidebar_template"}}
                                                        <div class="b-side-bar-container">
                                                            <h4>{{.PageData.SideBarTitle}}</h4>

                                                            {{range .PageData.SideBarLinks}}
                                                                 <a href="{{.Url}}">{{.Label}}</a>
                                                            {{end}}
                                                        </div>
                                                    {{end}}`)
    if (err2 != nil) {
        log.Print(err2)
    }

}
