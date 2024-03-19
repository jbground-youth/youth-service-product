package api

import (
	"github.com/gin-gonic/gin"
	"github.com/youth/product/pkg/data"
	"github.com/youth/product/pkg/e"
	"github.com/youth/product/pkg/services"
	"net/http"
	"strconv"
)

func GetItems(c *gin.Context) {
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	//validation 추가

	//articleService := article_service.Article{ID: id}
	//exists, err := articleService.ExistByID()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
	//	return
	//}

	appG.Response(http.StatusOK, e.SUCCESS, id)

}

func GetItem(c *gin.Context) {
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))
	//value := c.Param("type")

	//validation 추가

	//articleService := article_service.Article{ID: id}
	//exists, err := articleService.ExistByID()

	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
	//	return
	//}

	appG.Response(http.StatusOK, e.SUCCESS, id)

}

func CreateItem(c *gin.Context) {
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	service := services.Item{}
	err := service.Create()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

func UpdateItem(c *gin.Context) {
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	//articleService := article_service.Article{ID: id}
	//exists, err := articleService.ExistByID()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
	//	return
	//}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}

func DeleteItem(c *gin.Context) {
	appG := data.Gin{C: c}
	id, _ := strconv.Atoi(c.Param("id"))

	//articleService := article_service.Article{ID: id}
	//exists, err := articleService.ExistByID()
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
	//	return
	//}

	appG.Response(http.StatusOK, e.SUCCESS, id)
}
