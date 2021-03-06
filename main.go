package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"

	"net/http"
	// "fmt"
	"database/sql"
	_ "github.com/lib/pq"
	// "github.com/coopernurse/gorp"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=devtalks-dev sslmode=disable")
	PanicIf(err)
	return db
}

func PanicIf(err error){
	if err != nil {
		panic(err)
	}
}

type Video struct {
	Id int64
	Title string
}

func main() {
	m := martini.Classic()
	m.Map(SetupDB())
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(martini.Static("assets"))

	m.Get("/", func (ren render.Render, r * http.Request, db *sql.DB) {
		rows, err := db.Query("SELECT * from videos")
		PanicIf(err)
		defer rows.Close()

		videos := []Video{}
		for rows.Next() {
			v := Video{}
			err := rows.Scan(&v.Id, &v.Title)
			PanicIf(err)
			videos = append(videos, v)
			// fmt.Fprintf(rw, "Title: %s", title)
		}

		ren.HTML(200, "videos", videos)
	})

	m.Run()
}

// type Page struct {
// 	Title string
// }

// type Response map[string]interface{}

// // convert the Response map to string
// func (r Response) String() (s string) {
// 	b, err := json.Marshal(r)
// 	if err != nil {
// 		s = ""
// 		return
// 	}
// 	s = string(b)
// 	return
// }

// // func viewHandler(w http.ResponseWriter, req *http.Request) {
// // 	p := &Page{Title: "title"}
// // 	t, _ := template.ParseFiles("views/app.html")
// // 	t.Execute(res, p)
// // }

// func homeHandler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("homeHandler")
// 	filePath := "public/index.html"
// 	content, err := ioutil.ReadFile(filePath)
// 	if err != nil {

// 	}
// 	fmt.Fprintf(w, string(content))
// 	return
// }

// func jsonHandler(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// cast to string and pass it to response
// 	fmt.Fprint(w, Response{"success": true, "message": "Hello!"})
// 	return
// }

// func staticHandler(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("staticHandler")
// 	http.ServeFile(w, req, req.URL.Path[1:])
// 	// http.FileServer(http.Dir("public"))
// }

// func main() {

// 	r := mux.NewRouter()

// 	// http.Handle("/", homeHandler)
// 	r.HandleFunc("/", homeHandler)
// 	r.HandleFunc("/api/v1/videos", jsonHandler)

// 	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

// 	http.Handle("/", r)

// 	err := http.ListenAndServe(":8090", nil)
// 	fmt.Println(err)
// }
