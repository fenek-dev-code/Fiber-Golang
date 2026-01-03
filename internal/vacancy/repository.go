package vacancy

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Repository struct {
	log  *zerolog.Logger
	pool *pgxpool.Pool
}

func NewRepository(log *zerolog.Logger, pool *pgxpool.Pool) *Repository {
	return &Repository{log: log, pool: pool}
}

func (r *Repository) CreateVacancy(form *VacancyCreateForm) error {
	query := `INSERT INTO vacancies (email, name, role, type, salary, location)
	VALUES (@email, @name, @role, @type, @salary, @location)`
	args := pgx.NamedArgs{
		"email":    form.Email,
		"name":     form.Name,
		"role":     form.Role,
		"type":     form.Type,
		"salary":   form.Salary,
		"location": form.Location,
	}
	_, err := r.pool.Exec(context.Background(), query, args)
	return err
}
