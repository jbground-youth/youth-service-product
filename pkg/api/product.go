package api

import (
	"github.com/gin-gonic/gin"
	"github.com/youth/product/pkg/data"
	"github.com/youth/product/pkg/e"
	"net/http"
	"strconv"
)

func GetProducts(c *gin.Context) {

	//c.FullPath() == "/user/:id"
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	appG.Response(http.StatusOK, e.SUCCESS, id)

}

func GetProduct(c *gin.Context) {
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	appG.Response(http.StatusOK, e.SUCCESS, id)

}

func CreateProduct(c *gin.Context) {
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

func UpdateProduct(c *gin.Context) {
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

func DeleteProduct(c *gin.Context) {
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	appG.Response(http.StatusOK, e.SUCCESS, id)
}
