package cli

import (
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/christopherfranklin/gsql/core"
	"github.com/marianogappa/sqlparser"
	"github.com/rs/zerolog/log"
)

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{}
	for _, s := range suggestionsMap {
		suggestions = append(suggestions, s...)
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func getExecutor(file string) func(string) {
	// TODO: Use file for database here
	t := &core.Table{}
	return func(s string) {
		s = strings.TrimSpace(s)
		s = strings.ToLower(s)

		switch s {
		case "":
			return
		case ".quit", ".exit":
			log.Info().Msg("Goodbye!")
			os.Exit(0)
		default:
			// TODO: Process commands here!
			if strings.HasPrefix(s, ".") {
				// Unrecognized Meta-Command
				log.Error().Msg(fmt.Sprintf("Unrecognized command '%s'", s))
				break
			}
			// TODO: Prepare Statement w/ compiler
			query, err := sqlparser.Parse(s)
			if err == nil {
				core.ExecuteStatement(query, t)
				log.Info().Msg(t.GetTableSize())
			} else {
				log.Error().Msg(fmt.Sprintf("Unrecognized keyword at start of '%s'", s))
			}
		}
	}
}

func StartPrompt(file string) {
	p := prompt.New(
		getExecutor(file),
		completer,
		prompt.OptionTitle("Welcome to GSQL! A lightweight SQL Database!"),
		prompt.OptionPrefix("sql > "),
	)
	p.Run()
}
