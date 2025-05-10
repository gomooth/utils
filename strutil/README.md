# strutil 字符串处理工具

## 格式转换

- `strutil.Snake(str)` 驼峰转成蛇形字符串。 
  - XxYy --> xx_yy
  - AaBB --> aa_bb
  - HTMLElement -> html_element
- `strutil.Camel(str)` 蛇形转成驼峰字符串
  - xx_yy --> XxYy
  - XML_http_request --> XMLHttpRequest
