reactアプリをhttpsで公開する

# 確認環境
OS: Amazon Linux release 2 (Karoo)

# nginxのインストール
Amazon Linux 2 では標準で Nginx の YUM 向けパッケージが提供されていないため、extraリポジトリからインストールします。
```
$ sudo amazon-linux-extras install nginx1.12
```
バージョンの確認

```
$ nginx -v
nginx version: nginx/1.12.2
```

nginxの起動
```
$ sudo service nginx start
```

# nginxへのSSL設定
`/etc/nginx/conf.d`に、以下のファイルを配置します。
- `conf`ファイル(nginx設定ファイル)
- `crt`ファイル(サーバー証明書ファイル)
- `csr`ファイル(公開鍵ファイル)
- `key`ファイル(秘密鍵ファイル)

### オレオレ証明書の作成
```
$ cd ~ 
$ openssl genrsa 2048 > server.key
$ openssl req -new -key server.key > server.csr
$ openssl x509 -days 3650 -req -signkey server.key < server.csr > server.crt
$ sudo mv server.* /etc/nginx/conf.d/
$ cd /etc/nginx/conf.d
$ sudo chown root:root server.*
```
※`server.scr`作成時の答えは全て`enter`で回答しました。

### 設定ファイルの追加
設定ファイルを追加します。ここでは、`my_ssl_app`という名前にしました。
ちなみに、`/etc/nginx/conf.d/*.conf`は、`/etc/nginx/nginx.conf`から呼び出されます。ワイルドカード（`*.conf`）で呼び出されるため、名前は何でも構いません。
```
$ cd /etc/nginx/conf.d
$ touch my_ssl_app.conf
```
`my_ssl_app.conf`を以下のように編集します。

```
server {
    listen       443 ssl;
    server_name  localhost;
    ssl_certificate      /etc/nginx/conf.d/server.crt;
    ssl_certificate_key  /etc/nginx/conf.d/server.key;
    location / {
        proxy_pass http://127.0.0.1:3000;
    }
}
```

### nginxの再起動
```
sudo nginx -s reload
```

### 👀ブラウザで確認
ブラウザからhttps接続して確認すると、502 Bad Gateway のエラーページが表示されます。これは、reactアプリが起動していないためです。
(bad_gataway.png)

# reactアプリの作成

### nodejsのインストール（なければ）
```
$ curl --silent --location https://rpm.nodesource.com/setup_10.x | sudo bash -
$ sudo yum -y install nodejs
```

### reactアプリの作成
```
$ npx create-react-app my_app
```
### アプリの起動
```
$ cd my_app
$ npm start
```
※最新(2020/1/7時点)のreact-scriptでssl接続時に問題が発生しているようです。react-scriptが3.3.0だった場合、[リンク先](https://qiita.com/uemuram/items/a8f727b780bfd115bbaf)の対応（react-scriptの3.2.0へのバージョンダウン）を行ってください。

### 👀ブラウザで確認
ブラウザからhttps接続して確認すると、以下のページが表示されます。
(hello_react.png)

以上。