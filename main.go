package main

import (
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ngramReq struct {
	Text string `json:"text"`
}

func splitNum(str string, num int) []string {
	strs := []string{}
	for i := 0; i < len([]byte(str)); i += num { // NOTE: 暫定的にちょうどになる場合のみ（半端な数が来たら＼(^o^)／）
		strs = append(strs, string([]byte(str)[i:(i+num)]))
	}
	return strs
}

func generateNGramString(str string) string {
	var i, j int
	bytes := []byte(str)
	bytes2 := make([]byte, len(str)*len(str))
	for i = 0; i < len(str); i++ {
		for j = 0; j < len(str); j++ {
			bytes2[i*len(str)+j] = bytes[(len(str)+j+i)%len(str)]
		}
	}
	splitRes := splitNum(string(bytes2), int(math.Sqrt(float64(len(bytes2)))))
	res := strings.Join(splitRes, "\n")
	return string(res)
}

func main() {
	r := gin.Default()
	r.POST("/ngram", func(c *gin.Context) {
		var req ngramReq
		err := c.BindJSON(&req)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		ngramString := generateNGramString(req.Text)
		c.JSON(http.StatusOK, gin.H{
			"response": ngramString,
		})
	})
	r.Run(":8000")
}
