
# Node.js実行中に「JavaScript heap out of memory」が発生する

# tl:dr
 NODE_OPTIONSの `max-old-space-size`を変更します。

# 🔥現象例
```
 <--- Last few GCs --->

 [326:0x41400b0]   917511 ms: Mark-sweep 1511.0 (1774.9) -> 1510.0 (1775.4) MB, 2311.1 / 0.0 ms  allocation failure GC in old space requested
 [326:0x41400b0]   920160 ms: Mark-sweep 1510.0 (1775.4) -> 1509.8 (1728.9) MB, 2643.0 / 0.0 ms  last resort GC in old space requested
 [326:0x41400b0]   922034 ms: Mark-sweep 1509.8 (1728.9) -> 1509.8 (1719.9) MB, 1873.7 / 0.0 ms  last resort GC in old space requested


 <--- JS stacktrace --->

 ==== JS stack trace =========================================

 Security context: 0x1fe852a5891 <JSObject>
     1: push(this=0x1f442c2ea851 <JSArray[43037]>)
     2: query [/webapp/node_modules/pg/lib/client.js:512] [bytecode=0xe166cd56f31 offset=384](this=0x22a2a88e70d9 <Client map = 0x320165dfab1>,config=0x824a0505cb9 <String[186]: INSERT INTO "pancakes" ("value","createdAt","updatedAt","cake_id") VALUES ('','2021-07-11 01:07:39.748 +00:00','2021-07-11 01:07:39.748 +00:00',31613,60...

 FATAL ERROR: CALL_AND_RETRY_LAST Allocation failed - JavaScript heap out of memory
  1: node::Abort() [/usr/local/bin/node]
  2: 0x8cd14c [/usr/local/bin/node]
  3: v8::Utils::ReportOOMFailure(char const*, bool) [/usr/local/bin/node]
  4: v8::internal::V8::FatalProcessOutOfMemory(char const*, bool) [/usr/local/bin/node]
  5: v8::internal::Factory::NewUninitializedFixedArray(int) [/usr/local/bin/node]
  6: 0xd8366f [/usr/local/bin/node]
  7: 0xd977a5 [/usr/local/bin/node]
  8: v8::internal::JSObject::AddDataElement(v8::internal::Handle<v8::internal::JSObject>, unsigned int, v8::internal::Handle<v8::internal::Object>, v8::internal::PropertyAttributes, v8::internal::Object::ShouldThrow) [/usr/local/bin/node]
  9: v8::internal::Object::AddDataProperty(v8::internal::LookupIterator*, v8::internal::Handle<v8::internal::Object>, v8::internal::PropertyAttributes, v8::internal::Object::ShouldThrow, v8::internal::Object::StoreFromKeyed) [/usr/local/bin/node]
 10: v8::internal::Object::SetProperty(v8::internal::LookupIterator*, v8::internal::Handle<v8::internal::Object>, v8::internal::LanguageMode, v8::internal::Object::StoreFromKeyed) [/usr/local/bin/node]
 11: v8::internal::Runtime_SetProperty(int, v8::internal::Object**, v8::internal::Isolate*) [/usr/local/bin/node]
 12: 0x13b934d042fd
 Aborted (core dumped)
```

# 🚀対応
Node.js実行時に、`max-old-space-size'(ヒープ領域の最大値設定)を増やします。
こちらによると（https://qiita.com/kawanet/items/cfedd535990b32710c50）デフォルトは1400MBらしいので、それよりも大きい値を指定します。

```
node --max-old-space-size=3000 index.js
```

# 📖参考
ヒープ領域：コンピュータープログラミングにおいて、動的に確保可能なメモリの領域。
(https://ja.wikipedia.org/wiki/%E3%83%92%E3%83%BC%E3%83%97%E9%A0%98%E5%9F%9F)



	--max-old-space-size : ヒープメモリの領域の最大値を指定する
