package templater

import "html/template"

func GetBaseTemplate() *template.Template {
    baseTemplate := template.New("BaseTemplate")
    parsedBaseTemplate, _ := baseTemplate.Parse(getBaseTemplateContent())
    return parsedBaseTemplate
}

func getBaseTemplateContent() string {
    baseTemplateContent := `
                        {{define "BaseTemplate"}}
                            <!DOCTYPE html>

                            <html lang="en">

                            <head>
                                <meta charset="UTF-8">
                                <meta name="description" content="Domio is a marketplace for domains.">
                                <title>{{.PageTitle}}</title>
                                <link rel="stylesheet" href="/style.css" />
                                {{template "head" .}}
                            </head>

                            <body>
                                {{template "body" .}}
                                <script src="/bundle.js"></script>
                            </body>

                            </html>
                        {{end}}`
    return baseTemplateContent
}