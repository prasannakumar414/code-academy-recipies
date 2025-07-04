package postgresql

import (
	"context"

	"com.int.recipies/models"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type RecipiesDB struct {
	conn   *pgx.Conn
	logger *zap.Logger
}

func NewRecipiesDB(conn *pgx.Conn, logger *zap.Logger) *RecipiesDB {
	return &RecipiesDB{
		conn:   conn,
		logger: logger,
	}
}

func (db *RecipiesDB) CreateRecipie(ctx context.Context, recipie models.Recipie) error {
	query := "INSERT INTO recipies (id, name, description, instructions, prep_time) VALUES ($1,$2,$3,$4,$5)"
	_, err := db.conn.Exec(ctx, query,
		recipie.ID,
		recipie.Name,
		recipie.Description,
		recipie.Instructions,
		recipie.PrepTime,
	)
	return err
}

func (db *RecipiesDB) GetRecipie(ctx context.Context, id int) (*models.Recipie, error) {
	query := "SELECT id, name, description, instructions, prep_time FROM recipies WHERE id = $1"
	rows, err := db.conn.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	var recipie models.Recipie
	for rows.Next() {
		err = rows.Scan(
			&recipie.ID,
			&recipie.Name,
			&recipie.Description,
			&recipie.Instructions,
			&recipie.PrepTime,
		)
		if err != nil {
			return nil, err
		}
	}
	return &recipie, nil
}

func (db *RecipiesDB) DeleteRecipie(ctx context.Context, id int) (error) {
	return nil
}
func (db *RecipiesDB)	UpdateRecipie(ctx context.Context, recipie models.Recipie) (error) {
	return nil
}
func (db *RecipiesDB)	GetRecipiesCount(ctx context.Context) (int, error) {
	return 0, nil
}
func (db *RecipiesDB)	SearchRecipies(ctx context.Context, query string, skip int, limit int) ([]models.Recipie, error) {
	return []models.Recipie{},nil
}
func (db *RecipiesDB)	ListRecipies(ctx context.Context, skip int, limit int) ([]models.Recipie, error) {
	return []models.Recipie{}, nil
}