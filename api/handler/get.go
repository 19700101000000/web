package handler

import (
	"github.com/19700101000000/system-sample/api/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H(db.Categories()))
}

func GetUser(c *gin.Context) {
	observer, _ := getAuth(c)
	c.JSON(http.StatusOK, gin.H(db.User(c.Param("name"), UserList[observer].ID)))
}

/* get Galleries */
func GetMyGalleries(c *gin.Context) {
	var isSelf bool
	name := c.Param("name")

	if n, ok := getAuth(c); ok && name == n {
		isSelf = true
	} else {
		isSelf = false
	}

	c.JSON(http.StatusOK, db.MyGalleries(name, isSelf))
}
func GetGalleries(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H(db.Galleries()))
}

/* get wanteds */
func GetWorksWanteds(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H(db.WorksWanteds(c.Param("name"), false)))
}
func GetMyWanteds(c *gin.Context) {
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}

	c.JSON(http.StatusOK, db.WorksWanteds(name, ok))
}

/* get requests */
func GetMyRequests(c *gin.Context) {
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}
	c.JSON(http.StatusOK, db.WorksRequests(name, ok))
}
func GetWorksRequests(c *gin.Context) {
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}
	wanted, _ := strconv.Atoi(c.Param("wanted"))
	c.JSON(http.StatusOK, db.WorksRequests2Wanted(name, wanted))
}

/* get info */
func GetWorksInfo(c *gin.Context) {
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}
	c.JSON(http.StatusOK, db.WorksInfo(name))
}

/* get ranks */
func GetRates(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, db.MonitorsWithName(name))
}
