package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Directer *Directer `json:"directer"`


}

type Directer struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
var movies []Movie



func getMovies (w http.ResponseWriter , r *http.Request) {
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(movies)
defer r.Body.Close();

}

func deleteMovie (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/");
	params :=mux.Vars(r)
	fmt.Println(params)
	for index , item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
	defer r.Body.Close();
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
 _ = json.NewDecoder(r.Body).Decode(&movie)
 movie.ID =strconv.Itoa(rand.Intn(1000))
 movies=append(movies, movie)

 json.NewEncoder(w).Encode(movie)
 defer r.Body.Close();

}

func updateMovie (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index , val := range movies{
		if val.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie 

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID= params["id"]
			movies =append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	defer r.Body.Close();
}

func  getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params :=mux.Vars(r)

	for _ , movie := range movies{
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
		}
	}

	defer r.Body.Close();


}

func main(){
r :=mux.NewRouter()
movies = append(movies, Movie{
	ID: "1",Isbn: "6362",Title: "The Batman",Directer: &Directer{Firstname:"John",Lastname:"boy"}})

	movies = append(movies, Movie{
		ID: "2",Isbn: "63642",Title: "InterStaller",Directer: &Directer{Firstname:"John",Lastname:"boy"}})
		
r.HandleFunc("/movies",getMovies).Methods("GET")
r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
r.HandleFunc("/createMovie",createMovie).Methods("POST")
r.HandleFunc("/movies/{id}",updateMovie).Methods("POST")
r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

fmt.Println("Starting serever at 1000 Port ")

log.Fatal(http.ListenAndServe(":1000",r))

	
}