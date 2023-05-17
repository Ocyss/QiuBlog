package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"math"
	"net/http"
	"os"
	"time"
)

// Logger is the logrus logger handler
// https://github.com/toorop/gin-logrus
func Logger(logger log.FieldLogger, notLogged ...string) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}

	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, p := range notLogged {
			skip[p] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		if _, ok := skip[path]; ok {
			return
		}

		entry := logger.WithFields(log.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			//"latency":    latency, // time to process
			"clientIP": clientIP,
			//"method":     c.Request.Method,
			//"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			//"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("\"%s %s Code:%d (%dms)\"", c.Request.Method, path, statusCode, latency)
			if statusCode >= http.StatusInternalServerError {
				entry.Error(msg, clientUserAgent)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(msg, clientUserAgent)
			} else {
				entry.Info(msg)
			}
		}
	}
}

func LoggerDebug(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	entry := log.WithFields(log.Fields{
		//"httpMethod":   httpMethod,
		//"absolutePath": absolutePath,
		"handlerName": handlerName,
		"nuHandlers":  nuHandlers,
	})
	msg := fmt.Sprintf("%-6s %-25s", httpMethod, absolutePath)
	entry.Debug(msg)
}
