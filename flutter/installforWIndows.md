Flutter2 開発環境をWindowsにインストール

### 確認環境
Windows 10 Pro 

### Flutter2 SDK のダウンロード

flutterの[公式サイト](https://flutter.dev/)から、Flutter SDK をダウンロードします。
トップページの`Get started`ボタンを押します。

インストール対象となるOSを選択します（ここではwindows）。

少し下にある、`flutter_windows_2.*.*-stable.zip`(*はリリース状況により読み替えてください)ボタンを押すとダウンロードが始まります。


### Flutter2 SDK のインストール

## 配置
ダウンロードした`flutter_windows_2.*.*-stable.zip`ファイルを解凍し、任意の位置に展開します。
zip内のflatterフォルダごと展開します。


##  環境変数の設定
環境変数を設定し、pathを通します。
指定するのは、SDKの展開先のflatterフォルダ直下のbinです。
ここでは、`c:\Users\quzq`にflatterフォルダを配置してるので、`c:\Users\quzq\flatter\bin`を設定します。

## 設定したpathの確認
PowerShellを起動し、以下のコマンドが正しく動作するのを確認。

```
PS > where.exe flutter dart
C:\Users\quzq\flutter\bin\flutter
C:\Users\quzq\flutter\bin\flutter.bat
C:\Users\quzq\flutter\bin\dart
C:\Users\quzq\flutter\bin\dart.bat
```

## flutter doctorの実行
flutter doctorを実行し、インストールを完了するために必要なプラットフォームの依存関係があるかどうかを確認します。
※flutter doctorは、flutterフォルダの直下で実行します。

```
PS > flutter doctor
Running "flutter pub get" in flutter_tools...                      12.0s
Doctor summary (to see all details, run flutter doctor -v):
[√] Flutter (Channel stable, 2.2.3, on Microsoft Windows [Version 10.0.19043.1110], locale ja-JP)
[!] Android toolchain - develop for Android devices (Android SDK version 29.0.3)
    X Android license status unknown.
      Run `flutter doctor --android-licenses` to accept the SDK licenses.
      See https://flutter.dev/docs/get-started/install/windows#android-setup for more details.
[√] Chrome - develop for the web
[√] Android Studio (version 3.6)
[√] VS Code (version 1.57.0)
[√] Connected device (2 available)

! Doctor found issues in 1 category.
PS C:\Users\USER\flutter>
```


### android studioのインストール
[公式サイト](https://developer.android.com/studio)より、android studioをダウンロードします。
ダウンロードしたexeファイルを実行し、インストールします。


### ライセンスの取得
flutter doctorのコマンドを使って、ライセンスを取得します。
すべて`y`を押していく感じです。

```
PS C:\Users\USER\flutter> flutter doctor --android-licenses
Warning: File C:\Users\USER\.android\repositories.cfg could not be loaded.
6 of 7 SDK package licenses not accepted. 100% Computing updates...
Review licenses that have not been accepted (y/N)? y

1/6: License android-googletv-license:
---------------------------------------
Terms and Conditions

This is the Google TV Add-on for the Android Software Development Kit License Agreement.

...(以下、略)

```
再度、flutter doctorコマンドを確認。
