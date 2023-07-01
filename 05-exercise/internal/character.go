package charactercli

type Characters struct {
	Info    Info        `json:"info"`
	Results []Character `json:"results"`
}

type Info struct {
	Count int `json:"count"`
	Pages int `json:"pages"`
	Next  int `json:"next"`
}

type Character struct {
	CharacterID int    `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Gender      string `json:"gender"`
}

type RickAndMortyRepo interface {
	GetCharacters() (Characters, error)
}

func NewCharacter(characterID int, name string, status string, gender string) (c Character) {
	c = Character{
		CharacterID: characterID,
		Name:        name,
		Status:      status,
		Gender:      gender,
	}
	return
}
