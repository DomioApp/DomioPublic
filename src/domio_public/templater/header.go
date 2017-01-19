package templater

import (
    "html/template"
    "bytes"
)

type Link struct {
    Url string
}
type PageData struct {
    Links []Link
}

func GetHeader() template.HTML {

    t := template.New("Index")

    output, _ := t.Parse(`
                            <div class="b-top-bar-container">
                                <div class="left-column">
                                    <ul>
                                        {{range .Links}}
                                            <li>
                                             <a href={{.Url}}>{{.Url}}</a>
                                            </li>
                                        {{end}}
                                    </ul>
                                </div>
                            </div>
                        `)

    pageData := PageData{
        Links:[]Link{
            Link{Url:"https://domio.in"},
            Link{Url:"https://google.com"},
        },
    }

    var doc bytes.Buffer
    output.Execute(&doc, pageData)

    return template.HTML(doc.String())
}