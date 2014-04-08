package main

import (
	"github.com/codegangsta/martini"
	"net/http"
	"fmt"
	"github.com/lib/pq"
)

func main() {
	m := martini.Classic()
	m.Get("/", func (rw.= http.ResponseWriter, r * http.Request) {
		rw.Write([]byte("Hello world"))
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
