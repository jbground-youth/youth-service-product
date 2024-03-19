package app

import (
	"github.com/gin-gonic/gin"
	"github.com/youth/product/pkg/setting"
)

func Run() {

	gin.SetMode(setting.ServerSetting.RunMode)

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//if err := r.Run(":8080"); err != nil {
	//	log.Fatalf("could not run server: %v", err)
	//}

}
