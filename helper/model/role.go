package model

import (
	"github.com/nostracode/mangrove-api/model/domain"
	roleweb "github.com/nostracode/mangrove-api/model/web/role_web"
)

func ToRoleResponse(role *domain.Role) *roleweb.RoleResponse {
	return &roleweb.RoleResponse{
		UID:       role.UID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func ToRoleResponses(roles []*domain.Role) []*roleweb.RoleResponse {
	roleResponses := []*roleweb.RoleResponse{}

	for _, role := range roles {
		roleResponses = append(roleResponses, ToRoleResponse(role))
	}

	return roleResponses
}
