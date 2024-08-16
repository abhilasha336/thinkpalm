package driver

import (
	"database/sql"
	"fmt"

	//_ "github.com/jackc/pgx/v4"
	"github.com/abhilasha336/thinkpalm/internal/constants"
	"github.com/abhilasha336/thinkpalm/internal/dstructures"
	// "github.com/jmoiron/sqlx"
)

// ConnectDB initializes postgres DB
func ConnectDB(cfg dstructures.Database) (*sql.DB, error) {
	datasource := prepareConnectionString(cfg)
	databaseType := constants.DatabaseType
	db, err := sql.Open(databaseType, datasource)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db: %s err: %s", datasource, err)
	}
	db.SetMaxOpenConns(cfg.MaxActive)
	db.SetMaxIdleConns(cfg.MaxIdle)
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db(ping): %s err: %s", datasource, err)
	}
	return db, nil
}

func prepareConnectionString(cfg dstructures.Database) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=20 search_path=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DATABASE, cfg.Schema)
}
