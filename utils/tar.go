package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func addFileToTar(tw *tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball
		header := new(tar.Header)
		header.Name = path
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()
		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// copy the file data to the tarball
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}
	return nil
}

// Compress a file or folder with tar format, and evenually compress with g-zip encoding
func TarCompress(source, target string, compress bool) error {
	tarfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarfile.Close()
	var tw *tar.Writer
	if compress {
		gw := gzip.NewWriter(tarfile)
		tw = tar.NewWriter(gw)
		defer func() {
			tw.Close()
			gw.Close()
		}()
	} else {
		tw = tar.NewWriter(tarfile)
		defer func() {
			tw.Close()
		}()
	}
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

	err = filepath.Walk(source, func(path string, fi os.FileInfo, err error) error {
		// return on any error
		if err != nil {
			return err
		}

		// return on non-regular files (thanks to [kumo](https://medium.com/@komuw/just-like-you-did-fbdd7df829d3) for this suggested update)
		if !fi.Mode().IsRegular() {
			return nil
		}

		// create a new dir/path header
		header, err := tar.FileInfoHeader(fi, fi.Name())
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
			header.Name = filepath.ToSlash(header.Name)
		}
		// update the name to correctly reflect the desired destination when untaring
		//header.Name = strings.TrimPrefix(strings.Replace(path, source, "", -1), string(filepath.Separator))
		if fi.IsDir() {
			return nil
		}
		// write the header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// open files for taring
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		// copy path data into tar writer
		if _, err := io.Copy(tw, f); err != nil {
			return err
		}

		// manually close here after each path operation; defering would cause each path close
		// to wait until all operations have completed.
		f.Close()

		return nil
	})
	return err
}

// UnCompress tar archive to a given folder, evenually compressed with g-zip encoding
func TarUnCompress(archive, target string, uncompress bool) error {
	reader, err := os.Open(archive)
	if err != nil {
		return err
	}
	defer reader.Close()
	var tr *tar.Reader
	if uncompress {
		gr, err := gzip.NewReader(reader)
		if err != nil {
			return err
		}
		defer gr.Close()
		tr = tar.NewReader(gr)
	} else {
		tr = tar.NewReader(reader)
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	header, err := tr.Next()
	for header != nil && err == nil {
		path := filepath.Join(target, header.Name)
		if header.FileInfo().IsDir() {
			var mode = os.FileMode(header.Mode)
			_ = os.MkdirAll(path, mode)
			continue
		}
		var mode = os.FileMode(header.Mode)
		size := header.Size
		var b = make([]byte, size)
		_, err := tr.Read(b)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(path, b, mode); err != nil {
			return err
		}
		header, err = tr.Next()
	}

	return nil
}

// UnCompress tar archive to a given folder, evenually compressed with g-zip encoding, filtering the file/folder name and place in the output folder, without neasted folders
func TarUnCompressFilter(archive, target string, uncompress bool, filter string) error {
	reader, err := os.Open(archive)
	if err != nil {
		return err
	}
	defer reader.Close()
	var tr *tar.Reader
	if uncompress {
		gr, err := gzip.NewReader(reader)
		if err != nil {
			return err
		}
		defer gr.Close()
		tr = tar.NewReader(gr)
	} else {
		tr = tar.NewReader(reader)
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	header, err := tr.Next()
	for header != nil && err == nil {
		if strings.Contains(header.Name, filter) {
			path := filepath.Join(target, header.Name)
			if header.FileInfo().IsDir() {
				if strings.Contains(header.Name, filter) {
					var mode = os.FileMode(header.Mode)
					_ = os.MkdirAll(path, mode)
				}
				continue
			}
			var mode = os.FileMode(header.Mode)
			size := header.Size
			var b = make([]byte, size)
			_, err := tr.Read(b)
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(path, b, mode); err != nil {
				return err
			}
		}
		header, err = tr.Next()
	}

	return nil
}
