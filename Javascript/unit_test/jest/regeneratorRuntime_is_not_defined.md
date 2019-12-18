asyncを使用したコードをjestでテストすると「regeneratorRuntime is not defined」エラーが発生



# エラー内容

jestでテスト実行すると、asyncの記述の部分で、以下のようなエラーが発生しました。

```
    ReferenceError: regeneratorRuntime is not defined

       9 | }
      10 | 
    > 11 | const main = async () => {
         |              ^
```

# 対応

.babelrcにて、babelのターゲットを指定することで解決しました。
変更前：

```Javascript:.babelrc
{
  "presets": [
    "@babel/preset-env"
  ]
}
```

変更後：
```Javascript:.babelrc
{
  "presets": [
    [
      "@babel/preset-env", {
        "targets": {
          "node": "current"
        }
      }
    ]
  ]
}
```

