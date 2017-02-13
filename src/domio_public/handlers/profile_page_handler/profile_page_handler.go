package profile_page_handler

import (
    "html/template"
    "net/http"
    "domio_public/templater"
)

type PageData struct {
    PageTitle           string
    ProfileTopBarData   ProfileTopBarData
    ProfileMainAreaData ProfileMainAreaData
    SideBarData         templater.SideBarData
}

var profilePageTemplate *template.Template
var tokenCookie *http.Cookie

func init() {
    profilePageTemplate = templater.BuildTemplate(GetProfilePageTemplate)
}

func ProfilePageHandler(w http.ResponseWriter, req *http.Request) {
    var err error
    tokenCookie, err = req.Cookie("token")

    if (err != nil) {
        http.Redirect(w, req, "/login", http.StatusTemporaryRedirect)

    } else {
        templater.WriteTemplate(w, req, profilePageTemplate, GetPageName(), GetPageData())
    }

}

func GetUrl() string {
    return "/profile"
}

func GetPageName() string {
    return "ProfilePage"
}

func GetPageData() PageData {
    return PageData{
        PageTitle: "Domio - Marketplace for domains",
        ProfileTopBarData:GetProfileTopBarData(),
        ProfileMainAreaData:ProfileMainAreaData{Title:"Profile page"},
        SideBarData:templater.GetSideBarData(),
    }

}

func GetProfileTopBarData() ProfileTopBarData {

    return ProfileTopBarData{
        Links:[]templater.Link{
            {Url:"/profile/domains", Label:"My Domains"},
            {Url:"/profile/subscriptions", Label:"My Subscriptions"},
            {Url:"/profile/payments", Label:"Cards"},
        },
    }
}