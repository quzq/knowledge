
# 関数
##  関数の戻り値
-  戻り値はカンマ区切りで複数返せる。
  - return 'abc',100
  - 戻り値の型はタプルとなる。('abc', 100)

# データ型
  - dict型 {"abc": 1}  mutable
  - list型 [1, 2, 3, 4, 5]  mutable
  - tapple型 (1, 2, 3, 4, 5)  immutable
  - None =null

# 演算子
## 三項演算子
```
x = "OK" if n == 10 else "NG"
```

# high-order fuction
## map
```
print(list(map(lambda a: a+1, range(5))))
[1, 2, 3, 4, 5]
```
## reduce
```
from functools import reduce
print(reduce(lambda a, b: a+b, range(5)))
10
```
