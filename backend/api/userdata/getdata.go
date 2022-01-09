package userdata

import (
	"net/http"

	"github.com/chomosuke/backend/auth"
	"github.com/chomosuke/backend/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetData(c *gin.Context) {
	user := c.MustGet(auth.User).(db.User)
	c.JSON(http.StatusOK, user.Data.(bson.D).Map())
}
