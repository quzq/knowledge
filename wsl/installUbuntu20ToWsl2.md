

# WSLのインストール
WSLがまだインストールされていない場合、インストールが必要です。

管理者として PowerShell を開き、以下を実行。

```powershell
PS > dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
```

# WSL2 に更新

##  「仮想マシン プラットフォーム」のオプション コンポーネントの有効化
WSL2 に必要な、「仮想マシン プラットフォーム」オプション機能を有効にする必要があります。

管理者として PowerShell を開き、以下を実行。

```powershell
PS > dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```

## WSL2 を既定のバージョンとして設定
新しい Linux ディストリビューションをインストールするときに、WSL2 を既定のバージョンとして設定します。

管理者として PowerShell を開き、以下を実行。

```powershell
PS > wsl --set-default-version 2
```

# Ubuntu 20.04 LTS のインストール
WSL2にUbuntu20.04LTSをインストールします。
MIcrosoftStoreから「ubuntu」で検索し、「入手」（または「インストール」）ボタンを押してインストール。

# WSLのバージョン確認
PowerShell を開き、以下を実行。

```powershell
PS > wsl --list -v
  NAME              STATE           VERSION
* Ubuntu-20.04      Running         2
```
Ubuntu-20.04が、VERSION 2 （WSL2）で動作しているのを確認できます。


以上。

