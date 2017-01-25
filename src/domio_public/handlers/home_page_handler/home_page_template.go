package home_page_handler

import (
    "html/template"
    "domio_public/templater"
    "log"
    "bytes"
)

func GetTemplate() (*template.Template, error) {
    baseTemplate := templater.GetBaseTemplate()
    parsedHomeTemplate, parseErr := baseTemplate.Parse(getHomeTemplateContent())

    if (parseErr != nil) {
        log.Print(parseErr)
    }

    return parsedHomeTemplate, nil
}

func getHomeTemplateContent() string {
    homeTemplateContent := `
                            {{define "head"}}{{end}}

                            {{define "body"}}
                                {{.TopBar}}

                                <div class="b-main-container">
                                    {{.SideBar}}
                                    {{.MainArea}}
                                </div>
                            {{end}}

                            {{define "appstatus"}}{{.AppStatusInfo}}{{end}}
                        `
    return homeTemplateContent
}

func GetSideBar() template.HTML {
    return templater.GetSideBar()
}

type MainAreaData struct {
    Content template.HTML
}

func GetMainArea() template.HTML {

    t := template.New("MainArea")

    output, _ := t.Parse(`
                            <div class="b-main-area-container">
                                {{.Content}}
                            </div>
                        `)

    sideBarData := MainAreaData{
        Content:"main content here",
    }

    var doc bytes.Buffer
    output.Execute(&doc, sideBarData)

    return template.HTML(doc.String())
}