package api

import (
	"encoding/json"
	"fmt"
	characterscli "github.com/CodelyTV/golang-examples/05-exercise/internal"
	"io/ioutil"
	"net/http"
)

const (
	charactersEndpoint = "/character"
	apiURL             = "https://rickandmortyapi.com/api"
)

type rickAndMortyRepo struct {
	url string
}

func NewApiRepository() characterscli.RickAndMortyRepo {
	return &rickAndMortyRepo{url: apiURL}
}

func (rm *rickAndMortyRepo) GetCharacters() (Characters characterscli.Characters, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", rm.url, charactersEndpoint))
	contents, err := ioutil.ReadAll(response.Body)
	json.Unmarshal(contents, &Characters)
	return
}
