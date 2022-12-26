package model

import "database/sql"

type ResourceModel struct {
	ResourceID    sql.NullInt64  `json:"resource_id"`
	Surname       sql.NullString `json:"surname"`
	Nickname      sql.NullString `json:"nickname"`
	ClientID      sql.NullString `json:"client_id"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	CreateClient  sql.NullString `json:"create_client"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	UpdatedClient sql.NullString `json:"updated_client"`
	DeletedAt     sql.NullTime   `json:"deleted_at"`
	DeletedClient sql.NullString `json:"deleted_client"`
}
