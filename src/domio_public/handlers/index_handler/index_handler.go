package index_handler

import (
    "net/http"
    "html/template"
    //"domio_public/templates"
)

var Template *template.Template

type TemplateData struct {
    Title string
}

func IndexHandler(w http.ResponseWriter, req *http.Request) {
    //w.Write([]byte(templates.Hello("John")))

}

/*
func IndexHandler(w http.ResponseWriter, req *http.Request) {

    var err error

    t := template.New("index.html")
    t, err = t.ParseFiles(path.Join(config.Config.TemplatesFolder, "index.html"))

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
