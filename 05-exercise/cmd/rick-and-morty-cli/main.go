package main

import (
	characterscli "github.com/CodelyTV/golang-examples/05-exercise/internal"
	"github.com/CodelyTV/golang-examples/05-exercise/internal/cli"
	"github.com/CodelyTV/golang-examples/05-exercise/internal/storage/api"
	"github.com/spf13/cobra"
)

func main() {
	var repo characterscli.RickAndMortyRepo
	repo = api.NewApiRepository()

	rootCmd := &cobra.Command{Use: "characters-cli"}
	rootCmd.AddCommand(cli.InitCharactersCmd(repo))
	rootCmd.Execute()
}
