package fsutil

import (
	"encoding/base64"
	"os"
)

// Read2Base64 读文件并返回内容的 `base64` 编码。
// 一般用在图片处理时
func Read2Base64(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	sEnc := base64.StdEncoding.EncodeToString(data)
	return sEnc, nil
}
