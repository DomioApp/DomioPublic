package signup_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
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

type PageData struct {
    PageTitle     string
    LoginPageData LoginPageData
}

var signupPageTemplate *template.Template

func init() {
    signupPageTemplate = templater.BuildTemplate(GetSignupPageTemplate)
}

func SignupPageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, req, signupPageTemplate, GetPageName(), GetPageData())

}

func GetUrl() string {
    return "/signup"
}

func GetPageName() string {
    return "SignupPage"
}

func GetPageData() PageData {

    pageData := PageData{
        PageTitle: "Domio - Marketplace for domains. Signup to start earning by renting domains.",

        LoginPageData:LoginPageData{
            FormRows:[]FormRow{
                {Label:"Email", Name:"email", Type:"email"},
                {Label:"Password", Name:"password", Type:"password"},
            },
            SubmitButton: SubmitButton{Label:"Submit"},
        },
    }
    return pageData
}
