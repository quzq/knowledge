prettierのオプション一覧

# 表記について
- version 2.0移行の仕様に準拠。
- 見出しは設定ファイル上の表記にしています。セミコロンの後は設定値のデータ型。
- CLI表記とは、CLIで指定する場合のオプション名。

# オプション一覧

### printWidth : int
- CLI表記: `--print-width`
- デフォルト: 80
- 行の折り返しの長さ。

### tabWidth : int
- CLI表記: `--tab-width`
- デフォルト: 2
- インデントのスペース数。

### useTabs : bool
- CLI表記: `--use-tabs`
- デフォルト: false
- スペースではなくタブでインデントする。

### semi : `bool`
- CLI表記: `--no-semi`
- デフォルト: true
- ステートメントの最後にセミコロンを付与。

### singleQuote : `bool`
- CLI表記: `--single-quote`
- デフォルト: false
- 二重引用符の代わりに単一引用符を使用する。

### Quote Props : `as-needed|consistent|preserve`
- CLI表記: `--`
- デフォルト: as-needed
- 

### JSX Quotes
- CLI表記: `--`
- デフォルト: 
- 

### trailingCommas : `es5|none|all`
- CLI表記: `--trailing_comma`
- デフォルト: es5
  - es5: 最後の要素にES5で有効なカンマを付与する。object, array, その他。
  - none: 最後の要素にカンマを付与しない。
  - all: 最後の要素にカンマを付与する。

### bracketSpacing : `bool`
- CLI表記: `--no-bracket-spacing`
- デフォルト: true
- オブジェクトリテラルの角かっこの内側にスペースを付与。例: `true` ⇒ { foo: bar } `false` ⇒ {foo: bar}

### JSX Brackets
- CLI表記: `--`
- デフォルト: 
- 

### arrowParens : `always|avoid`
- CLI表記: `--arrow-parents`
- デフォルト: always
- アロー関数の単一パラメーターを()で囲む。例：`always` ⇒ (x) => x `avoid` ⇒ x => x

### Range
- CLI表記: `--`
- デフォルト: 
- 

### Parser
- CLI表記: `--`
- デフォルト: 
- 

### File Path
- CLI表記: `--`
- デフォルト: 
- 

### Require Pragma
- CLI表記: `--`
- デフォルト: 
- 

### Insert Pragma
- CLI表記: `--`
- デフォルト: 
- 

### Prose Wrap
- CLI表記: `--`
- デフォルト: 
- 

### HTML Whitespace Sensitivity
- CLI表記: `--`
- デフォルト: 
- 

### vueIndentScriptAndStyle : `bool`
- CLI表記: `--vue-indent-script-and-style`
- デフォルト: false
- Vueファイルの`<script>`および`<style>`タグ内のコードをインデントする。

### endOfLine : `lf|crlf|cr|auto`
- CLI表記: `--end-of-line`
- デフォルト: lf 
- ファイルの改行コードを指定します。
  - lf: Lf 
  - crlf: CrLf
  - cr: Cr
  - auto: 既存の改行コードを維持。複数の改行コードが混在する場合は、最初の行の改行コードを適用する。
  
### Embedded Language Formatting
- CLI表記: `--`
- デフォルト: 
- 

