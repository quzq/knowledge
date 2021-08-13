
# bullで管理するNodeJSの分散job

# bull　とは
NodeJSで分散ジョブとメッセージを処理するためのキューパッケージです。redisをベースに動作します。
kueの後継的なライブラリです。

# 確認環境
- node: v10.17.0
- [redis ( for windows )](https://github.com/microsoftarchive/redis/releases): 3.0.504

# express-generator にてプロジェクト作成
ここでは、`play_node_bull`というプロジェクトで作成します。

```
npx express-generator play_node_bull
cd play_code_bull
npm install
```

# bull　のインストール

```
npm install bull
```

# app.js の編集
app.jsに以下の行を追記します。
```Diff:App.js
var createError = require('http-errors');
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');

var indexRouter = require('./routes/index');
var usersRouter = require('./routes/users');
+ var jobsRouter = require('./routes/jobs');  

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/users', usersRouter);
+ app.use('/jobs', jobsRouter); 

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

module.exports = app;
```

# jos.js　の追加
### 1ブロック目
routesフォルダに、jobs.jsを追加します。
routerとBullを準備します。
```Javascript:routes/jobs.js
const express = require('express');
const router = express.Router();
const Bull = require("bull")
```
### 2ブロック目
重い処理のスタブを準備します。
```Javascript:routes/jobs.js
// 重い処理のスタブ
const heavyJob = async () => {
  return new Promise((resolve)=>{
    setInterval(() => {
      resolve();
     }, 10000);
  })  
}
```

### 3ブロック目
bullオブジェクトと処理名を定義します。
処理名は、キューの送信と受信を紐づけるキーとなります。
```Javascript:routes/jobs.js
const queue = new Bull('bulltest', { redis: { port: 6379, host: '127.0.0.1' } });

const processor = 'processor_name'; // 処理名
```

### 4ブロック目
エンキューするエンドポイントを定義します。
エンキューするqueue.add()には、データオブジェクトを処理名を指定します。
```Javascript:routes/jobs.js
router.post('/', (req, res, next) => {
  const data = { greetingTime: (new Date()).toString() }     // jobのデータ(object)
  queue.add(processor, data);
  res.send('ok');
});

module.exports = router;
```

### 5ブロック目
デキューするハンドラを記述します。
queue.process()に処理名と並行処理数、ハンドラ関数を指定します。
```Javascript:routes/jobs.js
const concurrency = 2;  // 並行処理数

queue.process(processor, concurrency, async (job, done) => {
  try {
    await heavyJob();
    done();
  } catch (e) {
    done(e);
  }
});
```

# 実行
アプリケーションを実行します。デフォルトでは3000番ポートで待機します。
```
npm start
```

# APIの呼び出し
curlコマンドで、作成した`jobs`APIを呼び出します。

```
curl -X POST 'http://127.0.0.1:3000/jobs'
```


# jobのオプション

```Javascript
interface JobOpts {
  priority: number, // オプションの優先度。1（最高の優先度）～MAX_INT（最低の優先度）を指定する。パフォーマンスにわずかな影響を与えるため、必要でない限り使用しない。
  delay: number, // このジョブを処理できるようになるまで待機するミリ秒数。正確に遅延するためには、サーバーとクライアントの時刻を同期する必要がある。
  attempts: number, //　ジョブが完了するまでのリトライ回数
  repeat: RepeatOpts, // cron仕様に従ってジョブを繰り返す。
  backoff: number | BackoffOpts, // ジョブが失敗した場合の自動再試行設定。遅延時間を設定するか、{type: 'fixed' or 'exponential')を指定。
  lifo: boolean, // 後入れ先出しにする (default false)
  timeout: number, // タイムアウトエラーでジョブが失敗するまでのミリ秒数
  jobId: number | string, // ジョブIDを上書きする。既に存在するIDを持つジョブを追加しようとしても、追加されない。
  removeOnComplete: boolean | number, // trueの場合、正常に完了したときにジョブを削除します。falseの場合は`completed`セットに保持される。
  removeOnFail: boolean | number, // trueの場合、処理に失敗したときにジョブを削除します。falseの場合は`faild`セットに保持される。
  stackTraceLimit: number, // スタックトレースに記録されるスタックトレース行の量を制限。
}
```