# valutil 值处理工具

## 断言

- `valutil.IsNil(any)` 判断任意类型的值是否为 `Nil`


## 类型强制转换

### Boolean

- `valutil.Bool(any)` 将任意值转成 `bool`，转换失败则返回 false
- `valutil.BoolWith(any, defaulted)` 将任意值转成 `bool`，转换失败则返回 defaulted
- `valutil.MustBool(any)` 将任意值转成 `bool`，转换失败则返回错误

转换规则如下：

1. 如果传入值是 boolean，直接强制转换返回；
2. 如果传入值是 string，则按以下规则转换：

	- "true, yes, y" -> true
	- "false, no, n" -> false
	- "" -> false
	- "0.0···001 ... 1 ... ∞" -> true
	- "-∞ ... -1 ... -0.1 ... 0" -> false
	- "other word" -> ERROR

3. 如果传入值是 数字，大于零返回 true，否则返回 false

    - 0.0···001 ... 1 ... ∞ -> true
    - -∞ ... -1 ... -0.1 ... 0 -> false

4. 如果传入值是其他类型，则返回 ERROR

### Integer

- `valutil.Int(any)` 将任意值转成整数，转换失败则返回 false
- `valutil.IntWith(any, defaulted)` 将任意值转成整数，转换失败则返回 defaulted
- `valutil.MustInt(any)` 将任意值转成整数，转换失败则返回错误

规则如下：

1. 如果输入值是 数字，则直接转换；
2. 如果输入值是 `boolean`，则 `true` 转为 1；`false` 转为 0；
3. 如果输入值是 `string`，则按字符串转换规则，并支持以下字符转换：

    - "true", "yes", "y" --> 1
    - "false", "no", "n" --> 0


## 值和指针转换

- `valutil.Ptr(any)` 返回值的指针
- `valutil.Val(ptr)` 返回指针的值
- `valutil.SlicePtr` 将值切片转换为指针切片
- `valutil.SliceVal` 将指针切片转换为值切片


## 随机数

- `valutil.Rand(min, max)` 生成 [min, max) 范围内的随机数
- `valutil.SecureRand(min, max)` 生成 [min, max) 范围内的加密安全随机数
- `valutil.RandAlphanum(num)` 生成 num 长度的随机字符串
- `valutil.RandSafeAlphanum(num)` 生成 num 长度的易辨识的随机字符串。适合用于验证码、密码等需要人工识别的场景
