WLS2にFlutter開発環境をインストール

# snapをWSL2上で動かす 
最初に以下のコマンドを実行し、snapを実行できるようにします。
ちなみに、Ubuntu 16.04 LTS以降のUbuntuにはデフォルトでsnapdがインストールされています。

:book: [WSL2- Ubuntu 20.04 Snap store doesn't work due to systemd dependency](https://github.com/microsoft/WSL/issues/5126)
```
sudo apt-get update && sudo apt-get install -yqq daemonize dbus-user-session fontconfig
sudo daemonize /usr/bin/unshare --fork --pid --mount-proc /lib/systemd/systemd --system-unit=basic.target
exec sudo nsenter -t $(pidof systemd) -a su - $LOGNAME
```
snapの動作を確認します。
```
snap version
```

# snapを使ってFlutterをインストール
[公式](https://flutter.dev/docs/get-started/install/linux)に従い、以下のコマンドでflutterをインストールします。
```
sudo snap install flutter --classic
```

flutter doctorコマンドでインストールする必要がある依存関係をチェック
:memo:このコマンド実行時、再度インストールが始まります。
```
flutter doctor
```
結果例 )
```
Running "flutter pub get" in flutter_tools...                      10.6s
Doctor summary (to see all details, run flutter doctor -v):
[✓] Flutter (Channel stable, 2.2.1, on Linux, locale C.UTF-8)
[✗] Android toolchain - develop for Android devices
    ✗ Unable to locate Android SDK.
      Install Android Studio from: https://developer.android.com/studio/index.html
      On first launch it will assist you in installing the Android SDK components.
      (or visit https://flutter.dev/docs/get-started/install/linux#android-setup for detailed instructions).
      If the Android SDK has been installed to a custom location, please use
      `flutter config --android-sdk` to update to that location.

[✗] Chrome - develop for the web (Cannot find Chrome executable at google-chrome)
    ! Cannot find Chrome. Try setting CHROME_EXECUTABLE to a Chrome executable.
[!] Android Studio (not installed)
[!] Connected device
    ! No devices available

! Doctor found issues in 4 categories.
```

