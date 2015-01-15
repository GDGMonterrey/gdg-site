package hello

import (
    "html/template"
    "net/http"
    //"path/filepath"
)

func init() {
    http.HandleFunc("/", home)
    http.HandleFunc("/posts", posts)
}

func home(w http.ResponseWriter, r *http.Request) {
  d := map[string]interface{}{"Titulo": "GDG Monterrey", "Tab" : 0}
  t := template.Must(template.ParseFiles("templates/root.html", "templates/base.html"))
  err := t.ExecuteTemplate(w,"base", d)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } 
}

type Site struct {
    Titulo string 
    Tab int
}
func posts(w http.ResponseWriter, r *http.Request) {
  d := Site {Titulo:"GGD Monterrey", Tab: 1}
  //  d := map[string]interface{}{"titulo": "GDG Monterrey", "tab" : 1}
  t := template.Must(template.ParseFiles("templates/posts.html", "templates/base.html"))
  err := t.ExecuteTemplate(w,"base", d)
      
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }                
   
}
