package index_handler

import (
    "net/http"
    "html/template"
    "domio_public/components/config"
    "path"
    "log"
)

var Template *template.Template

func IndexHandler(w http.ResponseWriter, req *http.Request) {

    var err error

    t := template.New("some template") // Create a template.
    t, err = t.ParseFiles(path.Join(config.Config.TemplatesFolder, "index.html"))  // Parse template file.
    if (err != nil) {
        log.Fatal(err)

    }
    /*user := GetUser() // Get current user infomration.*/
    t.Execute(w, "")  // merge.


}

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
