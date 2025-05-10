# userutil 用户相关处理工具

## 密码

- `userutil.NewHasher()` 密码加密器
  - `hasher.Sum(str)` 密码加密
  - `hasher.Check(input, cryptoText)` 密码校验
