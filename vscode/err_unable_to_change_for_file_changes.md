# Visual Studio Code is unable to watch for file changes in this large workspace" (error ENOSPC)
`この大規模なワークスペースでのファイルの変更をウォッチできません。`

`max_user_watches`の設定を変更する必要がある。
現在の設定値の確認：
```
cat /proc/sys/fs/inotify/max_user_watches
```
値の変更：
`/etc/sysctl.conf`を変更し、末尾に以下を追加。
```
fs.inotify.max_user_watches=524288
```
変更値のリロード：
```
sudo sysctl -p
```

