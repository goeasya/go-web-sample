package util_test

import (
	"gin-sample/util"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	out := util.GenerateUUID()

	if strings.Contains(out, "-") {
		t.Fatal("uuid should not contain -")
	}

	if len(out) != 32 {
		t.Fatal("out len must be equal 32")
	}
}

func TestGenerateRandomID(t *testing.T) {
	type testCaseStruct struct {
		Size    int
		Kind    int
		Pattern string
	}

	testCaseList := []testCaseStruct{
		{
			Size:    8,
			Kind:    util.RandomKindNum,
			Pattern: "^[0-9]*$",
		},
		{
			Size:    10,
			Kind:    util.RandomKindUpperLower,
			Pattern: "^[A-Za-z]+$",
		},
		{
			Size:    8,
			Kind:    util.RandomKindUpper,
			Pattern: "^[A-Z]+$",
		},
		{
			Size:    10,
			Kind:    util.RandomKindUpperNum,
			Pattern: "^[A-Z0-9]+$",
		},
		{
			Size:    8,
			Kind:    util.RandomKindLowerNum,
			Pattern: "^[0-9a-z]+$",
		},
		{
			Size:    8,
			Kind:    util.RandomKindLower,
			Pattern: "^[a-z]+$",
		},
		{
			Size:    10,
			Kind:    util.RandomKindAll,
			Pattern: "^[0-9A-Za-z]+$",
		},
	}

	for _, testCase := range testCaseList {
		out := util.GenerateRandomID(testCase.Size, testCase.Kind)
		if len(out) != testCase.Size {
			t.Fatalf("testcase: %d, %d, out len: %d, except len: %d",
				testCase.Size, testCase.Kind, len(out), testCase.Size)
		}

		result, _ := regexp.MatchString(testCase.Pattern, out)
		if !result {
			t.Fatalf("testcase: %d, %d, out must match patter %s",
				testCase.Size, testCase.Kind, testCase.Pattern)
		}
	}

}
