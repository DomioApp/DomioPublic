package login_page_handler

type FormRow struct {
    Label string
    Name  string
    Type  string
}

type SubmitButton struct {
    Label string
}

func GetLoginFormTemplate() string {
    loginFormTemplate := `
                            {{define "login_form_template"}}
                                {{with .PageData.LoginPageData}}
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