package comment

import (
	"context"
	"fmt"
)

// Commment - a representation of the comment
// structure for our services
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Services struct{
	Store Store
}

// NewServices - returns a pointer to a new
// services
func NewServices(store Store) *Services {
	return &Services{
		Store: store,
	}
}

func(s *Services) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}