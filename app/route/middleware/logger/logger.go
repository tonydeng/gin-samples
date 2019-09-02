package logger

import (
	"../../../config"
	"../../../utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type bodyLogWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var accessChannel = make(chan string, 100)

func (w bodyLogWrite) write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWrite) writeString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func SetUp() gin.HandlerFunc {
	go handleAccesChannel()

	return func(c *gin.Context) {
		bodyLogWrite := &bodyLogWrite{
			body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWrite

		startTime := utils.GetCurrentMilliTime()

		c.Next()

		responseBody := bodyLogWrite.body.String()

		var responseCode int
		var responseMsg string
		var responseData interface{}

		if responseBody != "" {
			response := utils.Response{}
			err := json.Unmarshal([]byte(responseBody), &response)

			if (err == nil) {
				responseCode = response.Code
				responseMsg = response.Message
				responseData = response.Data
			}
		}

		endTime := utils.GetCurrentMilliTime()

		if c.Request.Method == "POST" {
			_ = c.Request.ParseForm()
		}

		accessLogMap := make(map[string]interface{})

		accessLogMap["request_time"] = startTime
		accessLogMap["request_method"] = c.Request.Method
		accessLogMap["request_uri"] = c.Request.RequestURI
		accessLogMap["request_proto"] = c.Request.Proto
		accessLogMap["request_referer"] = c.Request.Referer()
		accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
		accessLogMap["reqeust_client_ip"] = c.ClientIP()

		accessLogMap["response_time"] = endTime
		accessLogMap["response_code"] = responseCode
		accessLogMap["response_msg"] = responseMsg
		accessLogMap["response_data"] = responseData

		accessLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)

		accessLogJson, _ := utils.JsonEncode(accessLogMap)

		accessChannel <- accessLogJson
	}
}

func handleAccesChannel() {
	if f, err := os.OpenFile(config.AppAccessLogName,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0666); err != nil {
		log.Println(err)
	} else {
		for accessLog := range accessChannel {
			_, _ = f.WriteString(accessLog + "\n")
		}
	}
	return
}
