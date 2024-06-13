# Wgee

Wgee 是一个命令行工具，用于从 github 代理网站快速下载 github 的文件，并支持一些可选的配置参数。

## 功能特性

- 从指定的 URL 下载文件。
- 可选地指定输出文件名。
- 可以自定义下载使用的域名。
- 显示当前配置的域名信息。

## 使用方法

要使用 Wgee，请使用以下选项运行可执行文件：
```shell
./wgee [-O 输出文件名] [-c 自定义域名] [-proxy] <URL>
示例：
./wgee http://baidu.com/index.html
程序默认提供一个代理服务器，如需替换可以直接-c 替换
```
### 参数选项

- `-O <输出文件名>`：指定下载的文件的输出文件名。
- `-c <自定义域名>`：使用自定义的域名进行下载，自定义的域名需要以斜杠 `/` 结尾。
- `-proxy`：显示当前配置的域名信息。
- `-help`：显示帮助信息和参数选项的使用说明。

## 示例

1. 下载文件并指定输出文件名：
```shell
./wgee -O output.txt https://example.com/file.txt
```


3. 使用自定义域名进行下载：
```shell
./wgee -c https://customdomain.com/ https://example.com/file.txt
```


5. 显示当前配置的域名信息：
```shell
./wgee -proxy
```

6. 显示帮助信息和参数选项的使用说明：
```shell
./wgee -help
```


## 注意事项

- 如果未提供 `-O` 参数，则默认将文件保存为下载的文件名。
- 使用 `-c` 参数更新后，会持久化保存配置，下次使用时将使用新的配置域名进行下载。
- 项目基于Cloudflare Workers，开源于GitHub hunshcn/gh-proxy
