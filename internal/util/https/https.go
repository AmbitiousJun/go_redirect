package https

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// ProxyRequest 代理当前请求到指定的 remote 地址上, shouldText 表示是否只代理文本请求
func ProxyRequest(c *gin.Context, remote string, shouldText bool) error {
	// 解析远程地址
	remoteURL, err := url.Parse(remote)
	if err != nil {
		return err
	}
	remoteURL.RawQuery = c.Request.URL.RawQuery

	// 创建请求
	req, err := http.NewRequest(c.Request.Method, remoteURL.String(), c.Request.Body)
	if err != nil {
		return err
	}

	// 复制请求头
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// 执行请求
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 判断是否是文本请求
	contentType := resp.Header.Get("content-type")
	if shouldText && !IsTextContent(contentType) {
		return fmt.Errorf("非文本响应: [%s]", contentType)
	}

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 响应的 headers 拷贝到 c
	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// 响应
	c.Status(resp.StatusCode)
	if _, err := c.Writer.Write(body); err != nil {
		return err
	}
	return nil
}
