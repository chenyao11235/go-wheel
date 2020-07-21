package gomath

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestGoMath math 库中常用的函数
func TestGoMath(t *testing.T) {
	// 向上取整
	assert.Equal(t, 4.0, math.Ceil(3.14), "")
	// 向下取整
	assert.Equal(t, 3.0, math.Floor(3.14), "")
	// 取绝对值
	assert.Equal(t, 3.1415926, math.Abs(-3.1415926), "")
	// 求底数
	assert.Equal(t, 3.0, math.Log2(8), "")
}
