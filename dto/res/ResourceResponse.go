package res

type ResourceDetail struct {
	ResourceID    interface{} `json:"resource_id"`
	Surname       string      `json:"surname"`
	Nickname      string      `json:"nickname"`
	ClientID      string      `json:"client_id"`
	CreatedAt     int64       `json:"created_at"`
	CreateClient  string      `json:"create_client"`
	UpdatedAt     int64       `json:"updated_at"`
	UpdatedClient string      `json:"updated_client"`
	DeletedAt     int64       `json:"deleted_at"`
	DeletedClient string      `json:"deleted_client"`
}
