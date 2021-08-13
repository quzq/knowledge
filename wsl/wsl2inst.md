WSL1を残したまま、WSL2のUbuntuでdockerを動かすまで

# tl;dr
WSL1のUbuntu18.04を消したくないので、ストアを使わずにWSL2へUbuntu18.04を手動インストールし、そこにdockerをインストールします。

# Windows Insider Program への参加
[公式](https://insider.windows.com/ja-jp/getting-started/)の通りに参加します。
- サイトにてWindows Insider Programへの参加する
- 参加したアカウントとWindows10への紐付け

# WSL2の準備
`windows update`より、新しいwindowsのビルドを取得します。

新しいビルドがインストールされたら、管理者として起動したpowershell上で以下のコマンドを実行します。
ちなみにWSL1を既に実行している場合は、２行目は不要です。

```	
Enable-WindowsOptionalFeature -Online -FeatureName VirtualMachinePlatform
Enable-WindowsOptionalFeature -Online -FeatureName Microsoft-Windows-Subsystem-Linux
```

# Ubuntuパッケージのダウンロード
microsoftストアのUbuntu18.04は、既にWSL1にセットアップされているため使えません。
そのため、Ubuntu18.04パッケージを[ここ](https://lxrunoffline.apphb.com/download/Ubintu/bionic)からダウンロード。
`ubuntu-bionic-core-cloudimg-amd64-root.tar.gz`というファイルがダウンロードされます。

# LxRunOfflineのダウンロード
WSLを管理するためのフル機能のユーティリティである`LxRunOffline`を、[ここ](https://github.com/DDoSolitary/LxRunOffline/releases)からダウンロード。`LxRunOffline-v3.4.1-msvc.zip`です

# Ubuntuをインストール
ダウンロードした`LxRunOffline`を適当な場所に解凍します。
ついでにダウンロードしたUbuntuパッケージも同じ場所に配置します。
powsershellを管理者として起動し、カレントを`LxRunOffline`に移動後、以下のWSLインストールコマンドを実行します。

```
LxRunOffline i -n {任意のディストリビューション名} -d {インストールしたい任意のディレクトリ} {パッケージファイル}
```

私の場合は、「D:\WSL\Ubuntu-18-wsl2」にインストールし、ディストリビューション名（今後、WSLの操作に使用する）を「Ubuntu-18-wsl2」としました。

```
> .\LxRunOffline i -n Ubuntu-18-wsl2 -d D:\WSL\Ubuntu-18-wsl2 -f ubuntu-bionic-core-cloudimg-amd64-root.tar.gz
```

# Ubuntuの準備
引き続きの`LxRunOffline`コマンドを使って、Ubuntuを起動します。まだ、この時点ではWSL2になっていません。
ディストリビューション名は「Ubuntu-18-wsl2」になっていますが、そこは、ご自分のつけたディストリビューション名に読み替えてください。
```
> .\LxRunOffline run -n {作成したディストリビューション名}
```

私の場合は、以下のコマンドを実行しました。

```
> .\LxRunOffline run -n Ubuntu-18-wsl2
```

起動したUbuntuに対し、以下のコマンドを実行します。
途中で「myuser」というユーザーを作成していますが、そこは任意に変えてください。
```
# apt update
# apt upgrade
# adduser myuser
# apt install sudo
# gpasswd -a myuser sudo
# login myuser
```

最後のコマンドで、作成したユーザーでログインしたので、`id`コマンドを使い、作成したユーザーのuidを確認します。

```
$ id
uid=1000(myuser) gid=1000(myuser) groups=1000(myuser),27(sudo)
```
uid（私の場合は1000）を確認したら、一旦Ubuntuからログアウトします。

# WSLのデフォルトユーザーの設定
管理者として起動したpowershell上で以下のコマンドを実行します。
```
> .\LxRunOffline su -n {作成したディストリビューション名} -v {上記で確認したuid}
```

私の場合は、以下のコマンドを実行しました。

```
> .\LxRunOffline su -n Ubuntu-18-wsl2 -v 1000

```

# WSL1 から WSL2 へのアップグレード
管理者として起動したpowershell上で以下のコマンドを実行します。

```
> wsl --set-version {作成したディストリビューション名} 2
```

私の場合は、以下のコマンドを実行しました。

```
> wsl --set-version Ubuntu-18-wsl2 2
```

以下のコマンドで正しくアップグレードされたか確認します。

```
> wsl -l -v
  NAME              STATE           VERSION
* Ubuntu-18.04      Running         1
  Ubuntu-18-wsl2    Running         2
```
Ubuntu-18.04がWSL1で実行中であり、Ubuntu-18-wsl2がWSL2で実行中であるのがわかります。

# dockerのインストールと確認
powershell上で以下のコマンドを実行し、WSL2のUbuntuにログインします。
```
> .\LxRunOffline run -n {作成したディストリビューション名}
```
Ubuntuにログイン後、[公式](https://docs.docker.com/install/linux/docker-ce/ubuntu/)通りの手順でdockerをインストール。

```
$ sudo apt-get update
$ sudo apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
$ sudo add-apt-repository    "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
$ sudo apt-get update
$ sudo apt-get install docker-ce docker-ce-cli containerd.io
```

### 👀 動作確認
dockerサービスの起動

```
$ sudo service docker start
```

hello-world実行

```
$ sudo docker run hello-world

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
 ```

 以上。