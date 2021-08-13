tesseract : 四次元超立方体




jTessBoxEditor
	tesseract2.0xと3.0x向けの自動学習ツール
	Site:http://vietocr.sourceforge.net/training.html
	DL:https://sourceforge.net/projects/vietocr/files/jTessBoxEditor/
	起動：java -Xms128m -Xmx1024m -jar jTessBoxEditor.jar
	
実行
	tesseract {ファイル名} stdout -l {言語識別子} 

	言語識別子：各言語の学習済ファイルの拡張子を除いた部分


座標ファイル（BOXファイル）の作成。教師データ
	tesseract <学習に使う画像> <言語>.<フォント>.exp<インデックス番号> -l jpn batch.nochop makebox
	ex)tesseract train_img00.bmp jpn.hiragi.exp0 -l jpn batch.nochop makebox
	行の構成
		<symbol> <left> <bottom> <right> <top> <page>
	jTessBoxEditorのTiff/Box GeneratorよりBOXファイルの自動生成ができる。
	作成されるファイル
	・.boxファイル
	・.tif(.boxとペアになる画像)
	・.font_properties
	
	BOXファイルの内容は、jTessBoxEditorのBox Editorにて確認できる。


trファイル（Shape Clusteringファイル。文字形状の特徴をクラスタリングする）
	.boxファイルと画像ファイルよりtrファイルを生成。
	tesseract {フォント画像ファイル} {言語}.{フォント}.exp{インデックス番号] nobatch box.train.stderr
	

トレーニングデータ作成

	unicharsetファイル作成
		unicharset_extractor {.boxファイル名}	※スペース区切りで複数指定可
		[作成ファイル]
			unicharset

	font_propertiesファイル作成（手入力等）
		trainingに使用するフォントの概要。
		内容は、[フォント名] [italic] [bold] [fixed] [serif] [fraktur]で構成される。

	学習する
		mftraining -F {.font_propertiesファイル} -U unicharset -O {.unicharsetファイル} {trファイル}
			C:\Users/KazunoriYamaguchi/Documents/tesseract/jTessBoxEditor/tesseract-ocr/mftraining
		作成されるファイル
		inttemp
		pffmtable
		shapetable

		cntraining {trファイル}
		作成されるファイル
		normproto

		wordlist2dawg dig.frequent_words_list dig.freq-dawg dig.unicharset
		wordlist2dawg dig.words_list dig.word-dawg dig.unicharset

ファイル結合
	結合して.traineddataファイルを作成する
	combine_tessdata {言語}.

	※以下のコマンドで結合したtraindedataを展開できる
	combine_tessdata -u {.traindedata} {言語}.



recorder: unicode compressor


