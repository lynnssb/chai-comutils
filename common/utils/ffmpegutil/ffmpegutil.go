/**
 * @Author:      wangxuebing
 * @FileName:    ffmpegutil.go
 * @Date         2023/9/2 17:13
 * @Description:
 **/
package ffmpegutil

import (
	"os"
	"os/exec"
)

// GenerationVideoToThumbnail 生成视频缩略图
func GenerationVideoToThumbnail(sourceFilePath, thumbnailPath string) (bool, error) {
	args := []string{
		"-i", sourceFilePath,
		"-ss", "00:00:03",
		"-vframes", "1",
		"-c:v", "mjpeg",
		"-pix_fmt", "yuv420p",
		"-color_range", "tv",
		"-y",
		thumbnailPath,
	}
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
