package filesutils

import (
	"bytes"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const fileMode = 0666

// Truncates a file and writes to it.
// If the path to the file does not exist WriteToFile creates it
func WriteToFile(path string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModeDir)
	if err != nil {
		return err
	}

	return os.WriteFile(path, content, fileMode)
}

type FileFilter func(path string) bool

// Walks through a directory (recursive) and filters files according to `filter`
func WalkDirectoryAndFilter(root string, filter FileFilter) (files []string, err error) {
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalln(err)
		}
		if !d.IsDir() && filter(path) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

// Receives the content of a YAML file and returns an array of the documents
// in it
func SplitYamlDocuments(fileContent []byte) ([]([]byte), error) {
	reader := bytes.NewReader(fileContent)
	decoder := yaml.NewDecoder(reader)
	var docs []([]byte)

	for {
		var d interface{}
		err := decoder.Decode(&d)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		dMarshal, err := yaml.Marshal(d)
		if err != nil {
			return nil, err
		}
		docs = append(docs, dMarshal)
	}

	return docs, nil
}

func CopyFile(copyFrom string, copyTo string) error {
	content, err := os.ReadFile(copyFrom)
	if err != nil {
		return err
	}

	return WriteToFile(copyTo, content)
}
