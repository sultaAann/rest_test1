package internal

import (
	"database/sql"
	"fmt"
	"time"
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

func (r repository) Create(t task) (int, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO task(title, description) VALUES ($1, $2) RETURNING id;",
		t.Title,
		t.Description).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, err
}

func (r repository) Update(id int, u task) (*task, error) {
	_, err := r.db.Exec(
		"UPDATE task SET title=$1, description=$2, is_completed=$3, updated_at=$4 WHERE id = $5",
		u.Title,
		u.Description,
		u.Is_completed,
		time.Now(),
		id,
	)

	if err != nil {
		return nil, err
	}

	return r.GetById(id)
}
