PostgresSQLのテーブルにjsonデータをインポート

# 確認環境
📀PostgreSQL 9.6.14

# 対象テーブルの定義
インポート対象のテーブルとして、`my_table`を定義します。

```SQL
create table my_table (
  id integer default nextval('steps_id_seq'::regclass) not null
  , name character varying(255) not null
  , age integer not null
  , primary key (id)
);
```
# インポートデータの定義
以下のファイルを作成し、`/home`下に配置しました。フルパスは`/home/sample.json`になります。

```Javascript:sample.json
[{ "name": "Mike", "age": 18 }, { "name": "Jhon", "age": 45 }, { "name": "Bob", "age": 23 }]
```
※改行を入れると、取り込み時にエラーが発生します。改行は入れないでください。

# データのインポート
psql等で以下を実行します。

```SQL
BEGIN;

create temporary table temp_json (values text) on commit drop;
copy temp_json from '/home/sample.json';

insert into my_table ("name", "age") 
select 
  values->>'name' as name,
  cast( values->>'age' as integer ) as age
from (
  select json_array_elements(replace(values,'\','\\')::json) as values 
  from temp_json
) a; 

COMMIT;
```

# 結果
👀テーブルの内容を確認します。
```
select * from my_table;

 id  | name | age
-----+------+-----
 101 | Mike |  18
 102 | Jhon |  45
 103 | Bob  |  23
(3 行)
```

以上。