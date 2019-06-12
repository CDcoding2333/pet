package department

type newDepartmentReq struct {
	Alias    string `json:"alias" mapstructure:"alias"`
	Brief    string `json:"brief" mapstructure:"brief"`
	ParentID uint   `json:"parent_id" mapstructure:"parent_id"`
	LogoURL  string `json:"logo_url" mapstructure:"logo_url"`
}
