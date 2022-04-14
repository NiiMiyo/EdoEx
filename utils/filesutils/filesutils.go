package filesutils

import (
	"archive/zip"
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

//
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

func ZipFiles(outputPath string, files map[string]string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	for filePath, pathInZip := range files {
		err := addToZipWriter(zipWriter, filePath, pathInZip)
		if err != nil {
			return err
		}
	}

	return nil
}

func addToZipWriter(zipWriter *zip.Writer, filePath string, pathInZip string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.IsDir() {
		files, err := WalkDirectoryAndFilter(
			filePath,
			func(string) bool { return true },
		)
		if err != nil {
			return err
		}

		for _, f := range files {
			zipPath := filepath.Join(pathInZip, f[len(filePath):])
			err = addToZipWriter(zipWriter, f, zipPath)
			if err != nil {
				return err
			}
		}

		return nil
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = pathInZip
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}
