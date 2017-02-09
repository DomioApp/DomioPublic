package rent_domain_page_handler

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

func GetRentDomainFormTemplate() string {
    addDomainFormTemplate := `
                            {{define "add_domain_form_template"}}
                                {{with .PageData.FormData}}
                                    <div class="b-add-domain-form-container">
                                        <form>
                                            <div class="b-form-rows-container">
                                                        <select>
                                                            <option>1 Month</option>
                                                            <option>3 Months</option>
                                                            <option>6 Months</option>
                                                            <option>12 Months</option>
                                                        </select>
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
        SubmitButton: SubmitButton{Label:"Rent now"},
    }
}