package internal

import (
	"database/sql"
)

type Repository interface {
	GetAll()
	Create()
	GetById()
	Update()
	Delete()
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r repository) GetAll() ([]task, error) {
	rows, err := r.db.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []task{}

	for rows.Next() {
		var t task
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Is_completed, &t.Created_at, &t.Updated_at)
		if err != nil {
			return nil, err
		}
		res = append(res, t)
	}

	return res, nil
}
