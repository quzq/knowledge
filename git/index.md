### git
- ファイルを管理対象から除外
  - git rm --cached {ファイル名}
  - 一度管理対象になったファイルは、このキャッシュが残っている限りgitignoreできない

### git
- reflog
  - HEADやブランチの移動履歴を確認する。出力の`HEAD@{num}`(numは整数)は、commitIDのエイリアスとして使用できる。
  - `git reset --hard HEAD@{0}`を行うことで、コマンドのundo代わりになる。
  
### git
- 現在のリモートアドレスを確認
  - git remote -v
- リモートの向き先を変更
  - git remote set-url origin {new url}

### git
- チェリーピック
  - git cherry-pick {コミットID}

### git
- git push --set-upstream origin master

### git
- 編集の取り消し
  - git checkout {filename}

### git
- ブランチのposition移動:
  - git reset --hard コミットID

### git
- ステージング
  - git add {ファイル名}
- スタッシュ内容のdiff表示
  - git diff {スタッシュ番号} {ファイル名}

### git
- ローカルブランチの削除
  - git branch -D {ローカルブランチ名}

### git
- git checkout -b (ブランチ名)
  - カレントのブランチから新しいブランチを派生させる

### git
- HEAD
  - 現在作業しているブランチの先頭コミットを指すポインタ。コミットIDのエイリアスとしても機能する。
- ブランチ
  - 特定のコミットを指す 40 文字の SHA-1 チェックサムだけを記録したシンプルなファイル
- HEAD~
  - HEADの1世代前の親コミット。２世代前はHEAD~~
- HEAD^
  - 複数ある親コミットの中のコミット。HEAD^とHEAD^2は同じ世代だが違うコミット
- reset
  - git reset コミットID
    - HEADポインタを指定のコミットIDに指定する。git reset HEADは、ポインタに変更はないが、その際にINDEXが再生成されることにより、git add が解除される。

### git
- add -i
  - 対話式ステージング
- add .
  - ワーキングツリーにあるファイルをまとめてadd
- 特定ファイルのrevert
  - git checkout {commit_id} {file}

### git
- branch変更直後で、内容が変わっていないはずなのに、git statusすると、変更済として表示されるファイルがある場合
  - ローカルファイルの改行コードがcrlfになっている。手動で改行をlfに変更し（vscodeで一括変更できる）、git addすると治る。

### git
- branch変更直後で、内容が変わっていないはずなのに、git statusすると、変更済として表示されるファイルがある場合
  - ローカルファイルの改行コードがcrlfになっている。手動で改行をlfに変更し（vscodeで一括変更できる）、git addすると治る。

  git
    remote add {alias} {リポジトリurl}
      リモートリポジトリ設定を追加
    remote rm {alies} 
      リモートリポジトリ設定を削除
    push {label} {ブランチ名}  
      ラベルのリモート先に対し、指定ブランチをpushする。
    push -u (--set-upstream) 
      labelを上流ブランチとして設定しながらpush。初回にuオプションをつけない場合、permission deniedが発生する。
    config --list
      全ての設定を表示
    リモート追跡ブランチ（remote-tracking branch）
      ローカルリポジトリにあって、リモートリポジトリの「あるブランチ（例えばmaster）の状態」をそのまま表しているブランチ。
    上流ブランチ(upstream branch)
      引数なしで git pull したとき対象になるブランチ


### git
    origin
      リモートリポジトリのアクセス先に対し、gitがデフォルトでつける名前。`git push -u`を行うと、リモートリポジトリ先origin、リモートブランチ名masterがデフォルトで設定される。


  ### git
  - カレントをブランチ名へ移動
    - git checkout nランチ名
  - ブランチのマージ
    - （マージ先をカレントににする）
    - git merge --no-ff マージ元ブランチ名
  - コミットの打ち消し
    - git revert <commit>

  ### git
  - git reset --hard HEAD^  →直前のコミットを取り消し(修正も消す)
  - git reset --soft HEAD^  →直前のコミットを取り消し(修正はworkspaceに残す)

  ### git
  - ブランチ分岐
    -（カレントブランチを作成元ブランチにしておく）
    - git checkout -b （作成するブランチ名）
    - git push -u origin （作成したブランチ名）

  ### git 
  クライアントのリモート情報のリフレッシュ
  - git remote update -p

  ### git 
- 最後のcommitの取り消し
  - git reset --soft HEAD^

### git
- stash popのキャンセル
  - git checkout --ours .

### git
- リモートブランチからローカルブランチを作成する:
  - git checkout -b ローカルに作成するブランチ名 origin/作成元のリモートのブランチ名
- ローカルブランチをリモートに反映:
  - git push -u origin 作成したブランチ名


リポジトリの比較
git diff {commitID} {commitID}

### git 
- rebase
  - 並行して開発していたブランチを同期させる
    - カレントのブランチをmasterと同期させたい場合、カレントブランチ上で、`git rebase origin/master`とする。取り消しは、`git reset`を使う
  - 複数のコミットを１ つにまとめる
    - git rebase -i [ひとまとめにする地点の一つ前のコミットID]









