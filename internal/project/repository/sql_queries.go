package repository

const (
	Create = `INSERT INTO projects (name) VALUES ($1) RETURNING *`
)
