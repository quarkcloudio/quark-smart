kitex工具

生成客户端代码:

```bash

 # 在项目根目录运行
kitex -module github.com/quarkcms/quark-hertz ./idl/admin.thrift

# 在一个服务里同时提供多个 service 时，可以使用 -combine-service 选项
kitex --combine-service -module github.com/quarkcms/quark-hertz ./idl/admin.thrift
```

生成服务端代码:

```bash

# 在项目应用目录运行，例如：cmd/admin/rpc
kitex -module github.com/quarkcms/quark-hertz -service admin -use github.com/quarkcms/quark-hertz/kitex_gen ../../../idl/admin.thrift

# 在一个服务里同时提供多个 service 时，可以使用 -combine-service 选项
kitex --combine-service -module github.com/quarkcms/quark-hertz -service admin -use github.com/quarkcms/quark-hertz/kitex_gen ../../../idl/admin.thrift
```