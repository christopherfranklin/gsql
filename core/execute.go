package core

import (
	"github.com/marianogappa/sqlparser/query"
	"github.com/rs/zerolog/log"
)

func ExecuteStatement(q query.Query, t *Table) {
	switch q.Type {
	case query.Insert:
		res := t.ExecuteInsert(q)
		if res == TableFull {
			log.Error().Msg("Table Full")
		} else if res == ExecuteFailed {
			log.Error().Msg("Error inserting row")
		} else {
			log.Info().Msg("Executed.")
		}
	case query.Select:
		log.Info().Msg("This is where we would do a select.")
	}
}
