package department

import (
	"net/http"

	"github.com/CDcoding2333/pet/internal/pkg/resp"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/common/log"
)

// Service ...
type Service interface {
	HandlerNewDepartment(ctx *gin.Context)
}

type service struct {
	d *gorm.DB
	r resp.Response
}

// NewService ...
func NewService(d *gorm.DB, r resp.Response) (Service, error) {
	return service{
		d: d,
		r: r,
	}, nil
}

func (s service) HandlerNewDepartment(ctx *gin.Context) {
	var v newDepartmentReq
	if err := ctx.BindJSON(&v); err != nil {
		log.Error(err)
		ctx.AbortWithStatusJSON(http.StatusOK, s.r.Result(err))
		return
	}

	role := &Department{Alias: v.Alias, Brief: v.Brief, ParentID: v.ParentID, LogoURL: v.LogoURL}
	if err := s.newDepartment(role); err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, s.r.Result(err))
		return
	}

	ctx.JSON(http.StatusOK, s.r.Succ)
}
