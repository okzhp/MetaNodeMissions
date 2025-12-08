# 一个基于go实现的博客系统后端服务，使用gorm和gin框架，提供基本的注册登录和增删改查功能。

## 1、包含功能接口：

1. 用户注册、登录
2. JWT鉴权
3. 发布/编辑/删除文章、发布/删除评论
5. 查询文章及评论、查询文章评论

## 2、运行方法
1. 下载项目同步依赖后，修改db.json文件中的数据库配置，执行go run main.go即可
> 注意：发布/编辑/删除文章、发布/删除评论 这些接口需要使用jwt的token才能请求。
> 调用注册接口注册成功后，使用注册的账户调用登录接口会返回一个token，后续请求需要在请求头中加入:
> Authorization: token值

## 3、目录结构说明

```
/mission4
    /controller ---------->控制器层,负责绑定参数，并调用service处理请求
        blog_controller.go
        user_controller.go
    /dto ----------------->数据传输层
        blog_request.go
        common_response.go
        user_response.go
    /model --------------->数据库模型
        comment.go
        post.go
        user.go
    /service ------------->具体的服务实现层，负责处理业务逻辑
        blog_service.go
        user_service.go
    /util ---------------->工具类
        crypto.go
    main.go
```
