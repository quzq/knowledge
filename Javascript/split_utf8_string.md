# UTF-8文字列を1文字単位に分割

## tl;dr
split('') ではUTF-8の4バイト文字を分割できません。
すべての文字を分割するには `spred operetor` 等を使用します。

## split('')を使用した場合

```Javascript
console.log('3バイト文字は分割できる'.split(''));

// output: [ '3', 'バ', 'イ', 'ト', '文', '字', 'は', '分', '割', 'で', 'き', 'る' ]

console.log('4バイト文字（𠀋）の場合'.split(''));

// output: [ '4',  'バ', 'イ', 'ト', '文', '字', '（', '�',  '�',  '）', 'の', '場', '合' ]

```
4バイト文字である `𠀋` が正しく分割されていません。

## 正しい方法

### spread operator を使う
```Javascript
console.log([...'4バイト文字（𠀋）の場合']);
  
// output: [ '4',  'バ', 'イ',  'ト', '文', '字',  '（', '𠀋', '）',  'の', '場', '合' ]
```


### Array.from を使う
```Javascript
console.log(Array.from('4バイト文字（𠀋）の場合'));
  
// output: [ '4',  'バ', 'イ',  'ト', '文', '字',  '（', '𠀋', '）',  'の', '場', '合' ]
```

### RegExpのuフラグを使う
```Javascript
console.log('4バイト文字（𠀋）の場合'.split(/(?=[\s\S])/u));
  
// output: [ '4',  'バ', 'イ',  'ト', '文', '字',  '（', '𠀋', '）',  'の', '場', '合' ]
```


