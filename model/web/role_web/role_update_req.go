package roleweb

type RoleUpdateReq struct {
	UID  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
}
