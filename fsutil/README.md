# fsutil 文件处理工具

## 大文件处理

- `fsutil.BlockRead(filename, handle)` 分块读取大文件，每次读取 4M 内容

## 文件下载

- `fsutil.Download` 下载远程文件

## 文件及目录处理

- `fsutil.Exist(filename)` 判断文件是否存在
- `fsutil.PathExist(path)` 判断路径/目录是否存在
- `fsutil.IsDir(path)` 判断所给路径是否为文件夹

- `fsutil.Clear(filepath)` 清理目录下的所有文件

- `fsutil.Copy(src, dst)` 复制文件，并自动创建不存在的目录，如果文件已存在则覆盖
- `fsutil.CopyFile(src, dst, overview)` 复制文件，并自动创建不存在的目录
- `fsutil.CopyPath(src, dst)` 递归复制目录，并自动创建不存在的目录，如果文件已存在则覆盖

- `fsutil.Read2Base64(path string)` 读文件并返回内容的 `base64` 编码，一般用在图片处理时。

- `fsutil.QiNiuFileHash(filename)` 获得文件内容的 hash
- `fsutil.QiNiuHash(reader, size)` 获得指定长度内容的 hash


