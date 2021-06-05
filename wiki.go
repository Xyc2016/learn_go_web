package main

import (
	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a"`
}
type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}
type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}
type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}
func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"a": 1,
		})
	})
	v2 := r.Group("/v2")
	v2.POST("/article", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "OK",
		})
	})
	r.Run("0.0.0.0:8000")
}
