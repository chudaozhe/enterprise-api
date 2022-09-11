# 企业展示型小程序 - 服务端
Gin + Mysql + Nginx + Redis

如果需要`管理员端`或`小程序端`，或表结构，请移步：https://github.com/chudaozhe/gin-vue-weapp

## 主要功能
- 登陆
- 管理员管理：添加管理员（密码以邮件形式发送），修改头像，重置密码（密码以邮件形式发送），启用/禁用，删除
- 图片空间：所有的图片集中管理
- 分类管理：文章分类，案例分类，单页分类
- 成功案例：支持搜索，缩略图，图集，显示/隐藏，删除
- 文章发布：支持搜索，缩略图，图集，显示/隐藏，删除
- 单页管理：适合做`关于我们`，`联系我们`等简单的介绍
- 快捷方式管理：图标加文字，点击可导航到指定页面（小程序首页焦点图下面的几个方块）
- 焦点图：图片和链接（小程序首页顶部出现的可切换图片）

## 底层支持
- Gin
- GORM
- Mysql
- Nginx
- Redis缓存
- SMTP发送邮件
- 中间件（api鉴权，跨域，logrus日志）

## 环境模式
三种环境模式：`debug`、`release`、`test`，分别对应3个配置文件
```
app/config/app-debug.json
app/config/app.json
app/config/app-test.json
```

在开始之前，你需要设置一种模式
```
app/config/env
```
如果不修改，则默认为`debug`模式

### 几个关键配置

1、发邮件配置
```
  "smtp_config": {
    "host": "smtp.mxhichina.com换成你自己的",
    "port": ":465",
    "ssl_port": ":25",
    "username": "aa@xxx.com换成你自己的",
    "password": "换成你自己的",
    "ssl": true
  },
```

2、redis配置
```
  "redis_config": {
    "host": "docker-redis",
    "port": ":6379",
    "password": "",
    "db": 0
  },
```

3、mysql配置
```
  "database": {
    "dsn": "root:@tcp(docker-mysql:3306)/ent?charset=utf8&parseTime=True&loc=Local&timeout=10ms"
  },
```

4、端口配置
```
  "app_port": ":7097",
```

5、域名配置
```
  "app_host": "//localhost",
```

## 快速开始
1、安装GO环境，IDE推荐GoLand

2、下载代码，在根目录执行
```
go mod tidy
```
![debug.jpg](https://ent.cuiwei.net/screenshots/backend/debug.jpg)

3、创建docker镜像

需要先切换到`release`或`test`模式
```
docker build -t ent:1.0-test .
```
