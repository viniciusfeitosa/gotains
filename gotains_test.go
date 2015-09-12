package gotains_test

import (
	"testing"

	"github.com/viniciusfeitosa/gotains"
)

var inputTestToExists = []struct {
	collectionTest interface{}
	itensToVerify  interface{}
}{
	{[]string{"test1", "test2", "test3"}, "test1"},
	{[]int{1, 2, 3}, 1},
	{[]int64{1, 2, 3}, int64(1)},
	{[]float64{1.1, 2.1, 3.1}, 1.1},
}

var inputTestToDontExists = []struct {
	collectionTest interface{}
	itensToVerify  interface{}
}{
	{[]string{"test1", "test2", "test3"}, "test4"},
	{[]int{1, 2, 3}, 4},
}

var inputTestToExistsInArray = []struct {
	collectionTest interface{}
	itensToVerify  interface{}
}{
	{[3]string{"test1", "test2", "test3"}, "test1"},
	{[3]int{1, 2, 3}, 1},
}

var inputTestToDontExistsInArray = []struct {
	collectionTest interface{}
	itensToVerify  interface{}
}{
	{[3]string{"test1", "test2", "test3"}, "test4"},
	{[3]int{1, 2, 3}, 4},
}

func TestContainsWhenTypeNeitherSliceNorArray(t *testing.T) {
	_, err := gotains.Contains("testError", "testError")
	if err == nil && err.Error() != "InterfaceSlice() given a non-slice type" {
		t.Error("Validation has a problem")
	}
}

func TestContainsWhenExistItem(t *testing.T) {
	for _, tt := range inputTestToExists {
		ret, err := gotains.Contains(tt.collectionTest, tt.itensToVerify)
		if err != nil {
			t.Error(err)
		}
		if !ret {
			t.Errorf("Error to verify slice when item exist with values %v and %v", tt.collectionTest, tt.itensToVerify)
		}
	}

	for _, tt := range inputTestToExistsInArray {
		ret, err := gotains.Contains(tt.collectionTest, tt.itensToVerify)
		if err != nil {
			t.Error(err)
		}
		if !ret {
			t.Errorf("Error to verify arrays when item exist with values %v and %v", tt.collectionTest, tt.itensToVerify)
		}
	}
}

func TestContainsWhenDontExistItem(t *testing.T) {
	for _, tt := range inputTestToDontExists {
		ret, err := gotains.Contains(tt.collectionTest, tt.itensToVerify)
		if err != nil {
			t.Error(err)
		}
		if ret {
			t.Errorf("Error to verify when item in slice don't exist with values %s and %s", tt.collectionTest, tt.itensToVerify)
		}
	}
	for _, tt := range inputTestToDontExistsInArray {
		ret, err := gotains.Contains(tt.collectionTest, tt.itensToVerify)
		if err != nil {
			t.Error(err)
		}
		if ret {
			t.Errorf("Error to verify when item in array don't exist with values %s and %s", tt.collectionTest, tt.itensToVerify)
		}
	}
}
