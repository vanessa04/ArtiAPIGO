package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/gorilla/handlers"


)

type  Article struct {
	Title     string   `json:"title, omitempty"`
	Image     string   `json: "image, omitempty"`
  Author    string   `json:"author, omitempty"`
	Paragrath string   `json:"paragrath, omitempty"`
}



var articles []Article



func GetArti(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(articles)
}



func main() {
	router := mux.NewRouter()
	articles = append(articles, Article{ Title: "This is An Article About Jon", Image: "http://www.whitelawandco.com/images/Jonathan.jpg", Author:"Vanessa Valoy #1", Paragrath:"This is a paragrath #1"})
	articles = append(articles, Article{ Title: "This is An Article About Solly", Image: "https://scontent-lga3-1.xx.fbcdn.net/v/t1.0-9/1377321_1380578442182164_769913331_n.jpg?_nc_cat=0&oh=60ea099ae3a6f01fdbfaf22a90f1cc1a&oe=5BD5719E", Author:"Vanessa Valoy #2", Paragrath:"This is a paragrath #2"})
  articles = append(articles, Article{ Title: "This is An Article About Sulaiman", Image: "https://media.licdn.com/dms/image/C4D03AQF_dCCLZf-cGA/profile-displayphoto-shrink_800_800/0?e=1536796800&v=beta&t=8-cdXUEOu2Rtf7XdclB6HP7preuXqGCjamW4oFhtBjQ", Author:"Vanessa Valoy #3", Paragrath:"This is a paragrath #3"})
	articles = append(articles, Article{ Title: "This is An Article About Yunise", Image: "https://scontent-lga3-1.xx.fbcdn.net/v/t1.0-9/10463892_870894956271462_6727171891539991069_n.jpg?_nc_cat=0&oh=6bac27658354b51a1370262e9faeb126&oe=5BD62FAC", Author:"Vanessa Valoy #3", Paragrath:"This is a paragrath #3"})


	router.HandleFunc("/articles", GetArti).Methods("GET")
  log.Print("localhost:8000")
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(router)))

}
