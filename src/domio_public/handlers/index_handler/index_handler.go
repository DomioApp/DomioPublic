package index_handler

import (
    "net/http"
    "html/template"
    "domio_public/components/config"
    "path"
)

var Template *template.Template

func IndexHandler(w http.ResponseWriter, req *http.Request) {

    t := template.New("some template") // Create a template.
    t, _ = t.ParseFiles(path.Join(config.Config.TemplatesFolder, "index.html"))  // Parse template file.
    /*user := GetUser() // Get current user infomration.*/
    t.Execute(w, nil)  // merge.


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
