package response

import (
	"net/http"
	"user-service/constants"
	ErrConstant "user-service/constants/error"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string `json:"status"`
	Mesaage string `json:"message"`
	Data    interface{}
	Token   *string // kenapa pointer ?
}

type ParamsHTTPResp struct {
	Code    int
	Err     error
	Mesaage *string
	Gin     *gin.Context
	Data    interface{}
	Token   *string
}

func HttpResponse(param ParamsHTTPResp) {
	param.Gin.JSON(param.Code, Response{
		Status:  constants.Success,
		Mesaage: http.StatusText(http.StatusOK),
		Data:    param.Data,
		Token:   param.Token,
	})

	message := ErrConstant.ErrInternalServerError.Error()

	if param.Mesaage != nil {
		message = *param.Mesaage
	} else if param.Err != nil {
		if ErrConstant.ErrorMapping(param.Err) {
			message = param.Err.Error()
		}
		param.Gin.JSON(param.Code, Response{
			Status:  constants.Error,
			Mesaage: message,
			Data:    param.Data,
		})

		return
	}

}
