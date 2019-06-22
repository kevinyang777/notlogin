package response

import (
	"github.com/gin-gonic/gin"
)

// func Message(status bool, message string) map[string]interface{} {
// 	return map[string]interface{}{"status": status, "message": message}
// }

// func Respond(w http.ResponseWriter, data map[string]interface{}) {
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }

func SuccessResponse(data *gin.Context, message map[string]interface{}) {
	data.JSON(200, gin.H{
		"message": "success",
	})
}

func FailResponse(data *gin.Context) {
	data.JSON(400, gin.H{
		"message": "fail",
	})
}
