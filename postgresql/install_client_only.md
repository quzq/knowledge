# Ubuntuにpsqlのみをインストール

# tl;dr
Ubuntu 16.04以降に、postgresqlクライアント(psql)のみをインストールします。

# インストール
以下を実行します。

```
$ sudo apt install postgresql-client  
```

# 👀確認

```
$ psql --version
psql (PostgreSQL) 12.2 (Ubuntu 12.2-4)
```

# 💡参考
dbホストへの接続例：

```
$ psql -h localhost -p 5432 -U postgres
```



