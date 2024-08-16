package usecaseslogic

import (
	"context"

	"github.com/abhilasha336/thinkpalm/internal/repodb"
)

// ThinkpalmUseCase holds OauthRepoImply interface
type ThinkpalmUseCase struct {
	useCase repodb.ThinkpalmRepoImplements
}

// THinkpalmUsecaseImply which implements functions
type ThinkpalmUsecaseImplements interface {
	GetPartnerId(ctx context.Context, clientID, clientSecret string) (string, string, error)
}

// NewThinkpalmUseCase function assign values to THinkpalmUseCase
func NewThinkpalmUseCase(repo repodb.ThinkpalmRepoImplements) ThinkpalmUsecaseImplements {
	return &ThinkpalmUseCase{
		useCase: repo,
	}
}

// function helps to implement businuss logics with data
func (think *ThinkpalmUseCase) GetPartnerId(ctx context.Context, clientID, clientSecret string) (string, string, error) {
	return think.useCase.GetPartnerId(ctx, clientID, clientSecret)
}
