package pokeapi

type PokeapiReponse struct {
	Count uint
	Next *string
	Previous *string
	Results []PokeapiResult
}

type PokeapiResult struct {
	Name string
	Url string
}
