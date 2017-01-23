package home_handler

import "domio_public/templater"

func GetTemplate() (string, error) {
    baseTemplate := templater.GetBaseTemplateAsString()
    return baseTemplate, nil
}