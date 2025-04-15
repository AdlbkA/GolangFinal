package connections

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang-auth-service/internal/app/config"
)

type Connections struct {
	DB *sqlx.DB
}

func (c *Connections) Close() {
	if c.DB != nil {
		_ = c.DB.Close()
	}
}

func New(cfg *config.Config) (*Connections, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database, cfg.DB.SSLMode))
	if err != nil {
		return nil, err
	}
	return &Connections{
		DB: db,
	}, nil
}
