package main

import (
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		print(e)
	}
}

func main() {

	err := os.Mkdir("FileDirsTempDir", 0775)
	check(err)
	defer os.RemoveAll("FileDirsTempDir")

	// tempfile
	f, err := os.CreateTemp("", "sample")
	check(err)
	defer os.Remove(f.Name())
	println("Temp file name:", f.Name())
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// tempdir
	dname, err := os.MkdirTemp("", "sampledir")
	println("Temp dir name:", dname)
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)

}
