# 2019-11-27 Wed
### nginx
  - conf
    - http
      - server
        - listen:
          - ipとポート番号を指定する。
          - 80  全てのIPv4アドレスの80番
          - [::]:80  全てのIPv6アドレスの80番
        - server_name: サーバー名を指定する
        - root: ドキュメントルートを指定する
        - location:
          - エイリアスの設定。パス毎のドキュメントルートを指定する
        - client_max_body_size: 最大アップロードサイス。413エラー発生時に確認する。
