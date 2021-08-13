javascriptで回帰直線を求める


# 回帰直線とは
組となるデータの、中心的な分布傾向を表す直線。
最小二乗法( least squares method )によって求められる。

# 最小二乗法の式

📖[最小二乗法(wikipedia)](https://ja.wikipedia.org/wiki/%E6%9C%80%E5%B0%8F%E4%BA%8C%E4%B9%97%E6%B3%95)

これを、javascript( ES6 )で関数化します。
`Σ`を`reduce`で表現しています。

```Javascript
// 回帰直線を求める（最小二乗法）
const lsm = coordinates => {
  const n = coordinates.length
  const sigX = coordinates.reduce((acc, c) => acc += c.x, 0)
  const sigY = coordinates.reduce((acc, c) => acc += c.y, 0)
  const sigXX = coordinates.reduce((acc, c) => acc += c.x * c.x, 0)
  const sigXY = coordinates.reduce((acc, c) => acc += c.x * c.y, 0)
  // a(傾き)を求める
  const a = (n * sigXY - sigX * sigY) / (n * sigXX - Math.pow(sigX, 2));
  // b(切片)を求める
  const b = (sigXX * sigY - sigXY * sigX) / (n * sigXX - Math.pow(sigX, 2));
  return { a, b }
}
```

# 👀動作確認
ドライバを作成して呼び出してみます。

```Javascript
  const coordinates = [
    { x: 200, y: 100 },
    { x: 250, y: 150 },
    { x: 300, y: 150 },
    { x: 312, y: 200 },
    { x: 330, y: 210 },
    { x: 100, y: 80 },
    { x: 110, y: 40 },
  ]

  const { a, b } = lsm(coordinates)
  console.log(a, b)
```

結果

```
0.6227602344112315 -9.665985075256124
```

### 📈グラフを書いて確認
canvasにグラフを描画して確認してみました。
青のドットがサンプル、赤が回帰直線です。


ソースコードは[ここ](https://github.com/quzq/function_lsm)。

以上。