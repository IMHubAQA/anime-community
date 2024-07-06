# 萌喵酱-社区服务

## 简介
提供社区业务后端服务，包括首页帖子、约妆、搜索、评论收藏点赞等场景

## 详细说明

### 技术选型
- 语言：golang
- web框架：beego
- orm: gorm
- 日志：zap
- 数据库：mysql
- 分布式缓存：redis
- 本地缓存：freeCache
- 搜索引擎：elasticsearch
- 消息队列：暂用redis list

### 项目文档
 - [figma](https://www.figma.com/design/cp8KS1Vix605UPezKABcmC/live-chat?node-id=0-1&t=28TBsFaB5ItGOh4X-0)
 - [系统设计文档](https://pet2y9q9b5.feishu.cn/wiki/Wx1ywN1MWiljlVk4L70ctmrVnSd)
 - [系统接口文档](https://pet2y9q9b5.feishu.cn/wiki/RoC4w4XJeiEYdBk8GlHcDUsLnwb)

### 项目结构
``` golang
anime-community
├── common // 工具代码
├── conf // 配置文件
├── controller // 接口层
├── dao // 数据层（包括：mysql、redis、http、rpc）
├── model // 实体类
├── router // 接口路由
└── service // 业务逻辑层
```

### 部署流程

```shell
go build
nohup ./anime-community > /dev/null 2>&1 &
```