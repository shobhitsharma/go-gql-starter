package resolvers

import (
	"github.com/shobhitsharma/go-gql-starter/db"
)

// Resolvers including query and mutation
type Resolvers struct {
	*db.DB
}
