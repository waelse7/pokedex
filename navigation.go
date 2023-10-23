package main

import "github.com/waelse7/pokedex/api"

type currentPage struct {
	Next *string
	Prev *string
	caughtPokemon map[string]*api.Pokemon
}

var page = currentPage{
	caughtPokemon: make(map[string]*api.Pokemon),
}

func SetNextPage(url *string){
	page.Next = url
}

func GetNextPage() *string {
	return page.Next
}

func SetPrevPage(url *string){
	page.Prev = url
}

func GetPrevPage() *string {
	return page.Prev
}

func SetCaughtPokemon(pokemon *api.Pokemon){
	page.caughtPokemon[pokemon.Name] = pokemon
}
func GetCaughtPokemon() map[string]*api.Pokemon {
	return page.caughtPokemon
}