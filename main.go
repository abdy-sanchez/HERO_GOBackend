package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Estructuras con la tada de los heros

type Hero struct {
	Name       string      `json:"name,omitempty"`
	Biography  *Biography  `json:"biography,omitempty"`
	PowerStats *PowerStats `json:"powerstats,omitempty"`
	Images     *Images     `json:"images,omitempty"`
}

type Biography struct {
	FullName string `json:"fullName,omitempty"`
}

type PowerStats struct {
	Intelligence int `json:"intelligence,omitempty"`
	Strength     int `json:"strength,omitempty"`
	Speed        int `json:"speed,omitempty"`
	Durability   int `json:"durability,omitempty"`
	Power        int `json:"power,omitempty"`
	Combat       int `json:"combat,omitempty"`
}

type Images struct {
	Xs string `json:"xs,omitempty"`
	Sm string `json:"sm,omitempty"`
	Md string `json:"md,omitempty"`
	Lg string `json:"lg,omitempty"`
}

//Array para guardar la info de los Heroes

var heroes []Hero

//Funcion para devolver el Heroe consultado

func GetHeroEndPoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	for _, item := range heroes {

		if item.Name == params["name"] {

			json.NewEncoder(w).Encode(item) //Se retorna el Personaje que coincide con el parametro enviado en la URL

			return

		}
	}

	//Si NO coincide, se retorna un objeto vacío -> json.NewEncoder(w).Encode(&Hero{})

	w.WriteHeader(http.StatusNotFound) //404 Not Found
	w.Write([]byte("No se encontró el heroe."))

}

//Funcion principal del programa

func main() {

	router := mux.NewRouter() //Declaración del enrutador

	//Se rellena la "Base de Datos" con los heroes

	heroes = append(heroes, Hero{Name: "Wolverine", Biography: &Biography{FullName: "John Logan"}, PowerStats: &PowerStats{Intelligence: 63, Strength: 32, Speed: 50, Durability: 100, Power: 89, Combat: 100}, Images: &Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg"}})
	heroes = append(heroes, Hero{Name: "Spider-Man", Biography: &Biography{FullName: "Peter Parker"}, PowerStats: &PowerStats{Intelligence: 90, Strength: 55, Speed: 67, Durability: 75, Power: 74, Combat: 85}, Images: &Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg"}})
	heroes = append(heroes, Hero{Name: "Iron Man", Biography: &Biography{FullName: "Tony Stark"}, PowerStats: &PowerStats{Intelligence: 100, Strength: 85, Speed: 58, Durability: 85, Power: 100, Combat: 64}, Images: &Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg"}})
	heroes = append(heroes, Hero{Name: "Black Widow", Biography: &Biography{FullName: "Natasha Romanoff"}, PowerStats: &PowerStats{Intelligence: 75, Strength: 13, Speed: 33, Durability: 30, Power: 36, Combat: 100}, Images: &Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg"}})
	heroes = append(heroes, Hero{Name: "Thor", Biography: &Biography{FullName: "Thor Odinson"}, PowerStats: &PowerStats{Intelligence: 69, Strength: 100, Speed: 83, Durability: 100, Power: 100, Combat: 100}, Images: &Images{Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg", Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg", Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg", Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg"}})

	//Endpoint API HEROES
	router.HandleFunc("/api/superhero", GetHeroEndPoint).Queries("hero", "{name}").Methods("GET")

	//activacion del servidor en el puerto correcto
	log.Fatal(http.ListenAndServe(":8080", router))

}
