# 全球地震信息爬虫程序与实时推送

### 已实现功能
- 可从Yahoo!防災速報等网站爬取最近发生的地震及相关情报，写入SQLite数据库
- 可爬取日本实时EEW信息并写入SQLite数据库
- 可对接Redis等实现消息队列，对接其他应用实现地震信息实时推送（自行扩展）
- 基于Go语言编写，可编译后运行在嵌入式设备等性能受限的设备上，路由器也能跑
- 可通过邮件、Telegram向预设的用户推送实时信息
- 高度灵活的配置

### 预期实现功能
- 可爬取中国地震台网地震信息、四川省地震局EEW、福建省地震局EEW、中国台湾中央气象署EEW等
- 实现Web API供外部查询信息
- 实现Telegram Bot查询信息

### 测试
1. 自行安装Go工具链
2. 进入根目录
3. 运行以下命令
    ```shell
    go build -ldflags="-s -w" -o earthquake-crawler cmd/main.go
    ```
4. 根据测试结果调整代码

### 编译与运行

1. 自行安装Go工具链并下载相关依赖
2. 进入根目录
3. 运行以下命令
    ```shell
    go build -ldflags="-s -w" -o earthquake-crawler cmd/main.go
    ```
4. 需要交叉编译的情况(以目标平台linux/arm64为例)
    ```shell
    GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o earthquake-crawler cmd/main.go
    ```
5. 使用命令行运行，首次运行会自动生成空配置文件于data/config.toml，请自行修改配置

### Docker
1. 不推荐使用Docker部署
2. 如需使用，源码提供了Dockerfile，可自行打包镜像，建立容器完成部署
3. 建议通过-v参数将数据库与配置从本地挂载至容器，以便调试配置与读取数据

