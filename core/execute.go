package core

import (
	"github.com/christopherfranklin/gsql/compiler"
	"github.com/rs/zerolog/log"
)

func ExecuteStatement(statement compiler.Statement) {
	switch statement.StatementType {
	case compiler.INSERT:
		log.Info().Msg("This is where we would do an insert.")
	case compiler.SELECT:
		log.Info().Msg("This is where we would do a select.")
	}
}
