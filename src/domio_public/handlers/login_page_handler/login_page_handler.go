package login_page_handler

import (
    "net/http"
    "html/template"
)

type Person struct {
    UserName string
}

func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
    t := template.New("Login") // Create a template
    output, _ := t.Parse("{{.UserName}}")

    p := Person{UserName: "Login"}

    output.Execute(w, p)

}
