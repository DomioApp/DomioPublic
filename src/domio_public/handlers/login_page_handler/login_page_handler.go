package login_page_handler

import (
    "net/http"
    "domio_public/templates/login_page"
)

func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte(login_page.Show("John")))
}
/*
func LoginPageHandler(w http.ResponseWriter, req *http.Request) {

    var err error

    t := template.New("login_page.html")
    t, err = t.ParseFiles(path.Join(config.Config.TemplatesFolder, "login_page.html"))

    if (err != nil) {
        log.Print(err)
    }

    templateData := TemplateData{Title:"hello there"}
    t.Execute(w, templateData)

}
*/

/*
func IndexHandler(w http.ResponseWriter, req *http.Request) {

    var err error
    Template, err = ace.Load("index", "", &ace.Options{DynamicReload: true, BaseDir:config.Config.TemplatesFolder})

    if err != nil {
        log.Println(err)
    }

    if err := Template.Execute(w, map[string]string{"Msg": "Hello Domio"}); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}*/
