package repository

const (
	Create           = `INSERT INTO goods (project_id, name, description, priority, removed, created_at) VALUES ($1, $2, $3, $4, $5, get_current_time()) RETURNING *`
	Update           = `UPDATE goods SET name = $2, description = COALESCE(NULLIF($3, ''), description) WHERE id = $1 RETURNING *`
	UpdatePriority   = `UPDATE goods SET priority = $1 WHERE id = $2 RETURNING *`
	Remove           = `UPDATE goods SET removed = true WHERE id = $1 RETURNING *`
	FindById         = `SELECT g.id, g.project_id, g.name, g.description, g.priority, g.removed, g.created_at FROM goods g LEFT JOIN projects p on p.id = g.project_id WHERE g.id= $1 and g.removed = false`
	ListWithoutLimit = `SELECT (SELECT COUNT(*) FROM goods WHERE removed = true) AS removed_count, id, project_id, name, priority, removed, description, created_at from goods ORDER BY created_at OFFSET $1`
	List             = `SELECT (SELECT COUNT(*) FROM goods WHERE removed = true) AS removed_count, (SELECT COUNT(*) FROM goods WHERE removed = true) AS removed, id, project_id, name, priority, removed, description, created_at from goods ORDER BY created_at OFFSET $1 LIMIT $2`
	Reprioritiize    = `UPDATE goods SET priority = priority + 1 WHERE priority < (SELECT MAX(priority) FROM goods);`
	GetReprioritiize = `SELECT id, priority from goods WHERE id != $1 ORDER BY priority`
)
