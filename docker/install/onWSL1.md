WSL + Ubuntu 上でdocker-composeを動作させる

# tl;dr
WSLにインストールしたdockerから、「Docker for Windows」デーモンにリモート接続します。

# 確認環境
📀WSL with Ubuntu 18.04LTS

# Docker for Windows のインストール
[ここ](https://qiita.com/fkooo/items/d2fddef9091b906675ca)等を参考に、Docker for Windowsをインストールします。
インストール後、リモート接続を許可するため、Docker for Windows の Settings画面 から、`General` -> `Expose daemon on tcp://localhost:2375 without TLS`のチェックをONにします。 
(画像)

# docker のインストール
[公式](http://docs.docker.jp/engine/installation/linux/docker-ce/ubuntu.html)の通りに、Docker CEをインストールします。

Docker の公式 GPG 鍵を追加します。

```
$ sudo apt-get update
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

鍵の fingerprint（フィンガープリント）が `9DC8 5822 9FC7 DD38 854A E2D8 8D81 803C 0EBF CD88` と表示されるのを確認します。

```
$ sudo apt-key fingerprint 0EBFCD88

pub   rsa4096 2017-02-22 [SCEA]
      9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid           [ unknown] Docker Release (CE deb) <docker@docker.com>
sub   rsa4096 2017-02-22 [S]
```

Dockerの stable （安定版）リポジトリをセットアップします。

```
$ sudo -E add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```
Docker CE をインストールします。

```
$ sudo apt-get install docker-ce
```

# Dockerデーモンへのリモート接続設定
homeディレクトリの`.bash_profile`に以下の行を追加。

```
export DOCKER_HOST=tcp://localhost:2375
```
編集を終えたら、以下の`home`ディレクトリにて以下のコマンドを実行し、環境変数の再読み込みを行います。

```
source .bash_profile
```

# docker-compose のインストール
docker-ompose は複数のコンテナを使う Docker アプリケーションを定義・実行するためのツールです。

```
sudo -i
curl -L https://github.com/docker/compose/releases/download/1.25.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
```
※`/download/`の後に続くバージョン指定は、[公式サイト](http://docs.docker.jp/compose/toc.html)を確認の上最新のバージョンへ変更してください。

# 動作確認
👀以下のコマンドで確認します。

```
$ curl -O https://keinos.github.io/hello-docker-compose/docker-compose.yml
$ docker-compose up
```

成功すると以下のように表示されます。

```
Starting hello_hello_1 ... done
Attaching to hello_hello_1
hello_1  |
hello_1  | Hello from Docker!
hello_1  | This message shows that your installation appears to be working correctly.
hello_1  |
hello_1  | To generate this message, Docker took the following steps:
hello_1  |  1. The Docker client contacted the Docker daemon.
hello_1  |  2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
hello_1  |     (amd64)
hello_1  |  3. The Docker daemon created a new container from that image which runs the
hello_1  |     executable that produces the output you are currently reading.
hello_1  |  4. The Docker daemon streamed that output to the Docker client, which sent it
hello_1  |     to your terminal.
hello_1  |
hello_1  | To try something more ambitious, you can run an Ubuntu container with:
hello_1  |  $ docker run -it ubuntu bash
hello_1  |
hello_1  | Share images, automate workflows, and more with a free Docker ID:
hello_1  |  https://hub.docker.com/
hello_1  |
hello_1  | For more examples and ideas, visit:
hello_1  |  https://docs.docker.com/get-started/
hello_1  |
hello_hello_1 exited with code 0
```


# 参考
[WSLでdockerのインストールからdocker-composeまで動かす](https://qiita.com/KEINOS/items/43d9415e351d80f78c8b)
[最もシンプルな docker-compose.yml の hello-world](https://qiita.com/tettsu__/items/85c96850d187e4386c24)
