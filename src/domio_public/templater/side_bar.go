package templater

import (
    "html/template"
    "bytes"
)

type SideBarData struct {
    Title string
    Links []Link
}

func GetSideBar() template.HTML {

    t := template.New("SideBar")

    output, _ := t.Parse(`
                            <div class="b-side-bar-container">
                                <h4>{{.Title}}</h4>

                                {{range .Links}}
                                     <a href="{{.Url}}">{{.Label}}</a>
                                {{end}}
                            </div>
                        `)

    sideBarData := SideBarData{
        Title: "Categories",
        Links:[]Link{
            Link{Url:"/domains/business", Label:"Business"},
            Link{Url:"/domains/social", Label:"Social"},
            Link{Url:"/domains/cars", Label:"Cars"},
        },
    }

    var doc bytes.Buffer
    output.Execute(&doc, sideBarData)

    return template.HTML(doc.String())
}