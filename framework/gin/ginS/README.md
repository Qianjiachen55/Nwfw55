# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
	"github.com/Qianjiachen55/Nwfw55/framework/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
