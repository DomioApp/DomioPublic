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
    loginFormTemplate := `
                            {{define "add_domain_form_template"}}
                                {{with .PageData.FormData}}
                                    <div class="b-login-form-container">
                                        <form>
                                            <div class="b-form-rows-container">
                                                {{range .FormRows}}
                                                     <div>
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
    return loginFormTemplate
}
func GetFormData() FormData {
    return FormData{
        FormRows:[]FormRow{
            {Label:"Name", Name:"name", Type:"text"},
            {Label:"Price", Name:"price", Type:"text"},
        },
        SubmitButton: SubmitButton{Label:"Add"},
    }
}