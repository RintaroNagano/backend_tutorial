# sign-server

## 概要
＠SaKu2110まで  

## 実装した機能
- [GET] /ping  
  - Request: なし
  - Response: {"message": "ping"} / HttpStatusCode(200)
- [POST] /signin
  - Request: Header: UserId `string`, Passwrod `string`
  - Response: {"token": 任意の文字列} / HttpStatusCode(200)

- [POST] /signup
  - Request: Header: UserId `string`, Passwrod `string`
  - Response: {"token": 任意の文字列} / HttpStatusCode(201)

## 実行手順
【任意】`make run`でapiを起動する