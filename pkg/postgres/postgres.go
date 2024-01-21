package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/MihailPodgorny/bot_english_words/config"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// Return new Postgresql db instance
func NewPsqlDB(c *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.DBName,
		c.Postgres.Password,
	)

	db, err := sqlx.Connect(c.Postgres.PgDriver, dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Connect")
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "db.Ping")
	}

	return db, nil
}
