package service

import "os/exec"

func ConvertOGGtoMP3(inputFile string, outputFile string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-vn", "-acodec", "libmp3lame", outputFile)
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
