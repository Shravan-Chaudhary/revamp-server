package user

import (
	"net/http"

	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/response"
	"github.com/Shravan-Chaudhary/revamp-server/internal/pkg/types"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	responseHandler *response.ResponseHandler;
	userRepo UserRepository;
}

func NewUserHandler(userRepo UserRepository, responseHandler *response.ResponseHandler) *UserHandler {
	return &UserHandler{
		responseHandler,
		userRepo,
	}
}

func (h *UserHandler) HandleCreateUser(c *gin.Context) {
	user := types.User{
		ID: 	  "1",
		FIRSTNAME: "Shravan",
		LASTNAME: "Chaudhary",
	}
	h.responseHandler.Send(c, http.StatusOK, response.Messages.Success, user,)
}