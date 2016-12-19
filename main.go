package main
import (
  "fmt"
  "net/http"
  "os"
  "github.com/russross/blackfriday"
  "github.com/julienschmidt/httprouter"
)
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

  r.HandlerFunc("POST", "/markdown", GenerateMarkdown)
  r.Handler("GET", "/", http.FileServer(http.Dir("public")))
  http.ListenAndServe(":"+port, r)
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
