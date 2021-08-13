シェルからPostgreSQLのpg_hba.confファイルの場所を探す

# tl;dr
シェルからPostgreSQLのpg_hba.confファイルの場所を取得するには、psqlのコマンド、`show hba_file`を使います。

# コマンド

```
psql -U {ユーザー名} -t -P format=unaligned -c 'show hba_file';
```

# 実行例

※ユーザー名を`postgres`としています。

```
$ psql -U postgres -t -P format=unaligned -c 'show hba_file';
/var/lib/postgresql/data/pg_hba.conf
```
※返される結果は、ホストOSの種類やPostgreSQLのバージョンにより異なります。


以上