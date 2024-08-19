package repodb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/abhilasha336/thinkpalm/internal/dstructures"
)

// ThinkpalmRepo holds db and config
type ThinkpalmRepo struct {
	repo *sql.DB
	cfg  *dstructures.EnvConfig
}

// THinkpalmRepoImply which implements functions
type ThinkpalmRepoImplements interface {
	RegisterUser(ctx context.Context, user dstructures.LoginRequest) error
	LoginUser(ctx context.Context, user dstructures.LoginRequest) error
}

// NewTHinkpalmRepo used to assign values to both database and config
func NewThinkpalmRepo(repo *sql.DB, cfg *dstructures.EnvConfig) ThinkpalmRepoImplements {
	return &ThinkpalmRepo{
		repo: repo,
		cfg:  cfg,
	}
}

// fn which retrieves partnerid with client id and client secret
func (think *ThinkpalmRepo) RegisterUser(ctx context.Context, user dstructures.LoginRequest) error {

	_, err := think.repo.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return errors.New("repo user insertion failed::" + err.Error())
	}

	return nil

}

func (think *ThinkpalmRepo) LoginUser(ctx context.Context, user dstructures.LoginRequest) error {
	var check dstructures.LoginRequestCheck

	rows := think.repo.QueryRow("select * from users where username=$1 and password=$2", user.Username, user.Password)
	err := rows.Scan(&check.Id, &check.Username, &check.Password)
	if check.Id == 0 {
		return errors.New("repo login error" + err.Error())
	}
	return nil

}
