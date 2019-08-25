package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Parms() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		// 目前用不到
		var (
			reqBody []byte
			err     error
		)
		if reqBody, err = ioutil.ReadAll(c.Request.Body); len(reqBody) != 0 && err == nil {
			var m map[string]interface{}
			if err := json.Unmarshal(reqBody, &m); err == nil {
				fmt.Println("set Body")
				c.Set("body", m)
			}
		}

		c.Set("example", "12345")

		// before request

		c.Next()
	}
}
