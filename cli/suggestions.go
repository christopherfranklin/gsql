package cli

import (
	prompt "github.com/c-bata/go-prompt"
)

type suggestionType int

const (
	Commands suggestionType = iota
	Keywords
)

var commandSuggestions = []prompt.Suggest{
	{Text: ".quit", Description: "Quit/Exit application"},
	{Text: ".exit", Description: "Quit/Exit application"},
}

var keywordSuggestions = []prompt.Suggest{
	{Text: "select", Description: "Select statement to read data from a table"},
	{Text: "insert", Description: "Insert statement to add data to a table"},
}

var suggestionsMap = map[suggestionType][]prompt.Suggest{
	Commands: commandSuggestions,
	Keywords: keywordSuggestions,
}
