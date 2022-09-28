package api

import (
	"database/sql"
	"net/http"

	db "github.com/Xebec19/pathshaala-server/db/sqlc"
	"github.com/Xebec19/pathshaala-server/utils"
	"github.com/gin-gonic/gin"
)

type createUsersRequest struct {
	FirstName string          `json:"first_name"`
	LastName  sql.NullString  `json:"last_name"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	Role      db.NullUserRole `json:"role"`
}

func (server *Server) createAccount(c *gin.Context) {
	var req createUsersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	arg := db.CreateUsersParams{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
	}
	account, err := server.query.CreateUsers(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, account)
}
