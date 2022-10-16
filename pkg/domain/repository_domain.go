package domain

type Column string

type CompanyDTO struct {
	ID          UUID         `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Employees   int64        `json:"employees"`
	Registered  bool         `json:"registered"`
	Type        CompanyType  `json:"company_type"`
	Status      EntityStatus `json:"status"`
}

type UpdateCompanyDTO struct {
	Description *string      `json:"description"`
	Employees   *int64       `json:"employees"`
	Registered  *bool        `json:"registered"`
	Type        *CompanyType `json:"company_type"`
}

type UserDTO struct {
	ID       UUID         `json:"id"`
	Name     string       `json:"name"`
	Password string       `json:"password"`
	Role     Role         `json:"role"`
	Email    string       `json:"email"`
	Status   EntityStatus `json:"status"`
}

type UpdateUserDTO struct {
	Name     *string `json:"name"`
	Password *string `json:"password"`
	Role     *Role   `json:"role"`
}
