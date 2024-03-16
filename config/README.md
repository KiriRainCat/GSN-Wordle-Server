# 配置文件说明

文件名: `config.yaml`

## 文件内容示例:

```yaml
# cSpell: disable

server: # 服务器设置
  port: 8005 # 服务器端口
  request_auth: "qwq" # 接口请求鉴权 (Authorization)

postgresql: # PostgreSQL 数据库设置
  dev_host: "127.0.0.1" # 开发环境连接地址
  dev_db: "jozutxqg" # 开发环境数据库名
  dev_user: "name" # 开发环境用户名
  dev_password: "pwd" # 开发环境密码
  host: "host" # 连接地址
  port: 5432 # 端口
  database: "name" # 数据库名
  user: "name" # 用户名
  password: "pwd" # 密码

redis: # Redis 数据库设置
  dev_host: "host" # 开发环境连接地址
  dev_port: 15254 # 开发环境端口
  dev_user: "name" # 开发环境用户名
  dev_password: "pwd" # 开发环境密码
  dev_db: 0 # 开发环境数据库
  host: "host" # 连接地址
  port: 11813 # 端口
  password: "pwd" # 密码
  db: 1 # 数据库
```
