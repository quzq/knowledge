package main

import "fmt"

func main() {

  // 変数
  var foo string = "1234" // cannot use singlequote 
  fmt.Println(foo)

  // array
  bar := []int{2, 3, 8, 2, 3, 8, 2, 3, 8, 2, 3, 8}
  // arrayの要素を取り出す
  for i := 0; i < 5; i++ {
    fmt.Println(bar[i])
  }
  // arrayの要素を列挙
  for i, v := range bar {
    fmt.Println(i, v)
  }

  m := map[string]int{"key1":100, "key2":200}
  // mapの要素を列挙
  for k, v := range m {
    fmt.Println(k, v)
  }

  // 条件分岐
  if(foo == "2234"){
    fmt.Println("match")
  }else{
    fmt.Println("unmatch")
  }

//  func baz(a int) int{
//    return a*2
//  }

//  fmt.Println(baz(3))

}
