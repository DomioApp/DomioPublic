package domain_edit_page_handler

import (
    "html/template"
    "log"
)

func GetDomainEditPageTemplate(parsedTemplate *template.Template) {
    _, err := parsedTemplate.New("main_template").Parse(`{{define "main_template"}}
                                                            <form>
                                                                {{with .PageData.DomainInfo}}
                                                                    <input name="domain_name" type="hidden" value="{{.Name}}"></input>
                                                                    <div data-domain="{{.Name}}">
                                                                        <h2>{{.Name}}</h2>

                                                                        <div>
                                                                            <label>Visible:</label>
                                                                            <input type="radio" name="visible">Yes</input>
                                                                            <input type="radio" name="visible">No</input>
                                                                        </div>

                                                                        <div>
                                                                            <label>Price Per Month, $:</label>
                                                                            <input name="price_per_month" type="number" value="{{.PricePerMonth}}"></input>
                                                                        </div>
                                                                        <button>Save</button>
                                                                    </div>
                                                                {{end}}
                                                            </form>
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }
}