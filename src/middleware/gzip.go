package middleware

import (
	"compress/gzip"
	"strings"

	"github.com/gin-gonic/gin"
)

func Gzip(c *gin.Context) {
	// Check if the client supports gzip
	if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
		c.Next()
		return
	}

	// Create gzip writer
	gz := gzip.NewWriter(c.Writer)
	defer gz.Close()

	// Replace the original ResponseWriter with a custom one that writes gzip
	c.Writer = &gzipWriter{Writer: gz, ResponseWriter: c.Writer}

	// Add gzip headers to the response
	c.Header("Content-Encoding", "gzip")
	c.Header("Vary", "Accept-Encoding")

	c.Next()
}

// gzipWriter wraps the http.ResponseWriter and compresses data before writing it.
type gzipWriter struct {
	gin.ResponseWriter
	Writer *gzip.Writer
}

// Write overrides the default Write method to compress the data before sending it to the client.
func (w *gzipWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}
