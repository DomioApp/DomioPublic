package templater

import "html/template"

func GetParsedTemplate(templateName string) (*template.Template, error) {
    t := template.New(templateName)

    return t.Parse(getTemplateAsString(templateName))
}

func getTemplateAsString(templateName string) string {
    if (templateName == "Index") {
        return "{{.LinksContainer}} <br> {{.UserName}}"
    }

    return "error"
}