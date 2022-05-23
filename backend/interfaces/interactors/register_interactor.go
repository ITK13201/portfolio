package interactors

import (
	"context"
	"github.com/ITK13201/portfolio/backend/domain"
	"github.com/ITK13201/portfolio/backend/ent"
	"github.com/ITK13201/portfolio/backend/internal/crypto"
)

type RegisterInteractor interface {
	CreateUser(registerVals domain.Register) (*ent.User, error)
}

type registerInteractor struct {
	sqlClient *ent.Client
}

func NewRegisterInteractor(sqlClient *ent.Client) RegisterInteractor {
	return &registerInteractor{
		sqlClient: sqlClient,
	}
}

func (r *registerInteractor) CreateUser(registerVals domain.Register) (*ent.User, error) {
	ctx := context.Background()
	hashedPassword := crypto.HashAndSalt(registerVals.Password)
	u, err := r.sqlClient.User.Create().SetUsername(registerVals.Username).SetHashedPassword(hashedPassword).Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}
