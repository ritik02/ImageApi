package main

import (
  "html/template"
  "net/http"
)

type Message struct {
     Title  string
     Body string
}


func HandleError(w http.ResponseWriter, errormessage string) {
     message := Message{Title: "Error Message", Body: errormessage}
     tmpl,_ := template.ParseFiles("templates/message.html")
     tmpl.Execute(w, message)
}
