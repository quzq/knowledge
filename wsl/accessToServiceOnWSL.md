
windowsからのWSL上のサービスへのアクセス

### tl;dr
WSL（またはWSL2）内で動作するサービスへのアクセス方法です。
`localhost`を使用することで簡単にアクセスできます。
ただし、`127.0.0.1`ではアクセスできません。

### WSL内のサービスを立ち上げる
ここでは、node+expressのテンプレートを使い、httpサーバーを起動させます。
※WSLにnodejsとgitがインストールされている必要があります。

WSLを開き、以下のコマンドを実行し、httpサービスを起動させます。
```
$ git clone https://github.com/quzq/play_express.git --depth 1
$ cd play_express
$ npm i
$ npm start
```
成功すると、以下のように出力されます。

```
Listening on port 3000
```

### Power Shell からアクセス
Windows Power Shell を起動し、上記のサービスへアクセスします。

```
PS > curl http://localhost:3000

```
httpサービスの頁を取得できました。
```
StatusCode        : 200
StatusDescription : OK
Content           : <!DOCTYPE html><html><head><title>Express</title><link rel="stylesheet" href="/stylesheets/style.cs
                    s"></head><body><h1>Express</h1><p>Welcome to Express</p></body></html>
RawContent        : HTTP/1.1 200 OK
                    Connection: keep-alive
                    Content-Length: 170
                    Content-Type: text/html; charset=utf-8
                    Date: Fri, 03 Jul 2020 07:18:31 GMT
                    ETag: W/"aa-z+ebXSEdArbZ+EXlN/WQjf6HV8c"
                    X-Powered-By: Expre...
Forms             : {}
Headers           : {[Connection, keep-alive], [Content-Length, 170], [Content-Type, text/html; charset=utf-8], [Date,
                    Fri, 03 Jul 2020 07:18:31 GMT]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 170
```

### webブラウザからアクセス
適当なwebブラウザに、`http://localhost:3000`を指定してアクセスします。

以下のようなページが表示されました。
