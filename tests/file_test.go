package tests

import (
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	// file not exist
	tmp1 := "./test.log"
	tmp2 := "./test/test.log"

	OpenFile(tmp1)
	OpenFile(tmp2)

	if _, err := os.Stat(tmp1); err != nil {
		t.Errorf("%s: file not exists", tmp1)
	}

	if _, err := os.Stat(tmp2); err != nil {
		t.Errorf("%s: file not exists", tmp2)
	}

}
