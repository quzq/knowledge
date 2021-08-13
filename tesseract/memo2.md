TesseractOCRについて
バージョン3系
	最新安定版： 3.05.02(2018/6/19) ※このバージョンは日本語の認識率に問題があるらしい
	認識エンジン：？
	日本語学習用データ
		https://github.com/tesseract-ocr/langdata/tree/master/jpn
	
バージョン4系
	最新： 4.0.0(2018/10/29) ※Ubuntu 18.04は4.00beta.1(2018/3/10)採用
	認識エンジン：ocrpy由来のLSTMベースのニューラルネットワークによる認識エンジン
	学習済みデータ（trained models）：
		tessdata_fast:新認識エンジン用学習済みモデル_速度重視
		tessdata_best:新認識エンジン用学習済みモデル_認識精度重視
		tessdata:tessdata_best + 旧式エンジン用データ。（コマンドパラメータ --oemにて、新旧エンジンを同時指定できる）
		※fastとbestの違いはニューラルネットワーク（LSTM）の学習時の重みの精度（integer/float）らしい
		※現状、基本的にHomebrew の--HEADオプションとかUbuntu or Debianでパッケージとして提供されているのはtessdata_fast。
		※scriptサブディレクトリに「Japanese.traineddata」という学習データが提供される。これは日本語と英語の学習用テキストを混在した状態で学習させたものらしい。	日本語学習用データ
		https://github.com/tesseract-ocr/langdata_lstm/tree/master/jpn

メモ
	出力結果に余計な空白が含まれる問題の対処方法。-cオプションでパラメータをセットする。
	例）tesseract -l jpn -c preserve_interword_spaces=1 sample.jpg  stdout
	※最新版では解決済？（windows版4.0.0では発生しない）

学習について
	公式：https://github.com/tesseract-ocr/tesseract/wiki/TrainingTesseract-4.00
	学習用ファイル
		forbidden_characters
			認識させない文字範囲をunicodeで指定する
		desired_characters
			認識したい文字。jpn.training_textを作成するのに使用する。
		jpn.training_text
		
		training/language-specific.sh
			JPN_FONTSの設定で対応フォントを増減できる。
		jpn.traineddata
			tesstrain.shにて学習させることで生成される結果ファイル。
		Unicharset
			文字セットを定義

	参考サイト
		Tesseract 4.0で日本語の認識をチューニングしよう
		https://qiita.com/masaoki/items/df665e285330d4da3cca
