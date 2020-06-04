package main

import (
	"fmt"
	"github.com/hellgate75/k8s-deploy/utils"
)

func main() {
	var err error
	//	err = utils.ZipCompress("C:/Users/Fabrizio/go/src/github.com/hellgate75/k8s-deploy", "C:/Users/Fabrizio/go/src/github.com/hellgate75/k8s-deploy.zip")
	//	if err != nil {
	//		fmt.Printf("Error zipping folder: %v", err)
	//	}
	//	err = utils.TarCompress("C:/Users/Fabrizio/go/src/github.com/hellgate75/k8s-deploy", "C:/Users/Fabrizio/go/src/github.com/hellgate75/k8s-deploy.tgz", true)
	//	if err != nil {
	//		fmt.Printf("Error tarring folder: %v", err)
	//	}
	var n int64 = 0
	var folder = ""
	var srcFolder = "C:/Users/Fabrizio/go/src/github.com/hellgate75/k8s-deploy"
	var dstFolder = "C:/Users/Fabrizio/tmp"
	n, folder, err = utils.CopyFileToFolder(srcFolder, dstFolder)
	if err != nil {
		fmt.Printf("Error copying folder %s, Error: %v\n", err)
	}
	fmt.Printf("Copied files: %v, to destination folder: %s\n", n, folder)

}
