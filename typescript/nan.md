NaNはNumber型

# tl;dr
NaN(Number.NaN)は数値ではありませんがNumber型です。

# 説明
数値、またはNaNの変数は以下のように定義します。

```
const foo: number = 123;
const foo: number = Number.NaN;
```

以下のような定義はできません。
```
const foo: number | NaN = Number.NaN;
```

NaNのチェックには`isNaN()`を使います。
```
const foo: number = parseInt('a')
const bar: number = Number('a')
console.log('foo', foo, isNaN(foo))
console.log('bar', bar, isNaN(bar))
```
実行結果
```
foo NaN true
bar NaN true
```

単純に真偽値でチェックすると0との区別がつかないので注意！
```
const foo: number = parseInt('a')
const bar: number = parseInt('0')
console.log('foo', foo, !foo)
console.log('bar', bar, !bar)

```
```
foo_ NaN true
bar_ 0 true
```
