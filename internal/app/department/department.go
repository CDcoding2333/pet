package department

import (
	"github.com/CDcoding2333/pet/internal/pkg/types"
)

// Department ...
type Department struct {
	ID       uint   `gorm:"primary_key" json:"id,omitempty"`
	Alias    string `json:"alias,omitempty" gorm:"not null"`
	ParentID uint   `json:"parent_id,omitempty"`
	Brief    string `json:"brief,omitempty"`
	LogoURL  string `json:"logo_url,omitempty"`
	Author   uint   `json:"author"`
	types.Model
}

func (s service) newDepartment(department *Department) error {
	return s.d.Create(department).Error
}

func (s service) delDepartments(ids ...uint) error {
	return s.d.Where("id in (?)", ids).Delete(Department{}).Error
}

func (s service) updateDepartment(id uint, params map[string]interface{}) error {

	updateParams := make(map[string]interface{}, 0)
	if v, ok := params["alias"]; ok {
		updateParams["alias"] = v
	}

	if v, ok := params["brief"]; ok {
		updateParams["brief"] = v
	}

	return s.d.Model(Department{}).Where("id = ?", id).Update(updateParams).Error
}
