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

func (r *Repository) GetAll() ([]Vacancy, error) {
	query := `SELECT * FROM vacancies;`
	row, err := r.pool.Query(context.Background(), query)
	if err != nil {
		r.log.Err(err).Msg("Ошибка при получения вакансий с Базы vacancy.repository.GetAll()")
		return nil, err
	}
	defer row.Close()
	// Этот метод работает только со Struct Tags `db:"id"`
	return pgx.CollectRows(row, pgx.RowToStructByNameLax[Vacancy])

	// Старыый метод
	// var vacancies []Vacancy
	// for row.Next() {
	// 	var vacancy Vacancy
	// 	err := row.Scan(&vacancy.Id, &vacancy.Email, &vacancy.Name, &vacancy.Role, &vacancy.Type, &vacancy.Salary, &vacancy.Location, &vacancy.CreatedAt, &vacancy.UpdatedAt)
	// 	if err != nil {
	// 		r.log.Err(err).Msg("Ошибка при сканировании вакансии из Базы vacancy.repository.GetAll()")
	// 		return nil, err
	// 	}
	// 	vacancies = append(vacancies, vacancy)
	// }
	// return vacancies, nil
}
