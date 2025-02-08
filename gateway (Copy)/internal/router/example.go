package router

import "server/internal/middleware"

func (m *routeStruct) Example() {
	subRoute := m.app.Group("/example")
	subRoute.Get("", m.handler.ExampleList)
	subRoute.Get("/:id", m.handler.ExampleDetail)
	subRoute.Use(middleware.AuthBearer())
	subRoute.Post("", m.handler.ExampleCreate)
	subRoute.Put("/:id", m.handler.ExamplePut)
	subRoute.Delete("/:id", m.handler.ExampleDelete)
}
