# apiエンドポイント
## /auth
### /sineUp
```
request-param
.email
.pass(8文字以上)
.username
```
```
return-param
.error
```
```
curl -X POST -d 'mail=test@test.com' -d 'password=11111111' -d 'username=test1' localhost:8080/api/signup
```
### /signin
__いったん無し__

```
request-param
.uid
```
```
return-param
.error
```
### /signout
__いったん無し__
```
request-param
```
```
return-param
.error
```

## /user
### /getUserInfo
- ログインしているuserの全情報
```
request-param
.uid
```
```
return-param
.mail
.pass
.default_col1
.default_col2
.default_col3
```
### /updateUserInfo
- userテーブルの情報を更新する時にたたく
```
request-param
.uid
.mail
.pass
.default_col1
.default_col2
.default_col3
```
```
return-param
.done
```

## /room
### /getRoomList
- ログインしているuserが所属しているroom一覧
```
request-param
.uid
```
```
return-param
.rid_list
```
### /joinRoom
- userがroomに参加する時にたたく
```
request-param
.uid
.rid
```
```
return-param
.done
```
### /leaveRoom
- userがroomから抜ける時にたたく
```
request-param
.uid
.rid
```
```
return-param
.done
```
### /makeRoom
- userがroomを作る
```
request-param
.roomname
.password
```
```
return-param
.error
```

## /event
### /getEventList
- 参加可能なevent一覧
```
request-param
.uid
```
```
return-param
.event_list
```
### /startEvent
- eventをスターとするときにたたく
```
request-param
.uid
.rid
```
```
return-param
.done
```
### /closeEvent
- eventを終了する時にたたく
```
request-param
.rid
```
```
return-param
.done
```
### /paticipateEvent
- イベントに参加する時にたたく
```
request-param
.uid
.eid
```
```
return-param
.done
```
### /exitEvent
- イベントから抜けるときにたたく
```
request-param
.uid
.eid
```
```
return-param
.done
```


# psql
```
docker exec -it <containerId> /bin/bash
psql -d commusto
set search_path = commusto;
\d
```