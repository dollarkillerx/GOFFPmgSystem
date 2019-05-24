# VF GO VideoFFMPeg Golang
Golang FFMPeg 云转码

### 总体架构
![](/README/SSR.png)

### 目录结构
```
.
├── api 用户相关服务
│   ├── dbops 数据库交互
│   │   ├── api.go 
│   │   ├── api_test.go
│   │   ├── conn.go 数据库链接
│   │   └── internal.go 内部操作sql相关
│   ├── defs 一切的配置
│   │   ├── apidef.go 
│   │   └── errs.go 错误定义
│   ├── handlers.go 头部处理
│   ├── main.go 
│   ├── response.go 返回定义
│   ├── session session相关
│   │   └── ops.go session的实现与cache
│   └── utils 常用工具库
│       ├── encryption.go 加密解密相关
│       ├── simpleTime.go 时间处理相关
│       ├── tootl_test.go 
│       └── uuid.go 
├── streamserver 视频相关服务
│   ├── defs.go  一些结构定义
│   ├── handlers.go
│   ├── limiter.go 流控模块
│   ├── main.go
│   ├── response.go
│   ├── streamserver 
│   └── VIDEOS //视频文件
│       ├── test1.mp4
│       └── test.mp4
├── scheduler 调度服务
│   ├── dbops 数据库相关
│   │   ├── api.go
│   │   ├── conn.go
│   │   └── internal.go
│   ├── handlers.go
│   ├── main.go
│   ├── response.go
│   ├── scheduler 
│   ├── taskruner  
│       ├── defs.go
│       ├── runner.go 调度器
│       ├── runner_test.go
│       ├── task.go
│       └── trmain.go 定时任务
├── template  模板
│   ├── home.html
│   ├── img
│   ├── script
│   │   └── home.js
│   └── userhome.html
├── web web服务
│   ├── client.go
│   ├── defs.go
│   ├── handlers.go
│   ├── main.go
│   └── web
├── go.mod
├── go.sum
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


#### Stream Server
- Streaming
- Upload files

#### Scheduler Server
- Timer
- 生成/消费模型下的 task runner

流控方案:
bucket 算法
channel. shared channel instead of shared memory


### 大前端
- GO 模板引擎 有两种text/template和html/template
- GO 模板 采用动态生成的模式

