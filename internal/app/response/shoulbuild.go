package response

import (
	"github.com/gin-gonic/gin"
	"spider/internal/app/types/types_common"
)

func ShouldBindForList(c *gin.Context, req types_common.IBaseListParam) error {
	if err := c.ShouldBindQuery(req); err != nil {
		return err
	}
	req.Adjust()
	return nil
}
