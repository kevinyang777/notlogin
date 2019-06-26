package response

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(data *gin.Context, message []byte) {
	var parsedData map[string]interface{}
	err := json.Unmarshal(message, &parsedData)
	if err != nil {
		panic(err)
	}
	data.JSON(200, gin.H{
		"message": parsedData,
	})
}

func FailResponse(data *gin.Context) {
	data.JSON(400, gin.H{
		"message": "fail",
	})
}
