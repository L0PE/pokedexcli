package pokeapi

type Pokemon struct {
	BaseExperience int64 `json:"base_experience"`
	Name 	 string
	Weight int64
	Height int64
	Types  []PokemonType
	Stats  []PokemonStat
}

type PokemonType struct {
	Slot int64
	Type struct{
		Name string
	}
}

type PokemonStat struct {
	Value int64 `json:"base_stat"`
	Stat 	struct{
		Name string
	}
}
