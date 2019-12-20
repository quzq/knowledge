WSLディストリビューションの保管場所を変更する

# 事前準備
Windows Power Shellを管理者として実行し、WSLサービスの再起動をしておきます。
これを怠ると、ディストリビューションの移動に失敗することがあります。

```
PS C:\> net stop LxssManager
PS C:\> net start LxssManager
```

# 移動先フォルダの作成
今回は、Dドライブの`WSL`フォルダに移動することにしました。
フォルダを作成し、権限を付与します。

```
PS C:\> cd d:
PS D:\> mkdir WSL
PS D:\> icacls D:\WSL /grant $env:USERNAME":(OI)(CI)(F)"
```

# LxRunOfflineのダウンロード
[ここ](https://github.com/DDoSolitary/LxRunOffline/releases)からLxRunOffline.zipをダウンロード。
ここでは、`LxRunOffline-v3.4.0.zip`をダウンロードしました。
zipファイルを適当なフォルダに展開します。今回は`C:\LxRunOffline`フォルダに展開しました。

# WSL ディストリビューション名の取得
移動操作のために、ディストリビューション名を取得しておきます。

```
PS D:\> cd c:\LxRunOffline
PS C:\LxRunOffline> .\LxRunOffline.exe list
Ubuntu-18.04
```
`Ubuntu-18.04`がディストリビューション名です。

# ディストリビューションの移動
moveコマンドを使用し、WSLフォルダを移動します。
- `-n` ディストリビューション名
- `-d` ディストリビューションの移動先

```
PS C:/LxRunOffline> .\LxRunOffline.exe move -n Ubuntu-18.04 -d d:\wsl\Ubuntu-18.04
```

以下のコマンドで、移動先を確認できます。

```
PS C:\LxRunOffline> .\LxRunOffline.exe get-dir -n Ubuntu-18.04
d:\wsl\Ubuntu-18.04
```

# 注意事項
ディストリビューションの移動を行うと、ディストリビューションのアンインストールしても、移動先フォルダは残ったままになるようです。

以上。

