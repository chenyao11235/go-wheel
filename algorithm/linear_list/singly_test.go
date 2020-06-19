package linear_list

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var s = &Singly{
    value: 0,
    next: &Singly{
        value: 1,
        next: &Singly{
            value: 2,
            next:  nil,
        },
    },
}

func TestSingly_GetValue(t *testing.T) {
    assertions := assert.New(t)
    assertions.Equal(s.GetValue(0), 0)
    assertions.Equal(s.GetValue(1), 1)
    assertions.Equal(s.GetValue(2), 2)
    assertions.Nil(s.GetValue(3))
}

func TestSingly_GetIndex(t *testing.T) {
    assertions := assert.New(t)
    assertions.Equal(s.GetIndex(0), 0)
    assertions.Equal(s.GetIndex(1), 1)
    assertions.Equal(s.GetIndex(2), 2)
    assertions.Nil(s.GetIndex(3))
}

func TestSingly_Add(t *testing.T) {
    assertions := assert.New(t)
    assertions.Equal(s.Add(3), 3)
    assertions.Equal(s.GetValue(3), 3)
    assertions.Equal(s.GetIndex(3), 3)
    assertions.Nil(s.GetIndex(4))
}

func TestSingly_TargetAdd(t *testing.T) {
    assertions := assert.New(t)
    //assertions.Nil(s.TargetAdd(3, 3))
    assertions.Nil(s.TargetAdd(0, 0))
    //assertions.Nil(s.TargetAdd(1.5, 2))
    for s.next != nil {
        t.Log(s.value)
        s = s.next
    }
    //assertions.Equal(s.GetIndex(3), 2.5)
}
