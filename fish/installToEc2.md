EC2(Amazon Linux2)へfishをインストール

# EPELレポジトリの追加
Extra Packages for Enterprise Linux (EPEL) リポジトリを追加します。

```
$ sudo amazon-linux-extras install -y epel
```

# Fishのインストール
Fishをインストールします。

```
$ sudo yum -y install fish
```

Fishシェルに入ります。

```
$ fish
Welcome to fish, the friendly interactive shell
Type help for instructions on how to use fish
>
```

# Fisherのインストール
Fishシェルに入った状態で、Fishのパッケージマネージャ`Fisher`をインストールします。

```
> curl https://git.io/fisher --create-dirs -sLo ~/.config/fish/functions/fisher.fish
```

👀確認

```
> fisher -v
```

🚀試しにテーマをインストール

```
> fisher add oh-my-fish/theme-agnoster
```

# その他
fishテーマ`theme-agnoster`等をインストールした場合、クライアント側に`powerlineフォント`がないと、一部文字化けします。
[ここ](https://github.com/powerline/fonts)等からダウンロードし、クライアントへフォントの追加が必要です。

📖[WSL環境 (Windows10) でPowerlineフォントをインストールする](https://qiita.com/mosh_shu/items/0dc06e8b4aecf7d68316)

以上。