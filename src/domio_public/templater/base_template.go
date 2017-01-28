package templater

import "html/template"

func GetBaseTemplate() *template.Template {
    baseTemplate := template.New("BaseTemplate")
    parsedBaseTemplate, _ := baseTemplate.Parse(getBaseTemplateContent())
    return parsedBaseTemplate
}

func getBaseTemplateContent() string {
    baseTemplateContent := `
                        {{define "base_template"}}
                            <!DOCTYPE html>

                            <html lang="en">

                            <head>
                                <meta charset="UTF-8">
                                <meta name="description" content="Domio is a marketplace for domains.">
                                <meta name="page" content="{{.PageName}}">

                                <title>{{.PageTitle}}</title>
                                <link rel="stylesheet" href="/style.css" />
                            </head>

                            <body>
                                {{template "main_template" .}}
                                <script src="/bundle.js"></script>
                            </body>

                            </html>
                        {{end}}`
    return baseTemplateContent
}