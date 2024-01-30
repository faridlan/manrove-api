package roleweb

type RoleUpdateReq struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty" validate:"required,gte=3"`
}
