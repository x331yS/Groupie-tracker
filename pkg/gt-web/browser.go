package gt_web

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func Openbrowser(url string) {
	var err error
	//Get Pc os
	switch runtime.GOOS {
	//Open browser with command line
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	case "linux":
		err = exec.Command("xdg-open", url).Start()

	default:
		fmt.Println("Can't find your os, please open yourself localhost:1111")

	}
	if err != nil {
		log.Fatal(err)
	}
}
