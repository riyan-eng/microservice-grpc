package router

import "server/internal/middleware"

func (m *routeStruct) Example() {
	subRoute := m.app.Group("/example")
	subRoute.GET("", m.handler.ExampleList)
	subRoute.GET("/:id", m.handler.ExampleDetail)
	subRoute.Use(middleware.AuthBearer())
	subRoute.POST("", m.handler.ExampleCreate)
	subRoute.PUT("/:id", m.handler.ExamplePut)
	subRoute.DELETE("/:id", m.handler.ExampleDelete)
}
