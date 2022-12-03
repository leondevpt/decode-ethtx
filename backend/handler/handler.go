package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leondevpt/decode-ethtx/backend/pkg"
	"log"
	"net/http"
)

type Req struct {
	Tx string `json:"tx" form:"tx" binding:"required"`
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, RespOK("pong"))
}

func DecodeHandler(c *gin.Context) {
	var req Req
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("parse input err:%s\n", err)
		c.JSON(http.StatusOK, BuildResponse(10001, fmt.Sprintf("bad request:%s", err.Error()), nil))
		return
	}
	fmt.Printf("req:%+v\n", req)
	if len(req.Tx) == 0 {
		c.JSON(http.StatusOK, BuildResponse(10002, "input is empty", nil))
		return
	}

	tx, err := pkg.DecodeTx(req.Tx)
	if err != nil {
		log.Printf("DecodeTx data err:%s\n", err)
		c.JSON(http.StatusOK, BuildResponse(20001, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, RespOK(tx))
}
