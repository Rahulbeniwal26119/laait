package object_test

import (
	"laait/object"
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello := &object.String{Value: "Hello World"}
	hello2 := &object.String{Value: "Hello World"}
	diff1 := &object.String{Value: "My name is Rahul"}
	diff2 := &object.String{Value: "My name is Rahul"}

	if hello.HashKey() != hello2.HashKey() {
		t.Errorf("string with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("string with same content have different hash keys")
	}

	if hello.HashKey() == diff1.HashKey() {
		t.Errorf("string with different content have same hash keys")
	}
}
