package main

import (
	"archive/zip"
	"log"
)

// 读取压缩文件

func main() {
	var path = "/Users/yostar/Perforce/NovaMacminiMasterPoj/Version/Baseline/GameDataTables/patch.update.06151650.zip"
	reader, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			log.Printf("directory:%s", file.Name)
		} else {
			log.Printf("file:%s", file.Name)
		}
	}

}
