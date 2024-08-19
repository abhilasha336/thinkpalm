package usecaseslogic

import (
	"context"

	"github.com/abhilasha336/thinkpalm/internal/dstructures"
	"github.com/abhilasha336/thinkpalm/internal/repodb"
)

// ThinkpalmUseCase holds OauthRepoImply interface
type ThinkpalmUseCase struct {
	useCase repodb.ThinkpalmRepoImplements
}

// THinkpalmUsecaseImply which implements functions
type ThinkpalmUsecaseImplements interface {
	RegisterUser(ctx context.Context, user dstructures.LoginRequest) error
	LoginUser(ctx context.Context, user dstructures.LoginRequest) error
}

// NewThinkpalmUseCase function assign values to THinkpalmUseCase
func NewThinkpalmUseCase(repo repodb.ThinkpalmRepoImplements) ThinkpalmUsecaseImplements {
	return &ThinkpalmUseCase{
		useCase: repo,
	}
}

// function helps to implement businuss logics with data
func (think *ThinkpalmUseCase) RegisterUser(ctx context.Context, user dstructures.LoginRequest) error {
	return think.useCase.RegisterUser(ctx, user)
}

func (think *ThinkpalmUseCase) LoginUser(ctx context.Context, user dstructures.LoginRequest) error {
	return think.useCase.LoginUser(ctx, user)
}
