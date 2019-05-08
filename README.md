# VF GO VideoFFMPeg Golang
Golang FFMPeg 云转码

### 总体架构
![](/README/SSR.png)

### API设计

#### 用户
- 创建(注册)用户:URL:/user Method:POST,SC:201创建成功,400参数错误,500内部错误
- 用户登录:URL了/user/:username Method:POST,SC:200,400,500
- 获取用户基本信息:URL:/user/:username Method:GET,SC:200,400,401没有通过验证,403通过验证but没有权限,500
- 用户注销:URL:/user/:username Method:DELETE,SC:204,400,401,403,500