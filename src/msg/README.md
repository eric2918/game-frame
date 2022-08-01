## 消息ID规则

- 错误码在 `protocol` 中进行定义。

#### 消息ID为 5 位数

| 1 | 01 | 01 |
| :------ | :------ | :------ |
| 服务模块 | 功能模块 | 具体消息 |

- 服务模块：1 位数进行表示，比如 1:静态资源，2:网关服务，3:账号服务
- 功能模块：2 位数进行表示，比如 01:养成功能，02:地图功能
- 具体消息：2 位数进行表示，比如 01:角色升级请求；02:角色升级响应


#### 服务模块初始值

| 服务模块 | 起始值 | 
| :------ | :------ |
| 静态资源 | 10000 |
| 网关服务 | 20000 |
| 账号服务 | 30000 |
| 游戏服务 | 40000 |
| 聊天服务 | 50000 |