# kss-cli

## target

创建常用的一些列操作的 cli 工具，通过命令方式能快速实现自己的的需求。

## list

[x] 本地创建一个 git 仓库，并且将远程仓库和本地仓库关联起来。
[ ] release 版本

## 命令

``` bash
go build -o kss-cli./cmd/app/main.go
```

## 使用

运行二进制文件后，会在用户目录下生成一个 .kssrc 文件，用于存储配置信息，修改成自己的信息即可

## 构建所有版本

``` bash
./build_all.sh
```
