package main

import (
	"archive/zip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//if err := encode("abc.zip", "/Users/yostar/workSpace/GoNewWork/ThinkLibrary"); err != nil {
	//	panic(err)
	//}
	if err := decode("/Users/yostar/Desktop", "/Users/yostar/Desktop/Update(2).zip"); err != nil {
		panic(err)
	}

}

func encode(dst, src string) error {
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(f)
	//buf := new(bytes.Buffer)

	if err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		name := obtained(src, path)
		w, err := zw.Create(name)
		if err != nil {
			return err
		}
		rf, err := os.Open(path)
		if err != nil {
			return err
		}
		//hash := checkCode(rf)
		//buf.WriteString(fmt.Sprintf("file:%s\tcheckCode:%s\n", name, hash))
		_, err = io.Copy(w, rf)
		return err
	}); err != nil {
		return err
	}
	//verify := src + string(os.PathSeparator) + "verify.txt"
	//w, err := zw.Create(obtained(src, verify))
	//if err != nil {
	//	return err
	//}
	//if _, err = io.Copy(w, buf); err != nil {
	//	return err
	//}
	return zw.Close()
}

func decode(dst, src string) error {
	zr, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("1 %s", err)
	}

	if dst != "" {
		if err = os.MkdirAll(dst, 0755); err != nil {
			return fmt.Errorf("2 %s", err)
		}
	}

	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)
		fmt.Println("path -> ", path)
		if !file.FileInfo().Mode().IsRegular() {
			continue
		}
		if file.FileInfo().Size() == 0 {
			continue
		}
		root := filepath.Dir(path)
		if err := os.MkdirAll(root, 0777); err != nil {
			return fmt.Errorf("37 %s", err)
		}
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return fmt.Errorf("3 %s", err)
			}
			continue
		}
		fmt.Printf("file name: %s file perm --> %v file size: %d \n", file.Name, file.Mode().Perm(), file.FileInfo().Size())
		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("4 %s", err)
		}
		fw, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if err != nil {
			return fmt.Errorf("5 %s", err)
		}

		_, err = io.Copy(fw, rc)
		if err != nil {
			return err
		}

		_ = fw.Close()
		_ = rc.Close()
	}
	return zr.Close()
}

func obtained(src, path string) string {
	str := strings.Split(src, string(os.PathSeparator))
	prefix := strings.Join(str[:len(str)-1], string(os.PathSeparator)) + string(os.PathSeparator)
	return strings.Replace(path, prefix, "", 1)
}

func checkCode(f io.ReadWriteCloser) string {
	h256 := sha256.New()
	if _, err := io.Copy(h256, f); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h256.Sum(nil))
}
