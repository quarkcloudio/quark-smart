# QuarkSmart

Overview:

QuarkSmart is a golang single service based on QuarkGo

Install:

1、重命名.env.example 改为 .env 

2、编辑.env文件，更改配置信息

3、执行下面的命令完成安装：
``` bash
# 第一步，安装依赖:
go mod tidy

# 第二步，创建vendor目录
go mod vendor

# 第三步，启动服务：
go run main.go
```

后台地址： http://127.0.0.1:3000/admin/

默认用户名：administrator 密码：123456