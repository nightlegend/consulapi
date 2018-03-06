package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	consulapi "github.com/nightlegend/consulapi/core/api"
	"github.com/nightlegend/consulapi/core/data/constdata"
	"github.com/nightlegend/consulapi/core/data/userdata"
	"github.com/nightlegend/consulapi/middleware"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
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

func init() {
	log.Info("set gin mode", gin.ReleaseMode)
	// gin.SetMode(gin.ReleaseMode)
}

// Start a application.
func Start() {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	// Default route "/"
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"stats": "success hit your api server..."})
	})

	// Delete a service.
	router.GET("/api/service/delete/:id", func(c *gin.Context) {
		serviceId := c.Param("id")
		flag := consulapi.ServiceDeRegister(serviceId)
		if flag {
			c.JSON(200, gin.H{"stats": "success"})
		} else {
			c.JSON(200, gin.H{"status": "fail", "message": "delete failed"})
		}
	})

	// Register a service.
	router.POST("/api/service/register", func(c *gin.Context) {
		var registerServiceInfo *ServiceRegistrInfo
		err := c.BindJSON(&registerServiceInfo)
		if err != nil {
			log.Println(err)
		}
		UUID := uuid.NewV4()
		// if err1 != nil {
		// 	log.Println(err1)
		// }
		registerServiceInfo.ID = UUID.String()
		if len(registerServiceInfo.TAGS) == 0 {
			registerServiceInfo.TAGS = []string{"docker", "server"}
		}
		flag := consulapi.RegisterService(registerServiceInfo.ID, registerServiceInfo.NAME, registerServiceInfo.TAGS,
			registerServiceInfo.ADDRESS, registerServiceInfo.PORT, constdata.NEW_REGISTER_TYPE)
		if flag {
			serviceList := consulapi.GetAllRegisterService()
			jsonString, _ := json.Marshal(serviceList)
			c.JSON(http.StatusOK, gin.H{"status": 200, "Message": "register success", "list": string(jsonString)})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 204, "Message": "register failed"})
		}
	})

	// Get the services list.
	router.GET("/api/service/list", func(c *gin.Context) {
		serviceList := consulapi.GetAllRegisterService()
		jsonString, _ := json.Marshal(serviceList)
		c.JSON(200, gin.H{"stats": "success", "list": string(jsonString)})
	})

	// Put a key-value to kv store.
	router.POST("/api/kv/put", func(c *gin.Context) {
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
	router.GET("/api/kv/get:key", func(c *gin.Context) {
		searchKey := c.Param("key")
		value := consulapi.Get(searchKey)
		c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": value})
	})

	// List all kv in kv store
	router.GET("/api/kv/list", func(c *gin.Context) {
		kvList := consulapi.ListAllKV()
		c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": kvList})
	})

	router.GET("/api/service/reload", func(c *gin.Context) {
		res := consulapi.ReloadData()
		if res {
			serviceList := consulapi.GetAllRegisterService()
			jsonString, _ := json.Marshal(serviceList)
			c.JSON(http.StatusOK, gin.H{"status": 200, "Message": "register success", "list": string(jsonString)})
			// c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": "Reload Data Successful"})
		} else {
			c.JSON(http.StatusOK, gin.H{"StatusCode": 201, "Message": "Reload Failed"})
		}
	})

	router.POST("/api/users/login", func(c *gin.Context) {
		var account *userdata.Accounts
		err := c.BindJSON(&account)
		log.Info(account.PASSWORD)
		log.Info(account.USERNAME)
		if err != nil {
			log.Error(err)
		}
		res := consulapi.Login(account)
		if res {
			c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": "Login Successed"})
		} else {
			c.JSON(http.StatusOK, gin.H{"StatusCode": 201, "Message": "Login Failed"})
		}
	})

	router.POST("/api/users/register", func(c *gin.Context) {
		var account *userdata.Accounts
		err := c.BindJSON(&account)
		if err != nil {
			log.Error(err)
		}
		res := consulapi.Register(account)
		if res {
			c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": "Login Successed"})
		} else {
			c.JSON(http.StatusOK, gin.H{"StatusCode": 201, "Message": "Login Failed"})
		}
	})

	// Server listen the port: 8015
	router.Run(":8015")
}
