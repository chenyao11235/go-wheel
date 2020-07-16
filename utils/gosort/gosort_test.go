package gosort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*介绍一些golang的标准库中的sort包的使用方法
 */

// 简单数据类型的排序
func TestSimpleSort(t *testing.T) {
	arr1 := []int{1, 4, 3, 2, 5}
	sort.Ints(arr1)
	assert.Equal(t, true, sort.IntsAreSorted(arr1), "")
	t.Log(arr1)
	arr2 := []float64{1.1, 2.3, 5.3, 3.4}
	sort.Float64s(arr2)
	assert.Equal(t, true, sort.Float64sAreSorted(arr2), "")
	t.Log(arr2)
	arr3 := []string{"e", "a", "d", "b", "z"}
	sort.Strings(arr3)
	assert.Equal(t, true, sort.StringsAreSorted(arr3), "")
	t.Log(arr3)
}

//复杂数据类型的排序，比如根据struct中的某个字段给struct排序
type Person struct {
	Name string
	Age  int
}

type Persons []Person

// Len, Less, Swap 其实是实现了sort包中的接口，使得struct可以sort包进行排序
func (p Persons) Len() int { return len(p) }
func (p Persons) Less(i, j int) bool {
	return p[i].Age < p[j].Age // 升序用<  降序用>  使用sort内置的reverse方法则不用修改less方法
}
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

var persons = Persons{
	{Name: "test1", Age: 20}, {Name: "test2", Age: 22}, {Name: "test3", Age: 21},
}

func TestCustomSort(t *testing.T) {
	sort.Sort(persons)
	t.Log(persons)
}

func TestReverse(t *testing.T) {
	sort.Sort(sort.Reverse(persons))
	t.Log(persons)
}

// 二分查找
func TestSearch1(t *testing.T) {
	x := 11
	s := []int{3, 6, 8, 11, 45}                                       // 注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x }) // 升序用>= , 降序用<=
	if pos < len(s) && s[pos] == x {
		t.Log(x, " 在 s 中的位置为：", pos)
	} else {
		t.Log("s 不包含元素 ", x)
	}
}

func TestSearch2(t *testing.T) {
	x := 11
	s := []int{45, 11, 8, 6, 3} // 注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] <= x })
	if pos < len(s) && s[pos] == x {
		t.Log(x, " 在 s 中的位置为：", pos)
	} else {
		t.Log("s 不包含元素 ", x)
	}
}

func TestSlice(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	t.Log(intervals)
}

func TestSlice1(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// 稳定排序
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	t.Log(intervals)
}
