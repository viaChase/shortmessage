# shortmessage

## 注册接口
* URL :127.0.0.1:8081/register

请求参数 :
```json
{
	"user_name":"aaa",
	"pass_ward":"123456"
}
```

成功返回200

## 登入接口

* URL :127.0.0.1:8081/login

请求参数 :
```json
{
	"user_name":"aaa",
	"pass_ward":"123456"
}
```

成功后会在header头 增加
UserId int
Jwt   string 
之后每次访问都需带入这两个header保持登入状态


## 添加通讯录接口

* URL :http://127.0.0.1:8081/home/contacts/add

请求参数 :
```json
{
"contactsId":13987761435,        //朋友id
"contactsName":"xasa"  //备注名
} 
```
## 添加通讯录列表接口

* URL :http://127.0.0.1:8081/home/contacts/list

请求参数 :
```json
{
	"search":"",     //关键词搜索,如果不带就直接搜全部
	"pageSize": 10,  //每次返回最多条数
	"nowPage":1      //当前页数
}
```

返回数据:
```json
{
    "count": 1,
    "contactsList": [
        {
            "friendName": "xasa",
            "friendId": 13987761435
        }
    ]
}
```

## 删除通讯录接口

* URL :http://127.0.0.1:8081/home/contact/delete

请求参数 :
```json
{
	"friendId":13987761435   //删除朋友id
}
```
## 更新通讯录接口

* URL :http://127.0.0.1:8081/home/contact/update

请求参数 :
```json
{
"contactsId":1,        //朋友id
"contactsName":"xasa"  //备注名
}

```

## 备忘录增加接口

* URL :http://127.0.0.1:8081/home/content/add

请求参数 :
```json
{
	"title":"this",
	"content":"this is a new text"
}
```


## 备忘录列表

* URL :http://127.0.0.1:8081/home/content/list

请求参数 :
```json
{
	"search":"", //搜索关键词
	"pageSize":100,
	"nowPage":1
}
```

返回数据 :
```json
{
    "count": 1,
    "contentlist": [
        {
            "contentId": 4,
            "title": "",
            "content": ""
        }
    ]
}
```

## 备忘录修改

* URL :http://127.0.0.1:8081/home/content/update

请求参数 :
```json
{
	"contentId":1,
	"title":"11",
	"content":"22"
}
```

## 备忘录删除

* URL :http://127.0.0.1:8081/home/content/delete

请求参数 :
```json
{
	"id":5
}
```

## 把消息置为已读
自己发送的消息默认已读，所有未读消息必须调用这个接口改成已读状态

* URL :http://127.0.0.1:8081/home/message/read

请求参数 :
```json
{
	"message_id":[] //消息id
}
```

## 删除消息

* URL :http://127.0.0.1:8081/home/message/delete

请求参数 :
```json
{
	"message_id":[] //消息id
}
```



## 发送信息接口

* URL :http://127.0.0.1:8081/home/message/send

请求参数 :
```json
{
	"targetUserId":1,
	"content":"xixi"
}
```


## 轮询 获得消息信息接口

* URL : http://127.0.0.1:8081/home/message/view

请求参数 :
```json
{
	"data_type":1,   // 0 获得所有已读的信息 1 获得未读的信息 拿来轮询
	"searchData":""  //只在dataType 为0的时候有效
}
```

返回数据:
```json
{
    "userMessageList": [
        {
            "userId": 3,
            "userName": "3",
            "messageList": [
                {
                    "messageId": 1,
                    "content": "xixi",
                    "createTime": 1588414560,
                    "isNew": false,
                    "isSelf": false //是否是自己发送的信息
                }
            ]
        }
    ]
}
```


## 当前数据备份

* URL :http://127.0.0.1:8081/home/backup/add

请求参数 :
```json
{
	"data_type":1  //1:对通讯录进行备份 2:对备忘录进行备份 3:对消息进行备份
}
```

## 备份数据的列表

* URL :http://127.0.0.1:8081/home/backup/list

请求参数 :
```json
{
    "count": 1,
    "data": [
        {
            "back_id": 4,
            "create_time": 1588416826,
            "data_type": 1
        }
    ]
}
```

## 备份数据还原

* URL :http://127.0.0.1:8081/home/backup/do

请求参数 :
```json
{
	"id":3
}
```

## 备份数据删除

* URL :http://127.0.0.1:8081/home/backup/delete

请求参数 :
```json
{
	"id":3
}
```