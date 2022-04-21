package filesutils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

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
