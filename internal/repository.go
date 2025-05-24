package internal

import (
	"database/sql"
	"fmt"
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
			fmt.Println(err)
			continue
		}
		res = append(res, t)
	}

	return res, nil
}

func (r repository) GetById(id int) (*task, error) {
	row := r.db.QueryRow("SELECT * FROM task WHERE id = $1;", id)

	var t task
	err := row.Scan(&t.Id, &t.Title, &t.Description, &t.Is_completed, &t.Created_at, &t.Updated_at)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r repository) DeleteById(id int) error {
	_, err := r.db.Exec("DELETE FROM task WHERE id = $1;", id)
	if err != nil {
		return err
	}
	return nil
}
