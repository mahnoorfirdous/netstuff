package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Notebook struct {
	Model        string `json:"model"`
	Manufacturer string `json:"manufacturer"`
	Width        string `json:"width"`
	Binding      string `json:"binding" binding:"oneof=Tape Spiral"`
}

func (sv *Server) AddNotebook(ctx *gin.Context) {
	request := Notebook{}
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	sv.mockdb.Notebooks = make([]Notebook, 1)
	sv.mockdb.Notebooks = append(sv.mockdb.Notebooks, request)
}
