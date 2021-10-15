package domain

import "context"

// availibilities struct
type Availibilities struct {
	AID   string `json:"AID"` //uuid string
	SID   string `json:"SID"`
	ADAY  string `json:"ADAY"`
	Times [2]int `json:"Times"`
}

// availibilitiesUseCase interface defines the functions all availibilitiesUseCase should have
type AvailibilitiesUseCase interface {
	Create(ctx context.Context, st *Availibilities) error
	GetByID(ctx context.Context, id string) (*Availibilities, error)
	Update(ctx context.Context, st *Availibilities) error
	Delete(ctx context.Context, id string) error
}

// AvailibilitiesRepository interface defines the functions all AvailibilitiesRepository should have
type AvailibilitiesRepository interface {
	Create(ctx context.Context, id string, st *Availibilities) error
	GetByID(ctx context.Context, id string) (*Availibilities, error)
	Update(ctx context.Context, st *Availibilities) error
	Delete(ctx context.Context, id string) error
}
