package logout_page_handler

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

var loginPageTemplate *template.Template

func init() {
    loginPageTemplate = templater.BuildTemplate(GetLogoutPageTemplate)
}

func LogoutPageHandler(w http.ResponseWriter, req *http.Request) {

    templater.WriteTemplate(w, loginPageTemplate, GetPageName(), GetPageData())

}

func GetUrl() string {
    return "/logout"
}

func GetPageName() string {
    return "LogoutPage"
}

func GetPageData() PageData {

    pageData := PageData{
        PageTitle: "Domio - Marketplace for domains. You've been logged out.",

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
