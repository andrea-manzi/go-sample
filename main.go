package main
import (
  "encoding/json"
  "fmt"
  "net/http"
  "os"
  "github.com/russross/blackfriday"
  "github.com/julienschmidt/httprouter"
)

type Book struct {
  Title string `json:"title"`
  Author string `json:"author"`
}


func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  r := httprouter.New()
  //r.GET("/", HomeHandler)
  // Posts collection
  r.GET("/posts", PostsIndexHandler)
  r.POST("/posts", PostsCreateHandler)
  // Posts singular
  r.GET("/posts/:id", PostShowHandler)
  r.PUT("/posts/:id", PostUpdateHandler)
  r.GET("/posts/:id/edit", PostEditHandler)

  r.HandlerFunc("GET", "/showbooks" , ShowBooks)

  r.HandlerFunc("POST", "/markdown", GenerateMarkdown)

  r.Handler("GET", "/", http.FileServer(http.Dir("public")))

  http.ListenAndServe(":"+port, r)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
  book := Book{"Building Web Apps with Go", "Jeremy Saenz"}
  js, err := json.Marshal(book)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}


func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
  markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
  rw.Write(markdown)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintln(rw, "Home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintln(rw, "posts index")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintln(rw, "posts create")
}
func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  id := p.ByName("id")
  fmt.Fprintln(rw, "showing post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintln(rw, "post update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintln(rw, "post delete")
}
func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintln(rw, "post edit")
}
