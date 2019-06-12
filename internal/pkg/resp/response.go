package resp

import (
	"net/http"

	"github.com/CDcoding2333/pet/internal/pkg/errs"
	"github.com/gin-gonic/gin"
)

// Response ...
type Response interface {
	Succ() gin.H
	Result(result interface{}) gin.H
}

type response struct {
}

const (
	//Code reply code
	Code = "code"
	//Message reply message
	Message = "message"
	//Data ...
	Data = "data"

	ok = iota
)

// NewResponse ...
func NewResponse() Response {
	return &response{}
}

func (r *response) Succ() gin.H {
	return gin.H{Code: ok, Message: "succ"}
}

// Result ...
func (r *response) Result(result interface{}) gin.H {
	switch result.(type) {
	case *errs.MError:
		err := result.(*errs.MError)
		return gin.H{
			Code:    err.Code,
			Message: err.Msg,
		}
	case error:
		return gin.H{
			Code:    http.StatusInternalServerError,
			Message: result.(error).Error(),
		}
	default:
		return gin.H{
			Code:    ok,
			Message: result,
		}
	}
}
