package gomvc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

// CreateFileFromString takes a filepath as the destination of the file
// to be created as well as the contents to be written to this file.
func CreateFileFromString(filepath string, contents string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return errors.Wrap(err, "CreateFileFromString: os.Create error")
	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(contents)
	w.Flush()

	if err != nil {
		return errors.Wrap(err, "CreateFileFromString: write string error")
	}
	return nil
}

func createStringFromFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func createDirIfNotExists(dir string) {
	if !dirExists(dir) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			panic(err)
		}
		log.Printf("created %s\n", dir)
	}
}

func dirExists(path string) bool {
	i, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return i.IsDir()
}

func fileExists(path string) bool {
	i, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !i.IsDir()
}

func addGoExt(s string) string {
	return fmt.Sprintf("%s.go", s)
}
