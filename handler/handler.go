// package handler

// import (
// 	"cyberreason_cross_domain/model"
// 	"cyberreason_cross_domain/service"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type HandlerPort interface {
// 	LoginHandlers(c *gin.Context)
// 	WidgetsHandlers(c *gin.Context)
// 	GroupsHandlers(c *gin.Context)
// 	GraphMalopsResolutionTrackingHandler(c *gin.Context)
// }

// type handlerAdapter struct {
// 	s service.ServicePort
// }

// func NewHanerhandlerAdapter(s service.ServicePort) HandlerPort {
// 	return &handlerAdapter{s: s}
// }

// func (h *handlerAdapter) LoginHandlers(c *gin.Context) {
// 	var login model.LoginRequest
// 	if err := c.ShouldBindJSON(&login); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
// 		return
// 	}
// 	cookies, err := h.s.LoginServices(login)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Login(cyberreason) successfully.", "cookie": cookies})
// }

// func (h *handlerAdapter) WidgetsHandlers(c *gin.Context) {
// 	var widgets map[string]interface{}
// 	if err := c.ShouldBindJSON(&widgets); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
// 		return
// 	}

// 	cookies := map[string]string{}
// 	for _, cookie := range c.Request.Cookies() {
// 		cookies[cookie.Name] = cookie.Value
// 	}
// 	// fmt.Println(cookies)
// 	// fmt.Printf("widgets %v", cookies)

// 	response, err := h.s.WidgetsServices(widgets, cookies)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
// 		return
// 	}

// 	// c.JSON(http.StatusOK, gin.H{"status": "OK", "message": response})
// 	c.Data(http.StatusOK, "application/json", response)
// }

// func (h *handlerAdapter) GroupsHandlers(c *gin.Context) {
// 	// ดึง cookies จาก request
// 	cookies := map[string]string{}
// 	for _, cookie := range c.Request.Cookies() {
// 		cookies[cookie.Name] = cookie.Value
// 	}
// 	// fmt.Printf("groups %v", cookies)

// 	response, err := h.s.GroupsServices(cookies)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
// 		return
// 	}

// 	c.Data(http.StatusOK, "application/json", response)
// }

package handler

import (
	"cyberreason_cross_domain/model"
	"cyberreason_cross_domain/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type HandlerPort interface {
	LoginHandlers(c *gin.Context)
	WidgetsHandlers(c *gin.Context)
	GroupsHandlers(c *gin.Context)
	GraphMalopsResolutionTrackingHandler(c *gin.Context)
}

type handlerAdapter struct {
	s service.ServicePort
}

func NewHanerhandlerAdapter(s service.ServicePort) HandlerPort {
	return &handlerAdapter{s: s}
}

func (h *handlerAdapter) LoginHandlers(c *gin.Context) {
	var login model.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	cookies, err := h.s.LoginServices(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	// เก็บ username ใน session
	session := sessions.Default(c)
	session.Set("userID", login.Username)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Login(cybereason) successfully.", "cookie": cookies})
}

func (h *handlerAdapter) WidgetsHandlers(c *gin.Context) {
	var widgets map[string]interface{}
	if err := c.ShouldBindJSON(&widgets); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	// ดึง userID จาก session
	session := sessions.Default(c)
	userID := session.Get("userID").(string)

	response, err := h.s.WidgetsServices(widgets, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", response)
}

func (h *handlerAdapter) GroupsHandlers(c *gin.Context) {
	// ดึง userID จาก session
	session := sessions.Default(c)
	userID := session.Get("userID").(string)

	response, err := h.s.GroupsServices(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", response)
}

func (h *handlerAdapter) GraphMalopsResolutionTrackingHandler(c *gin.Context) {
	var req model.MalopResolutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": "Invalid data"})
		return
	}

	filename, err := h.s.GraphMalopsResolutionTrackingServices(req.MalopResolution)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.File(filename)
}
