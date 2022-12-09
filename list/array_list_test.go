package list_test

import (
	"encoding/json"
	"fmt"
	"github.com/kkakoz/pkg/list"
	"testing"
)

func TestList(t *testing.T) {
	source := []int{5, 4, 3, 2, 1}
	arrayList := list.NewArrayList(list.WithSlice(source))

	arrayList.AddLast(1)

	fmt.Println(arrayList.Get(0))
	fmt.Println(arrayList.Get(1))

	fmt.Println(arrayList.Delete(1))
	fmt.Println(arrayList.Delete(0))

	arrayList.AddLast(2)

	fmt.Println(arrayList.Slice())

	data, _ := json.Marshal(arrayList)
	fmt.Println(string(data))

	// test filter by
	arrayList = list.NewArrayList(list.WithSlice([]int{1, 2, 3}))
	fmt.Println(arrayList)

	fmt.Println(arrayList.FilterBy(func(i int) bool {
		return i == 1
	}).Slice())
}
