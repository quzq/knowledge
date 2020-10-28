gitコマンドで、未追跡ファイルを一括削除

# tl;dr
`git clean`コマンドで未追跡ファイルを一括削除できます。

# デモ
現在のカレントブランチを初期状態に戻したいとします。
`git status`コマンドでブランチの状態を確認すると、変更ファイルが１つ、未追跡ファイルが２つあります。

```
$ git status
On branch develop
Your branch is up to date with 'origin/develop'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   src/index.js

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        core.33
        core.43
```

`git checkout .`コマンドを実行したあと、再度`git status`コマンドでブランチの状態を確認すると、未追跡ファイルが２つ残っています。

```
$ git status
On branch develop
Your branch is up to date with 'origin/develop'.

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        core.33
        core.43
```

そこで、`clean`コマンドで未追跡ファイルを一括削除します。
```
$ git clean -df
Removing core.33
Removing core.43
```

一括削除前に、どのファイルが削除対象なのか確認することもできます。
```
$ git clean -n
Would remove core.33
Would remove core.43
```

未追跡ファイルがクリアされ、完全に初期のブランチの状態に戻りました。
```
$ git status
On branch develop
Your branch is up to date with 'origin/develop'.

nothing to commit, working tree clean
```

以上。