package service

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type buildServer struct {
	store *models.BuildStore
}

func NewBuildServer() *buildServer {
	store := models.New()
	return &buildServer{store: store}
}

func (bs *buildServer) addBuildToQueryHandler(c *gin.Context) {
	type RequestBuild struct {
		Name string
	}

	var rb RequestBuild
	if err := c.ShouldBindJSON(&rb); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	buildId := bs.store.Save(rb.Name)
	c.JSON(http.StatusOK, gin.H{"buildId": buildId, "Message": "Build was saves successfully"})
}

func (bs *buildServer) getBuildByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	build, err := bs.store.GetBuildById(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, build)
}

func (bs *buildServer) getAllBuildsHandler(c *gin.Context) {
	builds := bs.store.GetAllBuildsFromQuery()
	c.JSON(http.StatusOK, builds)
}

func main() {
	router := gin.Default()
	server := NewBuildServer()

	router.POST("/build/", server.addBuildToQueryHandler)
	router.GET("/task/", server.getBuildByIdHandler)
	router.GET("/tasks/", server.getAllBuildsHandler)

	router.Run("localhost:" + os.Getenv("SERVERPORT"))
}
