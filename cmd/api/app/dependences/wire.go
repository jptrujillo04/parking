package dependences

import (
	"log"
	HandlerUser "parking/cmd/api/handler/register"
	"parking/internal/register"
)

type HandlerContainer struct {
	HandlerUser HandlerUser.Handler
}

func NewWire() HandlerContainer {
	dep, err := NewDependencies()
	if err != nil {
		log.Panicf(err.Error())
		return HandlerContainer{}
	}
	useCaseUser := register.NewUserUseCase(dep.UserRepository)
	return HandlerContainer{
		HandlerUser: HandlerUser.NewHandler(useCaseUser),
	}
}
