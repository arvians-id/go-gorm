package request

type ChangeRolesRequest struct {
	UserID uint64   `json:"user_id"`
	RoleID []uint64 `json:"role_id"`
}

type PaginationData struct {
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
	TotalPage    int `json:"total_page"`
	TotalData    int `json:"total_data"`
	CurrentPage  int `json:"current_page"`
}
