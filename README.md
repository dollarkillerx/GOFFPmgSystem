# VF GO VideoFFMPeg Golang
Golang FFMPeg 云转码

### 总体架构
![](/README/SSR.png)

### 目录结构
```
├── api
│   ├── dbops 数据库交互
│   ├── defs 一切的配置
│   ├── handlers.go
│   └── main.go 
├── README.md

```

### API设计

##### 用户
- 创建(注册)用户:URL:/user Method:POST,SC:201创建成功,400参数错误,500内部错误
- 用户登录:URL了/user/:username Method:POST,SC:200,400,500
- 获取用户基本信息:URL:/user/:username Method:GET,SC:200,400,401没有通过验证,403通过验证but没有权限,500
- 用户注销:URL:/user/:username Method:DELETE,SC:204not content,400,401,403,500

##### 用户资源
- List all videos:URL:/user/:username/videos Method:GET,SC:200,400,500
- Get one video:URL:/user/:username/videos:vid-id Method:GET,SC:200,400,500
- Delete one video:URL:/user/:username/videos/:vid-id Method:DE;ETE,SC:204,400,401,403,500

##### 评论
- Show comments:URL:/videos/:vid-id/comments Method:GET,SC:200,400,500
- Post a comment:URL:/videos/:vid-id/comments Method:POST,SC:201,400,500
- Delete a comment:URL:/videos/:vid-id/comment/:comment-id Method:DELETE,SC:204,400,401,403,500