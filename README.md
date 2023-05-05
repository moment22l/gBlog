# 博客-gBlog

## 描述

本项目暂时只包含博客系统的后端部分，只提供了博客系统所需的接口，前端部分暂时未着手制作。



## 技术选型

1. web框架: gin
2. 数据库: mysql
3. orm: gorm
4. 缓存: redis
5. 配置文件: yaml.v3
6. 文件存储: 七牛云
7. 用户认证: jwt



## TODO

- [x] 前期准备
  - [x] 配置文件
  - [x] gorm
  - [x] 日志
  - [x] 路由
- [x] swag文档
- [x] 配置管理
- [x] 图片管理
- [x] 广告管理
- [x] 菜单管理
- [x] 用户管理
- [x] 标签管理
- [ ] 文章管理
- [ ] 评论管理



## 项目结构

```
-gBlog
    |-api 接口目录
    |-config 配置目录
    |-core 服务内核目录
    |-docs 文档目录
    |-flag 命令行操作目录
    |-global 全局变量目录
    |-middleware 中间件目录
    |-models 数据库模型目录
    |-plugins 外部插件目录
    |-routers 路由目录
    |-service 公共服务目录
    |-uploads 文件目录
    |-utils 工具目录
    |-conf.yaml 配置文件
    |-main.go 程序入口
```



## 使用说明

### 前期准备

1. 克隆本项目到本地

   ```
   $ git clone https://github.com/moment22l/gBlog.git
   $ cd gblog
   ```

2. 打开本地**mysql**以及**redis**服务

3. 修改**conf.yaml**中的配置数据，以适配本机配置

4. **数据库迁移**，使用如下命令

   ```
   $ go run main.go -db
   ```

5. **创建管理员用户**，使用如下命令并根据提示输入信息

   ```
   $ go run main.go -u admin
   ```

### 启动服务器

```
$ go run main.go
```



## 注意事项

1. 如需使用图片文件上传，请自行到[七牛云官网](https://www.qiniu.com/)申请七牛云存储空间，并修改`conf.yaml`中相关内容
   - `access_key`
   - `secret_key`
   - `bucket` 空间名称
   - `cdn` 外网加速地址
   - `zone` 存储地区
   - `size` 限制图片大小
3. 如需使用邮件订阅，请自行在第三方邮箱申请服务，并修改`conf.yaml`中相关内容
   - `host` smtp服务器地址
   - `port` 服务器端口
   - `user` 发送方邮箱
   - `auth_code` 授权码
   - `default_from_email` 发送方名称

