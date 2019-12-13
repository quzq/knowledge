VSCodeで絵文字のオートコンプリート＆プレビュー

# VSCode拡張「Markdown Emoji」をインストール
markdownプレビューにマークダウン絵文字がプレビュー表示されるようになります。

# VSCode拡張「:emojisense:」をインストール
Extentionsから「emojisense」をキーワードに検索し、「:emojisense:」をインストールします。

# 機能
### エディタ内で、github emoji のプレビュー
エディタ内で`:smile:`と入力すると、以下のようにプレビュー付きで表示されます。

```
:smile😄:
```

### unicode emoji の入力補完
コロン（:）を入力して、unicode emoji の入力補完ができます。

### github emoji の入力補完
ダブルコロン（::）を入力して、github等で使われているemojiの入力補完ができます。

### 絵文字ピッカーの表示

`Emojisense: Pick an emoji...`コマンドを実行するか、ショートカットキーにより、絵文字ピッカーを表示できます。

ショートカットキー:
- Mac：cmd + i
- Linux：ctrl + alt + i
- Windows：ctrl + i 

# :emojisense:拡張の設定
いずれもデフォルトでは全て有効になっています。
- emojisense.unicodeCompletionsEnabled: 
  絵文字をUnicodeとして挿入する補完を有効にする。
  例）:smile_cat:->😸
- emojisense.markupCompletionsEnabled: 絵文字マークアップを挿入する補完を有効にする。例）::smile_cat->:smile_cat:
- emojisense.showOnColon: コロンを入力すると、絵文字補完を自動的に表示。


