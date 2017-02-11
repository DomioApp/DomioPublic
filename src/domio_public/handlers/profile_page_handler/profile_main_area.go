package profile_page_handler

type ProfileMainAreaData struct {
    Title string
}

func GetProfileMainAreaTemplate() string {
    return `{{ define "profile_main_area_template"}}

                {{with .PageData.ProfileMainAreaData}}

                    <div class="b-profile-main-area-container">
                        <button class="b-delete-account-button">Delete my account</button>
                    </div>

                {{end}}

            {{end}}
        `
}