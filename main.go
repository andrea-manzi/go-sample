package main
import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
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

  db := NewDB()


  r := httprouter.New()
  //r.GET("/", HomeHandler)
  // Posts collection
  r.GET("/posts", PostsIndexHandler)
  r.POST("/posts", PostsCreateHandler)
  // Posts singular
  r.GET("/posts/:id", PostShowHandler)
  r.PUT("/posts/:id", PostUpdateHandler)
  r.GET("/posts/:id/edit", PostEditHandler)

  r.Handler("GET", "/showbooks" ,  ShowBooks(db))

  r.HandlerFunc("POST", "/markdown", GenerateMarkdown)

  r.Handler("GET", "/", http.FileServer(http.Dir("public")))

  http.ListenAndServe(":"+port, r)
}


func ShowBooks(db *sql.DB) http.Handler {
  return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
    var title, author string
    err := db.QueryRow("select title, author from books").Scan(&title, &author)
    if err != nil {
      //insert object first
      stmt, _ := db.Prepare("INSERT INTO books  VALUES(?,?)")
      stmt.Exec("test","Andrea")
      db.QueryRow("select title, author from books").Scan(&title, &author)
    }
    fmt.Fprintf(rw, "The first book is '%s' by '%s'", title, author)
  })
}

func NewDB() *sql.DB {
  db, err := sql.Open("sqlite3", "example.sqlite")
  if err != nil {
  panic(err)
  }
  _, err = db.Exec("create table if not exists books(title text, author text)")
  if err != nil {
    panic(err)
  }
  return db
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
