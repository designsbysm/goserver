package tools

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIDParam(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return 0, ErrIDParameterRequired
	} else if id == 0 {
		return 0, ErrIDParameterRequired
	}

	return uint(id), nil
}
