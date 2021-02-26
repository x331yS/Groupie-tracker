package gt_web

import (
	"log"
	"os/exec"
	"runtime"
)

func Openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	case "linux":
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		log.Fatal(err)
	}
}
