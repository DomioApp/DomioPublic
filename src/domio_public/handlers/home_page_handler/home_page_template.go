package home_page_handler

import (
    "html/template"
)

func GetHomeTemplate(parsedTemplate *template.Template) {
    parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                    {{template "sidebar_template" .}}
                                                {{end}}`)

    parsedTemplate.New("sidebar_template").Parse(`{{ define "sidebar_template"}}
                                                        <div class="b-side-bar-container">
                                                            <h4>{{.SideBarTitle}}</h4>

                                                            {{range .Links}}
                                                                 <a href="{{.Url}}">{{.Label}}</a>
                                                            {{end}}
                                                        </div>
                                                    {{end}}`)
}
