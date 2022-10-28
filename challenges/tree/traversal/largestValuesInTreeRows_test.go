package traversal_test

import (
	"reflect"
	"testing"

	"github.com/abdalla/go_dsa/challenges/tree/traversal"
	"github.com/abdalla/go_dsa/tree"
	"github.com/segmentio/encoding/json"
)

func TestLargestValuesInTreeRows(t *testing.T) {
	bt := &tree.Tree{
		Value: -1,
		Left: &tree.Tree{
			Value: 5,
		},
		Right: &tree.Tree{
			Value: 7,
			Right: &tree.Tree{
				Value: 1,
			},
		},
	}

	got := traversal.LargestValuesInTreeRows(bt)
	want := []int{-1, 7, 1}

	if len(got) != len(want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}

func TestLargestValuesInTreeRows2(t *testing.T) {
	jsons := []byte(`
	{"value": -1,
	"left": {
			"value": 5,
			"left": {
					"value": -1,
					"left": {
							"value": 10,
							"left": null,
							"right": null
					},
					"right": null
			},
			"right": null
	},
	"right": {
			"value": 7,
			"left": null,
			"right": {
					"value": 1,
					"left": {
							"value": 5,
							"left": null,
							"right": null
					},
					"right": null
			}
	}
}`)
	var bt tree.Tree
	err := json.Unmarshal(jsons, &bt)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	got := traversal.LargestValuesInTreeRows(&bt)
	want := []int{-1, 7, 1, 10}

	if len(got) != len(want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}

}
