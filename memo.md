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
.uid
.mail
.name
```
```
curl -X POST -d 'uid={uid}' localhost:8080/api/getUserInfo
```

## /default_card
### /getDefaultCard
```
request-param
.uid
```
```
return-param
.error
.name
.hurigana
.birthday
.instagram
.twitter
.facebook
.free
```
```
curl -X POST -d 'uid={uid}' localhost:8080/api/getDefaultCard
```
### /updateDefaultCard
```
request-param
.uid
.name
.hurigana
.birthday
.instagram
.twitter
.facebook
.free
```
```
return-param
.error
```
```
curl -X POST
-d 'uid={uid}'
-d 'colName=テスト太郎'
-d 'colHurigana=てすとたろう'
-d 'colBirthday=2000-01-20'
-d 'colInstagram=http://hoge.com'
-d 'colTwitter=http://fuga.com'
-d 'colFacebook=http://hogefuga.com'
-d 'colFree=よろしく' localhost:8080/api/updateDefaultCard
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
.error
.rooms[
    .rid
    .name
    .password
    .admin
]
```
```
curl -X POST -d 'uid={uid}' localhost:8080/api/getRoomList
```
### /joinRoom
- userがroomに参加する時にたたく
```
request-param
.uid
.rid
.password
```
```
return-param
.error(passwordが違ったとき: isn't match password)
```
```
curl -X POST -d 'rid={rid}' -d 'uid={uid}' -d 'password={password}' localhost:8080/api/joinRoom
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
```
curl -X POST -d 'rid={rid}' -d 'uid={uid}' localhost:8080/api/leaveRoom
```
### /makeRoom
- userがroomを作る
```
request-param
.roomname
.password
.uid
```
```
return-param
.error
.rid
```
```
curl -X POST -d 'roomname=room1' -d 'password=1111' -d 'uid={uid}' localhost:8080/api/makeRoom
```

## /form
### /getForm
- そのroomのformとして登録されているcolListを返す
```
request-param
.rid
```
```
return-param
.colList[
    {col_name},
    {col_name}
]
.error
```
```
curl -X POST -d 'rid={rid}' localhost:8080/api/getForm
```
### /makeForm
- そのroomのformを登録する
```
request-param
.rid
.colList(`,`で区切る)
```
```
return-param
.error
```
```
curl -X POST -d 'rid={rid}' -d 'colList=名前,誕生日,得意教科,好きな食べ物,趣味' localhost:8080/api/makeForm
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