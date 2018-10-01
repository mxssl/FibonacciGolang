package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 0
	output := getNthFib(1)
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	expected := 1597
	output := getNthFib(18)
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}
