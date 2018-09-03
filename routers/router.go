package routers

import (
	"SHUPT-GO/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Post("/auth/login", controllers.AuthHandler)
	beego.Any("/graphql", controllers.GraphQLController)
}
