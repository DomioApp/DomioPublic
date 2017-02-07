package login_page_handler

func GetLoginFormTemplate() string {
    loginFormTemplate := `
                            {{define "login_form_template"}}
                                {{with .PageData.LoginPageData}}
                                    <div class="b-login-form-container">
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
    return loginFormTemplate
}