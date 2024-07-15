// package main

// import (
// 	"cyberreason_cross_domain/handler"
// 	"cyberreason_cross_domain/service"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	s := service.NewServiceAdapter()
// 	h := handler.NewHanerhandlerAdapter(s)

// 	router := gin.Default()
// 	config := cors.DefaultConfig()
// 	config.AllowOrigins = []string{"*"}
// 	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
// 	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
// 	config.AllowCredentials = true

// 	router.Use(cors.New(config))

// 	router.POST("/api/loginProxy", h.LoginHandlers)
// 	router.POST("/api/widgetsProxy", h.WidgetsHandlers)
// 	router.GET("/api/groupsProxy", h.GroupsHandlers)
// 	router.POST("/api/graphMalopsResolutionTracking", h.GraphMalopsResolutionTrackingHandler)

// 	err := router.Run(":8881")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

package main

import (
	"cyberreason_cross_domain/handler"
	"cyberreason_cross_domain/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	s := service.NewServiceAdapter()
	h := handler.NewHanerhandlerAdapter(s)

	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.POST("/api/loginProxy", h.LoginHandlers)
	router.POST("/api/widgetsProxy", h.WidgetsHandlers)
	router.GET("/api/groupsProxy", h.GroupsHandlers)
	router.POST("/api/graphMalopsResolutionTracking", h.GraphMalopsResolutionTrackingHandler)

	err := router.Run(":8881")
	if err != nil {
		panic(err.Error())
	}
}
