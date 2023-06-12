package converter

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	FFMPEGEncoder     = "ffmpeg"
	FFMPEGWinBinary   = "/binaries/ffmpeg.exe"
	FFMPEGLinuxBinary = "/binaries/ffmpeg"
	FFMPEGMacBinary   = "/binaries/ffmpeg_osx"
)

func GetEncoderName() string {
	if IsCommandAvailable(FFMPEGEncoder) {
		return FFMPEGEncoder
	} else {
		ffmpegBin, err := AvailableBinary()
		if err != nil {
			panic(fmt.Sprintf("command `%s` not found", FFMPEGEncoder))
		}
		return ffmpegBin
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

func AvailableBinary() (string, error) {
	switch runtime.GOOS {
	case "windows":
		if _, err := os.Stat(FFMPEGWinBinary); err == nil {
			return FFMPEGWinBinary, nil
		}
	case "linux":
		if _, err := os.Stat(FFMPEGLinuxBinary); err == nil {
			return FFMPEGLinuxBinary, nil
		}
	case "darwin":
		if _, err := os.Stat(FFMPEGMacBinary); err == nil {
			return FFMPEGMacBinary, nil
		}
	}
	return "", fmt.Errorf("no binaries found for Operating System %s", runtime.GOOS)
}
