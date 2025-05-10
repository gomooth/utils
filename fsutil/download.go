package fsutil

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
)

// Download 下载远程文件
func Download(filename, remoteUrl string) error {
	// 检查url
	if _, err := url.Parse(remoteUrl); nil != err {
		return err
	}

	resp, err := http.Get(remoteUrl)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("remote response failed, status code = %d", resp.StatusCode)
	}

	// 创建文件目录
	_ = os.MkdirAll(path.Dir(filename), fs.ModePerm)

	// 创建文件
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	// 保存下载文件
	_, err = io.Copy(out, resp.Body)
	return err
}
