package roleweb

type RoleCreateReq struct {
	Name string `json:"name,omitempty" validate:"required,gte=3"`
}
