package json

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFnSet(t *testing.T) {
	input1 := `{
		"key1": 123,
		"key2": {
		  "nested": 345.45,
		  "deepNesting": {
			"key3": 678.567575676,
			"key4": [
			  {
				"key5": 7987897.96878
			  }
			],
			"key5": [1, 2, 3]
		  }
		},
		"key3": [1, 2, 3],
		"key4": [
		  [1, 2],
		  [3, 4],
		  [3, 4, 5],
		  [
			3,
			4,
			{
			  "key8": 7987897.96878
			}
		  ]
		]
	  }`
	input2 := "newKey"
	input3 := "newVal"
	obj := make(map[string]interface{})
	err := json.Unmarshal([]byte(input1), &obj)
	assert.Nil(t, err)
	f := &fnSet{}
	v, err := f.Eval(obj, input2, input3)
	assert.Nil(t, err)
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bytes))
}
