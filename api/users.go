package api

import (
	"api-test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IndexUsers index
func IndexUsers(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func Show(c *gin.Context) {
	CT := c.Params.ByName("client_token")
	apparent := models.FindMidl(CT)
	app := models.FindApp(apparent.ClientToken)
	proxy := models.FindProxy(apparent.ProxyTag)
	node := models.FindNode(apparent.DummyNodeTag)
	c.JSON(200, app)
	c.JSON(200, proxy)
	c.JSON(200, node)
}
