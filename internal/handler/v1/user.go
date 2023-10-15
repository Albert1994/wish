package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"wish/internal/dto"
	"wish/internal/helpers"
)

func (h *Handler) initUsersRouters(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignUp)
		users.POST("/sign-in", h.userSignIn)
		users.POST("/auth/refresh", h.userRefresh)
	}
}

// @Summary User SignUp
// @Tags users-auth
// @Description create user account
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-up [post]
func (h *Handler) userSignUp(c *gin.Context) {
	var inp dto.UserSingUpInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	if err := h.services.Users.SingUp(inp); err != nil {
		if errors.Is(err, helpers.ErrUserAlreadyExists) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) userSignIn(context *gin.Context) {

}

func (h *Handler) userRefresh(context *gin.Context) {

}
