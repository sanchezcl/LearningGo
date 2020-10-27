package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PokemonSpecies struct {
	Name string `json:"name"`
}

type Pokemon struct {
	EntryNo int `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type Response struct {
	Name string `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

func main() {
	fmt.Println("*** Consuming A RESTful API ***")
	url := "http://pokeapi.co/api/v2/pokedex/kanto/"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Printf("%v's Pokemons: \n", responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}
}
