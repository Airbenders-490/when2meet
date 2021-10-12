package domain

import "context"

// when2meet struct
type When2meet struct {
	W2M_ID   string `json:"id"` //uuid string
	W2M_name string `json:"name"`
}

// When2meetUseCase interface defines the functions all when2meetUseCases should have
type When2meetUseCase interface {
	Create(ctx context.Context, st *When2meet) error
	GetByID(ctx context.Context, id string) (*When2meet, error)
	Update(ctx context.Context, st *When2meet) error
	Delete(ctx context.Context, id string) error
}

// When2MeetRepository interface defines the functions all when2meetRepositories should have
type When2meetRepository interface {
	Create(ctx context.Context, id string, st *When2meet) error
	GetByID(ctx context.Context, id string) (*When2meet, error)
	Update(ctx context.Context, st *When2meet) error
	Delete(ctx context.Context, id string) error
}
