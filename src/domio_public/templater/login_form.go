package templater

import (
    "html/template"
    "bytes"
)

type FormRow struct {
    Label string
    Name  string
    Type  string
}

type SubmitButton struct {
    Label string
}

type LoginPageData struct {
    FormRows     []FormRow
    SubmitButton SubmitButton
}

func GetLoginForm() template.HTML {

    t := template.New("LoginForm")

    output, _ := t.Parse(`
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
                        `)

    pageData := LoginPageData{
        FormRows:[]FormRow{
            FormRow{Label:"Email", Name:"email", Type:"email"},
            FormRow{Label:"Password", Name:"password", Type:"password"},
        },
        SubmitButton: SubmitButton{Label:"Submit"},
    }

    var doc bytes.Buffer
    output.Execute(&doc, pageData)

    return template.HTML(doc.String())
}