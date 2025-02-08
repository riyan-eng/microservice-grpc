package router

func (m *routeStruct) Role() {
	subRoute := m.app.Group("/role")
	subRoute.GET("", m.handler.RoleAccessList)
	// subRoute.GET("/:id", m.handler.ExampleDetail)
	// subRoute.Use(middleware.AuthBearer())
	// subRoute.POST("", m.handler.ExampleCreate)
	// subRoute.PUT("/:id", m.handler.ExamplePut)
	// subRoute.DELETE("/:id", m.handler.ExampleDelete)
}
