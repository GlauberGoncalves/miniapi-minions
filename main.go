package main

import(
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"	
)

// config global
// **************************************************************
// crie sua conta no themoviedb https://developers.themoviedb.org
// e pegue sua apikey
// var APIKEY string = "SUACHAVE"
var APIKEY string = os.Getenv("themoviedb_key")
var URL_API string = "https://api.themoviedb.org/3/"
var ARGS string = "&include_adult=false&language=pt-BR"
// **************************************************************


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w , "Api Minions")
	fmt.Println("Endpoint: Principal")
}

func allFilmes(w http.ResponseWriter, r *http.Request){

	fmt.Println("entrou no allFilmes")
	response, err := http.Get( URL_API + "search/movie?api_key=" + APIKEY + ARGS + "&query=minions")
	if err != nil {
		fmt.Printf("Ops, deu um erro", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)		
		fmt.Fprintf( w, string(data))
	}

	fmt.Println(APIKEY)
}

// rotas
// ************************************************** //

func handleRequests(){
	
	http.HandleFunc("/", homePage)
	http.HandleFunc("/filmes", allFilmes)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequests()
}