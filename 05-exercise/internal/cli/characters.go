package cli

import (
	"encoding/csv"
	"fmt"
	characterscli "github.com/CodelyTV/golang-examples/05-exercise/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

type CobraFn func(cmd *cobra.Command, args []string)

func InitCharactersCmd(repository characterscli.RickAndMortyRepo) *cobra.Command {
	charactersCmd := &cobra.Command{
		Use:   "characters",
		Short: "Print data about characters of Rick And Morty",
		Run:   runCharactersFn(repository),
	}

	return charactersCmd
}

func runCharactersFn(repository characterscli.RickAndMortyRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		characters, _ := repository.GetCharacters()

		fmt.Println(">> transfer to csv...")

		data := [][]string{
			{"ID", "Name", "Status", "Gender"},
		}

		for _, character := range characters.Results {
			data = append(data, []string{strconv.Itoa(character.CharacterID), character.Name, character.Status, character.Gender})
		}

		file, err := os.Create("characters.csv")
		if err != nil {
			log.Fatalf("Error al crear el archivo CSV: %s", err.Error())
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		writer.WriteAll(data)

		writer.Flush()
	}
}
