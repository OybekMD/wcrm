package models

type Policy struct {
	Role     string `json:"role"`
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
}

type RbacAllRolesResp struct {
	Roles []string `json:"role"`
}

type ListPolePolicyResponse struct {
	Policies []*Policy `json:"policies"`
}

type CreateUserRoleRequest struct {
	Role string `json:"role"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type SuperAdminMessage struct {
	Message string `json:"message"`
}

type AddPolicyRequest struct {
	Policy struct {
		Role    string `json:"role"`
		EndPoint string `json:"endpoint"`
		Method  string `json:"method"`
	} `json:"policy"`
}


type ListRolePolicyResp struct {
	Policies []*Policy `json:"policies"`
}
