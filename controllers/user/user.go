package controller

import (
	"net/http"
	"user-service/common"
	"user-service/common/response"
	"user-service/domain/dto"
	service "user-service/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service service.IServiceRegistry
}
type IUserController interface {
	Login(*gin.Context)
	Register(*gin.Context)
	Update(*gin.Context)
	GetUserLogin(*gin.Context)
	GetUserByUUID(*gin.Context)
}

func NewUserController(service service.IServiceRegistry) IUserController {
	return &UserController{service: service}
}

func (u *UserController) Login(ctx *gin.Context) {
	req := &dto.LoginRequest{}

	err := ctx.ShouldBindJSON(req)

	if err != nil {
		response.HttpResponse(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := common.ErrValidationResponse(err)
		response.HttpResponse(response.ParamsHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Mesaage: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})

		return

	}
	user, err := u.service.GetUser().Login(ctx, req)
	if err != nil {
		response.HttpResponse(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamsHTTPResp{
		Code:  http.StatusOK,
		Data:  user.User,
		Token: &user.Token,
		Gin:   ctx,
	})
}

func (u *UserController) Register(ctx *gin.Context) {
	req := &dto.RegisterRequest{}

	err := ctx.ShouldBindJSON(req)

	if err != nil {
		response.HttpResponse(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := common.ErrValidationResponse(err)
		response.HttpResponse(response.ParamsHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Mesaage: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})

		return

	}
	user, err := u.service.GetUser().Register(ctx, req)
	if err != nil {
		response.HttpResponse(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamsHTTPResp{
		Code: http.StatusOK,
		Data: user.User,
		Gin:  ctx,
	})
}

func (u *UserController) Update(ctx *gin.Context) {
	req := &dto.UpdateRequest{}
	uuid := ctx.Param("uuid")

	err := ctx.ShouldBindJSON(req)

	if err != nil {
		response.HttpResponse(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := common.ErrValidationResponse(err)
		response.HttpResponse(response.ParamsHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Mesaage: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})

		return

	}
	user, err := u.service.GetUser().Update(ctx, req, uuid)
	if err != nil {
		response.HttpResponse(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamsHTTPResp{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
}

func (u *UserController) GetUserLogin(ctx *gin.Context) {
	user, err := u.service.GetUser().GetUserLogin(ctx.Request.Context())
	if err != nil {
		response.HttpResponse(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})

		return
	}
	response.HttpResponse(response.ParamsHTTPResp{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})

}

func (u *UserController) GetUserByUUID(ctx *gin.Context) {
	user, err := u.service.GetUser().GetUserByUUID(ctx.Request.Context(), ctx.Param("uuid"))
	if err != nil {
		response.HttpResponse(response.ParamsHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})

		return
	}
	response.HttpResponse(response.ParamsHTTPResp{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})

}
