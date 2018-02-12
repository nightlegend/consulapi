package router

import (
	"encoding/json"
	consulapi "github.com/nightlegend/consulapi/core/api"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

type ServiceRegistrInfo struct {
	ID      string   `json:"id" binding:"required"`
	NAME    string   `json:"name" binding:"required"`
	TAGS    []string `json:"tags" binding:"required"`
	ADDRESS string   `json:"address" bindinbg:"required"`
	PORT    int      `json:"port" binding:"required"`
}

type KVEntry struct {
	KEY   string `json:"key" binding:"required"`
	VALUE string `json:"value" binding:"required"`
}

/*
 *
 * Support for CORS function.
 *
 */
func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			log.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
		c.Next()
	}
}

// Start a application.
func Start() {
	router := gin.New()
	router.Use(CORSMiddleware())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"stats": "success"})
	})
	// Delete a service.
	router.GET("/delete/:id", func(c *gin.Context) {
		serviceId := c.Param("id")
		flag := consulapi.ServiceDeRegister(serviceId)
		if flag {
			c.JSON(200, gin.H{"stats": "success"})
		} else {
			c.JSON(200, gin.H{"status": "fail", "message": "delete failed"})
		}
	})

	// Register a service.
	router.POST("/register", func(c *gin.Context) {
		var registerServiceInfo *ServiceRegistrInfo
		err := c.BindJSON(&registerServiceInfo)
		if err != nil {
			log.Println(err)
		}
		UUID := uuid.NewV4()
		registerServiceInfo.ID = UUID.String()
		if len(registerServiceInfo.TAGS) == 0 {
			registerServiceInfo.TAGS = []string{"docker", "server"}
		}
		flag := consulapi.RegisterService(registerServiceInfo.ID, registerServiceInfo.NAME, registerServiceInfo.TAGS,
			registerServiceInfo.ADDRESS, registerServiceInfo.PORT)
		if flag {
			c.JSON(http.StatusOK, gin.H{"status": 200, "Message": "register success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 204, "Message": "register failed"})
		}
	})

	// Get the services list.
	router.GET("/list", func(c *gin.Context) {
		serviceList := consulapi.GetAllRegisterService()
		jsonString, _ := json.Marshal(serviceList)
		c.JSON(200, gin.H{"stats": "success", "list": string(jsonString)})
	})

	// Put a key-value to kv store.
	router.POST("/kv/put", func(c *gin.Context) {
		var kvEntry *KVEntry
		err := c.BindJSON(&kvEntry)
		if err != nil {
			log.Println(err)
		}
		log.Println(kvEntry.KEY)
		log.Println(kvEntry.VALUE)
		flag := consulapi.Put(kvEntry.KEY, kvEntry.VALUE)
		if flag {
			c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"StatusCode": 100, "Message": "failed"})
		}
		// c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": "success"})
	})

	// Search key from kv store
	router.GET("/kv/get:key", func(c *gin.Context) {
		searchKey := c.Param("key")
		value := consulapi.Get(searchKey)
		c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": value})
	})

	// List all kv in kv store
	router.GET("/kv/list", func(c *gin.Context) {
		kvList := consulapi.ListAllKV()
		c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": kvList})
	})

	// Server listen the port: 8015
	router.Run(":8015")
}
