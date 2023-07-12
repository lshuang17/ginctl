### 使用方式
#### 1. go install
```shell
go install github.com/chnls/ginctl
```

#### 2. source
1. 生成命令二进制文件
    ```shell
    go build .
    ```
2. 将二进制文件放入$GOPATH/bin目录
3. 使用命令生成模块文件
   ```markdown
    ginctl init your_project [mod_name.com]
   ```
   
   ```markdown
    ginctl new [-di -u username] demo [pkgName]
    参数解析:
        -di: 注入依赖方式，google/wire，默认不使用
        -u: 文件创建人
        demo 生成模块
        pkgName: 包名，不指定默认模块名
   ```
   > `ginctl new -h` for help

