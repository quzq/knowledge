# 分割代入（Destructuring Assignment）を極める

## TL;DR

JavaScript の分割代入は、配列やオブジェクトから必要な値だけを取り出すための構文です。スプレッドオペレータ（`...`）を使うと、配列やオブジェクトから「残りの要素」や「残りのプロパティ」をまとめて取得できます。
また、存在しないプロパティに対してデフォルト値を割り当てることもできます。

## 1. 配列の例

### 基本的な使い方

```js
const colors = ["red", "green", "blue", "yellow", "purple"];

const [primary, secondary, ...restColors] = colors;

console.log(primary); // "red"
console.log(secondary); // "green"
console.log(restColors); // ["blue", "yellow", "purple"]
```

- `primary` には最初の要素 (`"red"`)
- `secondary` には 2 番目の要素 (`"green"`)
- 3 番目以降はすべて `restColors` に配列として格納されます

### 要素数が足りない場合・デフォルト値の利用

配列の要素数が、分割代入で指定した変数の数より少ない場合には、**対応する変数が `undefined`** となります。  
ただし、**デフォルト値を指定**しておくと、`undefined` が返ってきた場合にデフォルト値が適用されます。

```js
const array = [10, 20];

const [a, b, c = 30] = array;

console.log(a); // 10
console.log(b); // 20
console.log(c); // 30  (配列に3番目がないため、デフォルト値が使用される)
```

### スプレッドオペレータとの組み合わせ

スプレッドオペレータを使うと、「残りの要素」全体を一括で取得できます。もし配列の要素数が少ない場合は、スプレッド先に空の配列が代入されます。

```js
const array2 = [10, 20];

// 3番目以降を restNumbers で受け取る
const [x, y, ...restNumbers] = array2;

console.log(x); // 10
console.log(y); // 20
console.log(restNumbers); // [] (3番目以降が存在しないため空配列)
```

## 2. オブジェクトの例（スプレッドオペレータを含む）

次にオブジェクトです。オブジェクトの場合は、プロパティ名を指定して抽出します。スプレッドオペレータを使うことで「まだ分割していない、残りのプロパティ」をまとめて取得できます。さらに、存在しないプロパティに対してデフォルト値を指定することもできます。

```js
const user = {
  name: "Alice",
  age: 25,
  // city は定義されていない
};

const { name, city = "Unknown", ...restProps } = user;

console.log(name); // "Alice"
console.log(city); // "Unknown"
console.log(restProps); // { age: 25 }
```

- `name` には `"Alice"`
- `city` が存在しないため、デフォルト値として `"Unknown"` が適用
- `restProps` には `age: 25` がオブジェクトとして格納されます

もう少し複雑な例として、ネストしたプロパティの一部だけを展開し、残りをまとめることもできます。

```js
const config = {
  host: "localhost",
  port: 3000,
  credentials: {
    username: "admin",
    role: "editor",
  },
};

const {
  host,
  port,
  credentials: { username, ...restCredentials },
  ...restConfig
} = config;

console.log(host); // "localhost"
console.log(port); // 3000
console.log(username); // "admin"
console.log(restCredentials); // { role: "editor" }
console.log(restConfig); // {} （credentials 以外の残りプロパティは何もないので空オブジェクト）
```

## 3. 配列とオブジェクトを混ぜた例

実際のアプリケーションでは、オブジェクトのプロパティが配列であったり、配列の要素がオブジェクトであったりと複雑な構造が一般的です。以下は、そのようなケースでの分割代入の例です。

```js
const data = {
  numbers: [10, 20, 30, 40],
  user: {
    name: "Bob",
    info: {
      age: 28,
      // city プロパティは存在しない
    },
  },
};

const {
  numbers: [firstNum, secondNum, ...restNumbers],
  user: {
    name,
    info: { age, city = "Unknown" },
  },
} = data;

console.log(firstNum); // 10
console.log(secondNum); // 20
console.log(restNumbers); // [30, 40]
console.log(name); // "Bob"
console.log(age); // 28
console.log(city); // "Unknown"（存在しないのでデフォルト値が適用）
```

- 配列要素を指定した順番で変数に代入し、スプレッドオペレータで残りをまとめています。
- オブジェクトでもデフォルト値を活用し、存在しない `city` プロパティに `"Unknown"` を自動的に設定しています。
