package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/northbright/pathhelper"
)

func main() {
	router := gin.Default()
	// Set a new limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.GET("/", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, htmlStr)
	})

	router.POST("/", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Get current executable dir.
		dir, _ := pathhelper.GetCurrentExecDir()
		// Make absolute file path.
		dst := path.Join(dir, file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":8080")
}

var htmlStr = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
</head>
<body>
  <div align="center">
      <form method="POST" action="/" enctype="multipart/form-data">
          <input name="file" type="file">
          <input type="submit" value="Upload" />
      </form>
  </div>
</body>
</html>
`
