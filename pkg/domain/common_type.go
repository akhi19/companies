package domain

type UUID string

type EntityStatus string

const (
	EntityStatusActive   EntityStatus = "active"
	EntityStatusInactive EntityStatus = "inactive"
)

type CompanyType string

const (
	CompanyTypeCorporation        CompanyType = "corporations"
	CompanyTypeNonProfit          CompanyType = "non_profit"
	CompanyTypeCooperative        CompanyType = "cooperative"
	CompanyTypeSoleProprietorship CompanyType = "sole_proprietorship"
)

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleViewer Role = "viewer"
	RoleAny    Role = "any"
)
