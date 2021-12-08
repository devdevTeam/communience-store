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

## /default_cards
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
-d 'uid={uid}' \
-d 'colName=テスト太郎' \
-d 'colHurigana=てすとたろう' \
-d 'colBirthday=2000/01/01' \
-d 'colInstagram=http://hoge.com' \
-d 'colTwitter=http://fuga.com' \
-d 'colFacebook=http://hogefuga.com' \
-d 'colFree=よろしく' localhost:8080/api/updateDefaultCard
```

## /room
### /getRoomInfo
- roomの情報
```
request-param
.rid
```
```
return-param
.error
.rid
.name
.haveForm
.hash
```
```
curl -X POST -d 'uid={uid}' -d 'rid={rid}' localhost:8080/api/getRoomInfo
```
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
    .admin
]
```
```
curl -X POST -d 'uid={uid}' localhost:8080/api/getRoomList
```
### /getRoomUsers
- roomに所属するuser一覧
```
request-param
.rid
.haveForm
```
```
return-param
.error
.users[
    .uid
    .displayValue
]
```
```
curl -X POST -d 'rid={rid}' localhost:8080/api/getRoomUsers
```
### /getRoomUserDefaultCard
- roomに所属するuserのdefault_card一覧を取ってくる
```
request-param
.rid
```
```
return-param
.error
.defaultCardList[
    .name
    .hurigana
    .birthday
    .instagram
    .twitter
    .facebook
    .free
]
```
```
curl -X POST -d 'rid=b5a30400-847e-4d06-a266-094035df218c' localhost:8000/api/getRoomUserDefaultCard
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
.haveForm
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
.error
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
.haveForm
```
```
return-param
.error
.rid
```
```
curl -X POST -d 'roomname=room1' -d 'password=1111' -d 'uid={uid}' -d 'haveForm={bool}' localhost:8080/api/makeRoom
```
### /searchRoom
- roomの検索(name or rid)
```
request-param
.name
.rid
```
```
return-param
.error
.rooms[
    .name
    .rid
]
```
```
curl -X POST -d 'name=room1' -d 'rid={rid}' localhost:8080/api/searchRoom
```
### /searchRoomUser
- room内のuser検索(col_nameとvalue)
```
request-param
.rid
.haveForm
.colName
.value
```
```
return-param
.error
.users[
    .uid
    .displayValue
]
```
```
curl -X POST -d 'rid={rid}' -d 'haveForm=true' -d 'colName=名前' -d 'value=田中' localhost:8080/api/searchRoomUser
```
### /checkHash
- hashのチェック
```
request-param
.uid
.hash
```
```
return-param
.error
.rid
.haveForm
```
```
curl -X POST -d 'uid={uid}' -d 'hash={hash}' localhost:8080/api/checkHash
```
### /getRoomAdmin
- roomのAdmin確認
```
request-param
.uid
.rid
```
```
return-param
.admin
```
```
curl -X POST -d 'uid={uid}' -d 'rid={rid}' localhost:8080/api/getRoomAdmin
```

## /card_value
### /getCardValue
- roomのformのuserのデータを取得
```
request-param
.uid
.rid
```
```
return-param
.error
.cardValue [
    (値のリスト)
]
```
```
curl -X POST -d 'uid={uid}' -d 'rid={rid}' localhost:8080/api/getCardValue
```
### /updateCardValue
- card_valueを更新
```
request-param
.uid
.rid
.valueList(`,`で区切る)
```
```
return-param
.error
```
```
curl -X POST -d 'uid={uid}' -d 'rid={rid}' -d 'valueList=田中太郎,お米' localhost:8080/api/updateCardValue
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
    (項目の名前の配列)
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
.displayBoolList(`,`で区切る)
```
```
return-param
.error
```
```
curl -X POST \
-d 'rid={rid}' \
-d 'colList=名前,誕生日,得意教科,好きな食べ物,趣味' \
-d 'displayBoolList=true,false,false,false,false' \
localhost:8080/api/makeForm
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
.events[
    .eid
    .org_uid
    .name
]
```
```
curl -X POST -d 'uid={uid}' localhost:8080/api/getEventList
```
### /startEvent
- eventをスターとするときにたたく
```
request-param
.password
.org_uid
.rid
```
```
return-param
.error
```
```
curl -X POST -d 'password=1111' -d 'org_uid=75POaUuQ97fc2ZfGB0hDOQj8llv2' -d 'rid=709689af-87bb-42f9-8f5e-1844b49596a2' localhost:8080/api/startEvent
```
### /joinEvent
```
request-param
.eid
.uid
.password
```
```
return-param
.error
```
```
curl -X POST -d 'password={password}' -d 'uid={uid}' -d 'eid={eid}' localhost:8080/api/joinEvent
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