package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	a := NewArray(7)
	a.Print()
	assert.Nil(t, a.Insert(0, 0), "no error")
	assert.Nil(t, a.Insert(1, 1), "no error")
	assert.Equal(t, 2, a.Len(), "len is 2")
	assert.Nil(t, a.Insert(2, 2), "no error")
	assert.Nil(t, a.Insert(3, 3), "no error")
	a.Print()
	assert.Nil(t, a.Insert(3, 3), "no error")
	a.Print()

	if v, err := a.Find(3); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
	a.Print()

	if v, err := a.Delete(3); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
	a.Print()
}
