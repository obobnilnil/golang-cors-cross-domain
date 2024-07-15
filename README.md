<div align="center">
   <img src="https://www.ardanlabs.com/images/training-landing/go/go-intro.svg" alt="Go" width="300"/>
</div>

<H1>Handling Cross-Domain Issues in Go with Gin Framework by Backend as a Proxy</H1>

This guide details a robust solution to handle cross-domain issues when your frontend application needs to communicate with external APIs, like CyberReason,
which are typically blocked by browser security policies due to the Same-Origin Policy. By setting up a backend service that functions as an intermediary,
we can effectively bypass these cross-domain restrictions, allowing seamless API interactions.

<H1>Problem Overview</H1>
Integrating external APIs into your frontend can lead to a Cross-Domain Policy Error. This is because modern web browsers enforce a security measure known as the Same-Origin Policy, 
which prevents web pages from making requests to a domain different from the one that served them.

<H1>Solution: Implementing Backend as a Proxy</H1>
To navigate this limitation, a backend service is employed to perform API calls to the external service (e.g., CyberReason) on behalf of the frontend. 
This backend then provides its own API endpoint, which is accessible to the frontend without any cross-domain issues. 

<pre>
   router := gin.Default()
   store := cookie.NewStore([]byte("secret"))
   router.Use(sessions.Sessions("mysession", store))
   // You start a Gin server.
   config := cors.DefaultConfig()
   config.AllowOrigins = []string{"*"}
   config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}
   config.AllowHeaders = []string{"Origin", "Content-Type", "X-Auth-Token", "Authorization"}
   config.AllowCredentials = true
   // After you create a line of configuration, you have to apply it to the router you have made.
   router.Use(cors.New(config))
   // Then, for every endpoint with a router parameter, you will allow CORS policy.
   router.POST("/api/loginProxy", h.LoginHandlers)
   router.POST("/api/widgetsProxy", h.WidgetsHandlers)
   router.GET("/api/groupsProxy", h.GroupsHandlers)
   router.POST("/api/graphMalopsResolutionTracking", h.GraphMalopsResolutionTrackingHandler)
</pre>

<H1>***The code shown in the repository is an example of how to configure CORS policy for every endpoint created by the backend to fix cross-domain problems.</H1>
Golang/Gin-framework

<div align="center">
    <img src="https://i.pinimg.com/originals/17/f9/9c/17f99c7a041cb2ddca13e3d5aa9ef28f.jpg" alt="Winter" width="300"/>
</div>
