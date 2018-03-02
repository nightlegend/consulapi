# CONSUL API SERVER
A consul api server developed by golang.

## Description
> a high level api for consul, it`s provide register/deplet/query service api, hope it can make your work efficent and convenience.


## Run

### go run
<pre>
git clone https://github.com/nightlegend/consulapi.git ${golang_workspace}/src/github.com/nightlegend/
cd ${golang_workspace}/src/github.com/nightlegend/apigateway
go get
go build -a
export CONSUL_API="localhost:8500"
export MONGO_URL="localhost:27017"
go run server.go
</pre>

#### if you saw below, Congratulations.
<pre>
time="2018-03-02T16:10:59+08:00" level=info msg="new consul api..."
time="2018-03-02T16:10:59+08:00" level=info msg="10.222.49.65:8500"
time="2018-03-02T16:10:59+08:00" level=info msg=dev
time="2018-03-02T16:10:59+08:00" level=info msg="set gin moderelease"
time="2018-03-02T16:10:59+08:00" level=info msg="Init application configure..."
time="2018-03-02T16:10:59+08:00" level=info msg="Load environment is dev"
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> github.com/nightlegend/consulapi/router.Start.func1 (2 handlers)
[GIN-debug] GET    /api/service/delete/:id   --> github.com/nightlegend/consulapi/router.Start.func2 (2 handlers)
[GIN-debug] POST   /api/service/register     --> github.com/nightlegend/consulapi/router.Start.func3 (2 handlers)
[GIN-debug] GET    /api/service/list         --> github.com/nightlegend/consulapi/router.Start.func4 (2 handlers)
[GIN-debug] POST   /api/kv/put               --> github.com/nightlegend/consulapi/router.Start.func5 (2 handlers)
[GIN-debug] GET    /api/kv/get:key           --> github.com/nightlegend/consulapi/router.Start.func6 (2 handlers)
[GIN-debug] GET    /api/kv/list              --> github.com/nightlegend/consulapi/router.Start.func7 (2 handlers)
[GIN-debug] GET    /api/service/reload       --> github.com/nightlegend/consulapi/router.Start.func8 (2 handlers)
[GIN-debug] Listening and serving HTTP on :8015
</pre>

> you can hit http://localhost:8015 now.

## API

### register service

```go
router.POST("/api/service/register", func(c *gin.Context) {
	var registerServiceInfo *ServiceRegistrInfo
	err := c.BindJSON(&registerServiceInfo)
	if err != nil {
		log.Println(err)
	}
	UUID, err1 := uuid.NewV4()
	if err1 != nil {
		log.Println(err1)
	}
	registerServiceInfo.ID = UUID.String()
	if len(registerServiceInfo.TAGS) == 0 {
		registerServiceInfo.TAGS = []string{"docker", "server"}
	}
	flag := consulapi.RegisterService(registerServiceInfo.ID, registerServiceInfo.NAME, registerServiceInfo.TAGS,
		registerServiceInfo.ADDRESS, registerServiceInfo.PORT, constdata.NEW_REGISTER_TYPE)
	if flag {
		c.JSON(http.StatusOK, gin.H{"status": 200, "Message": "register success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": 204, "Message": "register failed"})
	}
})
```

### delete service
```go
router.GET("/api/service/delete/:id", func(c *gin.Context) {
	serviceId := c.Param("id")
	flag := consulapi.ServiceDeRegister(serviceId)
	if flag {
		c.JSON(200, gin.H{"stats": "success"})
	} else {
		c.JSON(200, gin.H{"status": "fail", "message": "delete failed"})
	}
})
```

### list service
```go
router.GET("/api/service/list", func(c *gin.Context) {
	serviceList := consulapi.GetAllRegisterService()
	jsonString, _ := json.Marshal(serviceList)
	c.JSON(200, gin.H{"stats": "success", "list": string(jsonString)})
})
```

### reload service
```go
router.GET("/api/service/reload", func(c *gin.Context) {
	res := consulapi.ReloadData()
	if res {
		c.JSON(http.StatusOK, gin.H{"StatusCode": 200, "Message": "Reload Data Successful"})
	} else {
		c.JSON(http.StatusOK, gin.H{"StatusCode": 201, "Message": "Reload Failed"})
	}
})
```
