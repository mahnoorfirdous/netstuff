package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	Password    string
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
		return
	}

	argument := RegisterParams{
		Email:       request.Email,
		Name:        request.Name,
		AccountType: request.AccountType,
		Balance:     0.0,
	}

	sv.mockdb.Users[request.Email] = LocalStorage{Registration: argument}
	sv.mockdb.storeHashedPassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": argument})
}

func (as *AllStore) storeHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	return string(bytes), err
}

func (as *AllStore) verifyHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
