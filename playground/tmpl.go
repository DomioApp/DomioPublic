//http://stackoverflow.com/questions/41856021/how-to-parse-multiple-strings-into-a-template-with-go
//http://golang-examples.tumblr.com/post/87553422434/template-and-associated-templates

package main

import (
    "html/template"
    "os"
)

func main() {
    t1, _ := template.New("t1").Parse(`{{define "one"}}I'm #1.{{end}}`)
    t2, _ := t1.New("t2").Parse(`{{define "two"}}I'm #2, including #1: {{template "one" .}}{{end}}`)
    t2.ExecuteTemplate(os.Stdout, "two", nil)
}