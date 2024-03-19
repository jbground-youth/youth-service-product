package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/youth/product/pkg/api"
	"github.com/youth/product/pkg/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/items")
	apiv1.Use(jwt.JWT())

	apiv1.GET("/", api.GetItems)
	apiv1.GET("/:type", api.GetItems)
	apiv1.GET("/:type/:id", api.GetItem)
	apiv1.POST("/:id", api.CreateItem)
	apiv1.POST("/:id", api.UpdateItem)
	apiv1.POST("/:id", api.DeleteItem)

	api2 := r.Group("/products")

	api2.GET("/", api.GetProducts)
	api2.GET("/:id", api.GetProduct)
	api2.POST("/", api.CreateProduct)
	api2.PUT("/:id", api.UpdateProduct)
	api2.DELETE("/:id", api.DeleteProduct)

	return r
}
