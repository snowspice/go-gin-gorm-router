# go-gin-gorm-router
    简单示例：go restful demo ,包含路由，添删改查、分页查询等
#  1. gorm语法
    https://blog.csdn.net/lkysyzxz/article/details/79803648
#  2. 目录说明
## 2.1 main.go
       启动函数，初始化服务器，初始化db信息
## 2.2 api/router
    定义路由信息，后续中间件可放到 router.go中。
## 2.3 /api/models
    实体定义，及对实体的操作。
## 2.4 api/database
     数据库连接初始化，当前仅mysql,后续可增加多数据源连接
## 2.5 api/apis
     接口实现方法
## 2.6 api/common
    常用工具类，当前仅分页工具类
## 2.7 api/config
    采用consul作为配置服务器，初始化配置信息
### 2.7.1 安装consul
    https://github.com/snowspice/consul


