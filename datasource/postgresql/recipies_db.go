package postgresql

import (
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type RecipiesDB struct {
	conn *pgx.Conn
	logger *zap.Logger
}

func NewRecipiesDB(conn *pgx.Conn, logger *zap.Logger) *RecipiesDB{
	return &RecipiesDB{
		conn: conn,
		logger: logger,
	}
}