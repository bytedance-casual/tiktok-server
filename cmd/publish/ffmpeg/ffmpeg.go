package ffmpeg

import (
	"bytes"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

// GetVideoShot 截取第一帧视频帧
func GetVideoShot(shotPath string, videoPath string) (err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		//Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output(shotPath, ffmpeg.KwArgs{"vframes": 1, "format": "image2", "codec": "png"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return err
	}
	return nil
}
