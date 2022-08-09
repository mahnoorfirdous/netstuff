package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email       string `json:"email" binding:"required"`
	Name        string `json:"name" binding:"required"`
	AccountType string `json:"accounttype" binding:"required,oneof=Premium Free"`
	Password    string `json:"password" binding:"required"`
}

type RegisterParams struct {
	Email       string
	Name        string
	AccountType string
	Balance     float32
}

// type UserResponse struct {
// 	Email       string `json:"email"`
// 	Name        string `json:"name"`
// 	AccountType string `json:"accounttype"`
// 	// PasswordChangedAt time.Time `json:"password_changedat"`
// 	// RegisteredAt      time.Time `json:"registered_at"`
// 	Balance float32 `json:"balance"`
// }

//better to separate route function with actual storage function, associate storage with storage class

func (sv *Server) Register(ctx *gin.Context) {
	request := RegisterRequest{}
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	argument := RegisterParams{
		Email:       request.Email,
		Name:        request.Name,
		AccountType: request.AccountType,
		Balance:     0.0,
	}

	sv.mockdb.Users[request.Email] = LocalStorage{Registration: argument}

	ctx.JSON(http.StatusOK, gin.H{"response": argument})
}
