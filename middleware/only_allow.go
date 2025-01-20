package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/dto"
	"github.com/Caknoooo/go-gin-clean-starter/utils"
)

func OnlyAllow(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole := ctx.MustGet("role").(string)

		for _, role := range roles {
			if userRole == role {
				ctx.Next()
				return
			}
		}

		err := fmt.Sprintf(dto.ErrRoleNotAllowed.Error(), userRole)
		response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_TOKEN_NOT_VALID, err, nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}
}
