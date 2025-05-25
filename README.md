## REST_TEST

# How to run
`docker compose up`

### API avaible in localhost:8080

1. `POST /tasks/` — создать задачу
2. `GET /tasks/` — получить список всех задач
3. `GET /tasks/{id}/` — получить задачу по ID
4. `PUT /tasks/{id}/` — обновить задачу
5. `DELETE /tasks/{id}/` — удалить задачу

Used: ***Golang, Gin, database/sql, github.com/lib/pq, PostgreSQL, Docker, Docker Compose***
