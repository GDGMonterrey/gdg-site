package hello

import (
    "html/template"
    "net/http"
)

func init() {
    http.HandleFunc("/", home)
    http.HandleFunc("/posts", posts)
    http.HandleFunc("/single", single)
}

func home(w http.ResponseWriter, r *http.Request) {
  d := map[string]interface{}{"Titulo": "GDG Monterrey", "Tab" : 0}
  t, err := template.ParseFiles("templates/root.html", "templates/base.html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else{
    err := t.ExecuteTemplate(w,"base", d)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    } 
  }
  
}

type Site struct {
    Titulo string 
    Tab int
}

func posts(w http.ResponseWriter, r *http.Request) {
  d := Site {Titulo:"GDG Monterrey", Tab: 1}
  
  t, err := template.ParseFiles("templates/posts.html", "templates/base.html")
  
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else{
    err := t.ExecuteTemplate(w,"base", d)
      
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    } 

  }
     
}

func single(w http.ResponseWriter, r *http.Request) {
  d := Site {Titulo:"GDG Monterrey", Tab: 1}
  t, err := template.ParseFiles("templates/single.html"); 
  
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
  
      err := t.Execute(w, d)
      if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
      }
  }
}
