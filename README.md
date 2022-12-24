# Q-A

CSA初级考核内容:问答社区

# 目录

- ### Q-A

  - 目录
  - 功能实现
    - 基础功能
  - 接口说明
    - 用户
      - 注册POST
      - 登陆POST
      - 改密码POST
      - 删除账号DELETE
    - 问题
      - 创建问题POST
      - 获取所有回答GET
      - 获取回答详情GET
      - 修改问题POST
      - 点赞POST
      - 取消点赞DELETE
      - 删除问题DELETE
    - 回答
      - 创建回答POST
      - 修改回答POST
      - 删除回答DELETE

## 功能实现

### 基础功能

- 用户注册登陆修改密码删除账号
- 创建修改删除问题
- 点赞/取消点赞
- 创建修改删除回答
- 获取回答详情
- 使用mysql存储所有数据

## 接口说明

[Q-A - Team Workspace (postman.co)](https://cannat-csa-primary.postman.co/workspace/e6fe6b10-2c29-44b9-88a9-482257117e97/collection/25061249-b3e9f6ba-f679-41c0-8784-ecb29c14c71a)

### 注册POST

> localhost:8080/register

#### BODY

| KEY      | DESCRIPTION |
| -------- | ----------- |
| uesrname | 必填        |
| password | 必填        |

### 登录POST

> ##### localhost:8080/login

#### BODY

| KEY      | DESCRIPTION |
| -------- | ----------- |
| uesrname | 必填        |
| password | 必填        |

### 修改密码POST

> ##### localhost:8080/user/password

#### BODY

| KEY          | DESCRIPTION |
| ------------ | ----------- |
| uesrname     | 必填        |
| old_password | 必填        |
| new_password | 必填        |

### 删除DELETE

> ##### localhost:8080/delete

#### BODY

| KEY      | DESCRIPTION |
| -------- | ----------- |
| uesrname | 必填        |
| password | 必填        |

### 创建问题POST

> localhost:8080/question

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

#### BODY

| KEY  | DESCRIPTION |
| ---- | ----------- |
| txt  | 必填        |

### 获取所有问题GET

> localhost:8080/question

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

### 获取问题详情GET

> localhost:8080/question/:question_id

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

#### Param

| KEY         | DESCRIPTION |
| ----------- | ----------- |
| question_id | 必填        |

### 修改问题POST

> localhost:8080/question/:question_id

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

#### 路由参数

| KEY         | DESCRIPTION |
| ----------- | ----------- |
| question_id | 必填        |

#### BODY

| KEY     | DESCRIPTION |
| ------- | ----------- |
| new_txt | 必填        |

### 点赞POST

> localhost:8080/question/like?question_id=2

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

#### PARAM

| KEY         | DESCRIPTION |
| ----------- | ----------- |
| question_id | 必填        |

### 取消点赞DELETE

> localhost:8080/question/like?question_id=2

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

#### PARAM

| KEY         | DESCRIPTION |
| ----------- | ----------- |
| question_id | 必填        |

### 删除问题DELETE

> localhost:8080/question/:question_id

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

#### 路由参数

| KEY         | DESCRIPTION |
| ----------- | ----------- |
| question_id | 必填        |

### 创建回答POST

> localhost:8080/answer

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

#### BODY

| KEY         | DESCRIPTION |
| ----------- | ----------- |
| txt         | 必填        |
| question_id | 必填        |

### 修改回答POST

> localhost:8080/answer/update

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

#### BODY

| KEY       | DESCRIPTION |
| --------- | ----------- |
| new_txt   | 必填        |
| answer_id | 必填        |

### 删除回答POST

> localhost:8080/answer/:answer_id

#### Cookies

| KEY      | DESCRIPTION |
| -------- | ----------- |
| username | 必填        |

#### 路由参数

| KEY       | DESCRIPTION |
| --------- | ----------- |
| answer_id | 必填        |

