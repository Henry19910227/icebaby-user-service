package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/Henry19910227/icebaby-user-service/pkg/jwt"
	"github.com/spf13/viper"

	"github.com/Henry19910227/icebaby-user-service/global"
	"github.com/Henry19910227/icebaby-user-service/internal/controller"
	"github.com/Henry19910227/icebaby-user-service/internal/middleware"
	"github.com/Henry19910227/icebaby-user-service/internal/repository"
	"github.com/Henry19910227/icebaby-user-service/internal/service"
	"github.com/Henry19910227/icebaby-user-service/pkg/db"
	"github.com/Henry19910227/icebaby-user-service/pkg/logger"
	"github.com/Henry19910227/icebaby-user-service/pkg/upload"
	"github.com/gin-gonic/gin"

	_ "github.com/Henry19910227/icebaby-user-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	mysqlDB         *sql.DB
	userService     service.UserService
	loginService    service.LoginService
	registerService service.RegisterService
	jwtTool         jwt.Tool
	viperTool       *viper.Viper
)

func init() {
	setupViper()
	setupLogger()
	setupDB()
	setupLoginService()
	setupUserService()
	setupRegisterService()
	setupTokenTool()
}

// @title Henry
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	router := gin.New()
	router.Use(gin.CustomRecovery(middleware.Recover()))            //加入攔截panic中間層
	router.Use(gin.Logger())                                        //加入路由Logger
	router.Use(middleware.Cors())                                   //加入解決跨域中間層
	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	controller.NewUserController(router, userService, jwtTool)
	controller.NewLoginController(router, loginService, jwtTool)
	controller.NewRegisterController(router, registerService)

	router.Run(":9090")
}

func setupViper() {
	vp := viper.New()
	vp.SetConfigFile("./config/config.yaml")
	if err := vp.ReadInConfig(); err != nil {
		log.Fatalf(err.Error())
	}
	var mode string
	flag.StringVar(&mode, "m", "debug", "獲取運行模式")
	flag.Parse()
	vp.Set("Server.RunMode", mode)
	viperTool = vp
}

func setupDB() {
	setting := db.NewMysqlSetting(viperTool)
	mysqlDB = db.NewDB(setting)
}

func setupLogger() {
	setting := logger.NewGPLogSetting(viperTool)
	logger, err := logger.NewGPLogger(setting)
	if err != nil {
		log.Fatalf(err.Error())
	}
	global.Log = logger
}

func setupLoginService() {
	setting, err := upload.NewUploadSetting("./config/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	userService = service.NewUserService(repository.NewUserRepository(mysqlDB), upload.NewUploadTool(setting))
}

func setupRegisterService() {
	registerService = service.NewRegisterService(repository.NewUserRepository(mysqlDB))
}

func setupUserService() {
	loginService = service.NewLoginService(repository.NewUserRepository(mysqlDB))
}

func setupTokenTool() {
	setting, err := jwt.NewJWTSetting("./config/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}
	jwtTool = jwt.NewJWTTool(setting)
}
