package web

import (
	"fmt"
	"go_redirect/internal/pathresv"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Listen 启动 web 服务进行监听
func Listen() error {
	r := gin.Default()
	r.GET("/:idx/*vars", handleRedirect)
	return r.Run(":5555")
}

// handleRedirect 处理重定向
func handleRedirect(c *gin.Context) {
	idxStr := c.Param("idx")
	idx, err := strconv.Atoi(idxStr)
	if err != nil {
		c.String(http.StatusBadRequest, "无效的模板索引: "+idxStr)
		return
	}
	vars := c.Param("vars")
	res, err := pathresv.Handle(idx, vars)
	if err != nil {
		c.String(http.StatusBadRequest, "转换失败: "+err.Error())
		return
	}
	fmt.Println()
	log.Printf("idx: [%d], vars: [%s], redirect-link: [%s]\n\n", idx, vars, res)
	c.Redirect(http.StatusFound, res)
}
