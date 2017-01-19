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
                            <div class="b-links-list-container">
                                {{range .Links}}
                                    <a href={{.Url}}>{{.Url}}</a>
                                {{end}}
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