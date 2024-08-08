package web

import (
	"fmt"
	"go_redirect/internal/pathresv"
	"go_redirect/internal/util/https"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Listen 启动 web 服务进行监听
func Listen() error {
	r := gin.Default()
	r.Any("/:idx/*vars", handleRedirect)
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

	// 1 匹配模板
	res, err := pathresv.Handle(idx, vars)
	if err != nil {
		c.String(http.StatusBadRequest, "转换失败: "+err.Error())
		return
	}

	// 2 尝试代理请求
	err = https.ProxyRequest(c, res, true)
	if err == nil {
		log.Printf("请求被代理 => {idx: [%d], vars: [%s], remote: [%s]}\n", idx, vars, res)
		return
	}
	printError := fmt.Sprintf("请求无法被代理 => {method: %s, url: %s, error: %s}\n", c.Request.Method, c.Request.URL, err)
	log.Println(printError)

	// 3 重定向
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "非 GET 请求, 无法重定向", "error": printError})
		return
	}
	log.Printf("执行重定向 => {idx: [%d], vars: [%s], redirect-link: [%s]}\n\n", idx, vars, res)
	c.Redirect(http.StatusFound, res)
}
