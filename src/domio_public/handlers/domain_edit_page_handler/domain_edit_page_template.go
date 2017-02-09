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
                                                                            <p>Is visible: {{.IsVisible}}</p>
                                                                            <label>Visible:</label>

                                                                            {{if .IsVisible}}
                                                                                <input type="radio" value="visible" name="visible" checked>Yes</input>
                                                                                <input type="radio" value="invisible" name="visible">No</input>
                                                                            {{else}}
                                                                                <input type="radio" value="visible" name="visible">Yes</input>
                                                                                <input type="radio" value="invisible" name="visible" checked>No</input>
                                                                            {{end}}

                                                                        </div>

                                                                        <div>
                                                                            <label>Price Per Month, $:</label>
                                                                            <input name="price_per_month" type="number" value="{{.PricePerMonth}}"></input>
                                                                        </div>

                                                                        <div>Nameservers:</div>
                                                                        <div>
                                                                            <label>NS1: {{.NS1}}</label>
                                                                        </div>

                                                                        <div>
                                                                            <label>NS2: {{.NS2}}</label>
                                                                        </div>

                                                                        <div>
                                                                            <label>NS3: {{.NS3}}</label>
                                                                        </div>

                                                                        <div>
                                                                            <label>NS4: {{.NS4}}</label>
                                                                        </div>

                                                                        <button>Save</button>
                                                                        <button value="{{.Name}}" class="b-delete-domain-button">Delete</button>

                                                                    </div>
                                                                {{end}}
                                                            </form>
                                                         {{end}}`)
    if (err != nil) {
        log.Print(err)
    }
}