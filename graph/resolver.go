//go:generate go run github.com/99designs/gqlgen generate

package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/NICKNAME-wengreen/BigDemo/graph/model"
)

type Resolver struct{
	todos []*model.Todo
}
