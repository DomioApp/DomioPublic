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
    DomainName   string
    FormRows     []FormRow
    SubmitButton SubmitButton
}

func GetRentDomainFormTemplate() string {
    addDomainFormTemplate := `
                            {{define "add_domain_form_template"}}
                                {{with .PageData.FormData}}
                                    <div class="b-add-domain-form-container">
                                        <form>
                                            <div class="b-form-rows-container"></div>

                                            <div>
                                                <button value="{{.DomainName}}">{{.SubmitButton.Label}}</button>
                                            </div>
                                        </form>
                                    </div>
                                {{end}}
                            {{end}}
                        `
    return addDomainFormTemplate
}
func GetFormData(domainName string) FormData {
    return FormData{
        DomainName:domainName,
        SubmitButton: SubmitButton{Label:"Rent now"},
    }
}