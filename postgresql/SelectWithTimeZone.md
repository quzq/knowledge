

# まとめ
日本時間として渡した日付は、utcタイムゾーンのホストで処理する前に、jstタイムゾーンの日付として前処理する必要がある。
例）日本で動作するUIから、当日日付をEC2（utcタイムゾーン）上のAPIに送った場合

```
function API(当日日付){
  const today = moment.tz(当日日付, 'Asia/Tokyo')   // jstとして処理
  ( …メイン処理 )
}
```


モデル定義
```Javascript
const Test = sequelize.define("test",{
  created_at: {
    type: Sequelize.DATE,
    allowNull: false
  },
},{
  timestamps: false,
  freezeTableName: true
});
```

モデルをDBと同期させて物理テーブルを作成し、スキーマーを確認。

```
tanawari=> \d test
                                  テーブル "public.test"
     列     |            型            |                      修飾語
------------+--------------------------+---------------------------------------------------
 id         | integer                  | not null default nextval('test_id_seq'::regclass)
 created_at | timestamp with time zone | not null
インデックス:
    "test_pkey" PRIMARY KEY, btree (id)
```
created_atカラムは、`timestamp with time zone`型として定義されています。

utcタイムゾーンのホスト上で以下を実行しました。日本時間は、2020年2月22日午前8時25分です。

```Javascript
const created_at = new Date()
console.log(created_at)
await Test.create({
  created_at: created_at,
})
```

console.log()からは、「2020-02-21T23:25:51.558Z」が出力されています。utcなので、日本時間の9時間前です。


testテーブルの中身を確認します。
```
=> select * from test;
 id |         created_at
----+----------------------------
  1 | 2020-02-21 23:25:51.558+00
```

2月22日に入力したデータなので、これを単純に以下のように抽出してみました。

```Javascript
const dateFrom = moment('2020-02-21 00:00:00')
const dateTo = moment('2020-02-22 00:00:00')
console.log(dateFrom, dateTo)
const count = await Test.count({
  where: {
    created_at: { [Op.gte]: dateFrom.toDate(), [Op.lt]: dateTo.toDate() },
  }
})
```

結果、ヒットしません。
console.log()からは **moment("2020-02-21T00:00:00.000") moment("2020-02-22T00:00:00.000")** が出力されました。
つまり、utcの2月22日として検索されてしまいました。

なので、これを以下のようにタイムゾーンを加えて修正します。

```Javascript
const dateFrom = moment.tz('2020-02-21 00:00:00','Asia/Tokyo')      // updated!
const dateTo = moment.tz('2020-02-22 00:00:00','Asia/Tokyo')        // updated!
console.log(dateFrom, dateTo)
const count = await Test.count({
  where: {
    created_at: { [Op.gte]: dateFrom.toDate(), [Op.lt]: dateTo.toDate() },
  }
})
```
ヒットしました。
console.log()からは、**moment.parseZone("2020-02-21T09:00:00.000+09:00") moment.parseZone("2020-02-22T00:00:00.000+09:00")** が出力されました。
