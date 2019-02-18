package rest

import (
	"net/http"

	"talk-docker/pkg"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Repository  ...
type Repository interface {
	ListTalks() (*[]pkg.Talk, error)
}

// Server ...
type Server struct {
	repo Repository
}

// NewServer  ...
func NewServer(repo Repository) *Server {
	return &Server{
		repo: repo,
	}
}

// Run  ...
func (s *Server) Run() error {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/ping", ping())
	router.GET("/talks", listTalks(s.repo))

	return router.Run()
}

func ping() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func listTalks(repo Repository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		talks, err := repo.ListTalks()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, talks)
	}
}
