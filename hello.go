package hello

import (
    "fmt"
    "html/template"
    "net/http"
    "path/filepath"
)

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, guestbookForm)
}

const guestbookForm = `
<html>
  <body>
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`
type Person struct {
    Name string //exported field since it begins with a capital letter
}
func sign(w http.ResponseWriter, r *http.Request) {
    //p := Person{Name:"Mary"}
    d := map[string]interface{}{"Name": "Gold"}
    if t, err := template.ParseFiles(filepath.Join("", "root.html")); err != nil {
    // Something gnarly happened.
      http.Error(w, err.Error(), http.StatusInternalServerError)
    } else {
    // return to client via t.Execute
      err := t.Execute(w, d)
      if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
      }
    }
   /* err := signTemplate.Execute(w, d)
    //err := signTemplate.Execute(w, r.FormValue("content"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }*/
}

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

const signTemplateHTML = `
<html>
  <body>
    <p>You wrote:</p>
    <pre>{{.Name}}</pre>
  </body>
</html>
`