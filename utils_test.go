package main

import (
	"testing"
)

func asset(t *testing.T, exp interface{}, act interface{}, msg string) {
	if exp != act {
		t.Errorf(msg+" - expected: %s actual: %s", exp, act)
	}
}

func assetError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf(msg+" - Error: %s", err)
	}
}
