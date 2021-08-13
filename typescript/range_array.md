typescriptで連番の配列を１行で作成する

# tl;dr
`[...Array(n)].map((_:undefind, idx:number) => idx))`で連番の配列を作成できます。

# 説明
`[...Array(number)]`を実行すると、number個の`undefined`の一次元配列を返します。

```Javascript
const foo = [...Array(5)]
console.log(foo)
```
実行結果：
```
[ undefined, undefined, undefined, undefined, undefined ]
```

そこで、map関数のindexを使って値を作成します。
```Typescript
const foo = [...Array(5)].map((_:undefined, idx:number) => idx))
console.log(foo)
```
実行結果
```
[ 0, 1, 2, 3, 4 ]
```

以上。

