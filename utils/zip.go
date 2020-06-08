package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Compress a file or folder with ZIP format
func ZipCompress(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	var fullBaseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
		fullBaseDir = source
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			var baseDirPath = fullBaseDir
			if runtime.GOOS == "windows" && strings.Contains(baseDirPath, "/") {
				baseDirPath = strings.ReplaceAll(baseDirPath, "/", "\\")
			}
			if runtime.GOOS == "windows" && strings.Contains(path, "/") {
				path = strings.ReplaceAll(path, "/", "\\")
			}
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, baseDirPath))
			if header.Name == baseDirPath {
				header.Name = ""
			}
		}
		if info.IsDir() {
			if header.Name != "" {
				header.Name += "/"
			}
		} else {
			header.Method = zip.Deflate
		}

		if header.Name == "" {
			return nil
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

// Uncompress zip archive to a given folder
func ZipUnCompress(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer reader.Close()
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

// UnCompress zip archive to a given folder, filtering the file/folder name and place in the output folder, without neasted folders
func ZipUnCompressFilter(archive, target string, filter string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer reader.Close()
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			if strings.Contains(file.Name, filter) {
				os.MkdirAll(path, file.Mode())
			}
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()
		if strings.Contains(path, filter) {
			_, name := filepath.Split(file.Name)
			path = filepath.Join(target, name)
			targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer targetFile.Close()

			if _, err := io.Copy(targetFile, fileReader); err != nil {
				return err
			}
		} else {
			continue
		}
	}
	return nil
}
