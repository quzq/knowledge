# linuxパッケージのダウンロード
https://lxrunoffline.apphb.com/download/Ubintu/bionic

# LxRunOfflineのダウンロード
https://github.com/DDoSolitary/LxRunOffline/releases
LxRunOffline-v3.4.1-msvc.zip

# パッケージをインストール
powsershell 管理者として
```
> .\LxRunOffline i -n Ubuntu-18-wsl2 -d D:\WSL\Ubuntu-18-wsl2 -f ubuntu-bionic-core-cloudimg-amd64-root.tar.gz
> .\LxRunOffline run -n Ubuntu-18-wsl2

```

```
# apt update
# apt upgrade
# adduser user
# apt install sudo
# gpasswd -a user sudo      # add user in sudo group
# login user
```

```
$ id
uid=1000(user) gid=1000(user) groups=1000(user),27(sudo)
```


```
 .\LxRunOffline su -n Ubuntu-18-wsl2 -v 1000

```


 wsl --set-version Ubuntu-18-wsl2 2

apt install sudo
apt install vim


    5  sudo
    6  sudo apt-get update
    7  sudo
    8  sudo apt-get update
    9  sudo apt-get install     apt-transport-https     ca-certificates     curl     gnupg-agent     software-properties-common
   11  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
   12  sudo apt-key fingerprint 0EBFCD88
   13  sudo add-apt-repository    "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
   14  sudo apt-get update
   15  sudo apt-get install docker-ce docker-ce-cli containerd.io
   16  docker -v
   17  sudo docker run hello-world
   21  sudo service docker start