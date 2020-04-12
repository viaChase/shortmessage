# shortmessage

## 注册接口
* URL ：127.0.0.1:8081/register

请求参数 ：
```json
{
	"user_name":"jd",
	"pass_ward":"123456"
}
```

成功返回200

## 登入接口

* URL ：127.0.0.1:8081/login

请求参数 ：
```json
{
	"user_name":"jd",
	"pass_ward":"123456"
}
```

成功后会在header头 增加
Userid int
Jwt   string 
之后每次访问都需带入这两个header保持登入状态


## 添加通讯录接口

* URL ：127.0.0.1:8081/home/add_contacts

请求参数 ：
```json
{
"contactsId":1,        //朋友id
"contactsName":"xasa"  //备注名
} 
```
## 添加通讯录列表接口

* URL ：127.0.0.1:8081/home/contacts_list

请求参数 ：
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
            "friendName": "sss",
            "friendId": 1
        }
    ]
}
```

## 删除通讯录接口

* URL ：127.0.0.1:8081/home/del_contact

请求参数 ：
```json
{
	"friendId":0   //删除朋友id
}
```
## 更新通讯录接口

* URL ：127.0.0.1:8081/home/update_contact

请求参数 ：
```json
{
"contactsId":1,        //朋友id
"contactsName":"xasa"  //备注名
}

```

## 发送信息接口

* URL ：127.0.0.1:8081/home/send_message

请求参数 ：
```json
{
	"targetUserId":1,   //发送对象的id
	"content":"test"    //内容
}

```


## 轮询 获得消息信息接口

请求参数 :
```json
{
	"lastTime":0,     //上次请求信息时间，如果是当登入直接用 0
	"searchData":""   //搜索关键词
}
```

返回数据:
```json
{
    "lastTime": 1586664166,    //下次请求时间
    "userMessageList": [
        {
            "userId": 2,       //消息发送人id
            "userName": "",    //该人的备注名,
            "messageList": [
                {
                    "messageId": 1,
                    "content": "test",
                    "createTime": 1586664166,
                    "isNew": true
                }
            ]
        }
    ]
}
```