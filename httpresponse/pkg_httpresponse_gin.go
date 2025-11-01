package httpresponse

import (
	"github.com/gin-gonic/gin"
)

type GinSuccess struct {
	CTX  *gin.Context
	Data interface{}
}

type GinError struct {
	CTX     *gin.Context
	Error   error  `json:"error"`
	Message string `json:"message"`
}

func GinJSONSuccess(p *GinSuccess) {
	JSONSuccess(p.CTX.Writer, p.CTX.Request, p.Data)
}

func GinJSONError(p *GinError) {
	JSONError(p.CTX.Writer, p.CTX.Request, p.Error, p.Message)
}
