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
	err = utils.TarCompress("C:/Users/Fabrizio/go/src/github.com/hellgate75/k8s-deploy", "C:/Users/Fabrizio/go/src/github.com/hellgate75/k8s-deploy.tgz", true)
	if err != nil {
		fmt.Printf("Error tarring folder: %v", err)
	}
}
