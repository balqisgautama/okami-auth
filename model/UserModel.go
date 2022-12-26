package model

import "database/sql"

type UserModel struct {
	UserID        sql.NullInt64  `json:"user_id"`
	Email         sql.NullString `json:"email"`
	Password      sql.NullString `json:"password"`
	ClientID      sql.NullString `json:"client_id"`
	Status        sql.NullInt16  `json:"status"`
	Sysadmin      sql.NullInt16  `json:"sysadmin"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	CreatedClient sql.NullString `json:"create_client"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
	UpdatedClient sql.NullString `json:"updated_client"`
	DeletedAt     sql.NullTime   `json:"deleted_at"`
	DeletedClient sql.NullString `json:"deleted_client"`
}
