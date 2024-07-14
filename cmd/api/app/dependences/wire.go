package dependences

import "log"

type HandlerContainer struct {
	//HandlerLocation handlerLocation.Handler
}

func NewWire() HandlerContainer {
	_, err := NewDependencies()
	if err != nil {
		log.Panicf(err.Error())
		return HandlerContainer{}
	}
	//useCaseLocation := location.NewUseCaseLocation(dep.LocationRepository)
	return HandlerContainer{
		//HandlerLocation: handlerLocation.NewHandler(useCaseLocation),
	}
}
