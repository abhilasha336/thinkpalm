package repodb

import (
	"context"
	"database/sql"

	"github.com/abhilasha336/thinkpalm/internal/dstructures"
	"github.com/sirupsen/logrus"
)

// ThinkpalmRepo holds db and config
type ThinkpalmRepo struct {
	repo *sql.DB
	cfg  *dstructures.EnvConfig
}

// THinkpalmRepoImply which implements functions
type ThinkpalmRepoImplements interface {
	GetPartnerId(ctx context.Context, clientID, clientSecret string) (string, string, error)
}

// NewTHinkpalmRepo used to assign values to both database and config
func NewThinkpalmRepo(repo *sql.DB, cfg *dstructures.EnvConfig) ThinkpalmRepoImplements {
	return &ThinkpalmRepo{
		repo: repo,
		cfg:  cfg,
	}
}

// fn which retrieves partnerid with client id and client secret
func (think *ThinkpalmRepo) GetPartnerId(ctx context.Context, clientID, clientSecret string) (string, string, error) {

	var (
		partnerID, redirectUri string
		err                    error
	)

	GetCredentials := `
	SELECT 
		partner_id,redirect_uri 
	FROM partner_api_credential
	WHERE
	client_id=$1
	AND
	client_secret=$2
	`
	row := think.repo.QueryRowContext(
		ctx,
		GetCredentials,
		clientID,
		clientSecret,
	)

	err = row.Scan(
		&partnerID,
		&redirectUri,
	)
	if err != nil {
		logrus.Errorf("GetPartnerId- scan error:%v", err)
		return "", "", err
	}

	return partnerID, redirectUri, nil

}
