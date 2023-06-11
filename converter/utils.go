package converter

import (
	"fmt"
	"runtime"

	"os/exec"
)

func GetEncoderName() string {
	if IsCommandAvailable(FFMPEGEncoder) {
		return FFMPEGEncoder
	} else {
		panic(fmt.Sprintf("command `%s` not found", FFMPEGEncoder))
	}
}

func IsCommandAvailable(name string) bool {
	var wh string
	if runtime.GOOS == "windows" {
		wh = "where"
	} else {
		wh = "which"
	}
	cmd := exec.Command(wh, name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
