package controllers

import (
	controller "user-service/controllers/user"
	service "user-service/services"
)

type Registry struct {
	service service.IServiceRegistry
}

type IControllerRegistry interface {
	GetUserController() controller.IUserController
}

func NewControllerRegistry(service service.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (u *Registry) GetUserController() controller.IUserController {
	return controller.NewUserController(u.service)
}
