package storage

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestSaveFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "files")
	if err != nil {
		t.Fatal(err)
	}
	local, err := NewLocal(dir)
	if err != nil {
		t.Fatal(err)
	}

	err = local.Save("1/test.png", bytes.NewBufferString("Hello World"))
	if err != nil {
		t.Fatal(err)
	}


}
