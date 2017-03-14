package add_domain_page_handler

type FormRow struct {
    Label string
    Name  string
    Type  string
}

type SubmitButton struct {
    Label string
}

type FormData struct {
    FormRows     []FormRow
    SubmitButton SubmitButton
}

func GetAddDomainFormTemplate() string {
    addDomainFormTemplate := `
                            {{define "add_domain_form_template"}}
                                {{with .PageData.FormData}}
                                    <div class="b-add-domain-form-container">
                                        <form>

                                            <div class="b-form-rows-container">
                                                {{range .FormRows}}
                                                     <div class="b-form-row">
                                                        <label>{{.Label}}</label>
                                                        <input type="{{.Type}}" autocomplete="off" name="{{.Name}}"></input>
                                                     </div>
                                                {{end}}
                                            </div>

                                            <div>
                                                {{with .SubmitButton}}
                                                <input type="submit" value="{{.Label}}"/>
                                                {{end}}
                                            </div>

                                        </form>
                                    </div>
                                {{end}}
                            {{end}}
                        `
    return addDomainFormTemplate
}
func GetFormData() FormData {
    return FormData{
        FormRows:[]FormRow{
            {Label:"Name:", Name:"name", Type:"text"},
            {Label:"Price:", Name:"price_per_month", Type:"number"},
            {Label:"Tags:", Name:"tags", Type:"string"},
        },
        SubmitButton: SubmitButton{Label:"Add"},
    }
}