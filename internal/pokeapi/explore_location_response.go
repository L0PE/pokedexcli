package pokeapi

type ExploreLocationReponse struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon ExploreLocationPokemonResult
}

type ExploreLocationPokemonResult struct {
	Name string
	Url string
}
