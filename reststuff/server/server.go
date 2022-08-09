package restapi

import (
	"fmt"
	"net/http"
	token "samplerest/token"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
}

type Server struct {
	sconfig Config
	token   token.Maker
	router  *gin.Engine
	mockdb  AllStore
}

type LocalStorage struct {
	Registration RegisterParams
	Purchases    Notebook
}

type AllStore struct {
	Users     map[string]LocalStorage
	Notebooks []Notebook
}

func NewServer(sconfig Config) (*Server, error) {

	server := &Server{mockdb: AllStore{Users: make(map[string]LocalStorage)}}
	_, err := token.NewPasetoMaker(sconfig.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("../templates/*.html")
	//Routing
	router.POST("/accounts", server.Register)
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"content": "Simple index page",
		})
	})
	//TODO: Only special users like admin should be able to use this handle
	router.PUT("/notebooks", server.AddNotebook)
	server.router = router
	return server, nil
}

func (sv *Server) Start(address string) error {
	return sv.router.Run(address)
}
