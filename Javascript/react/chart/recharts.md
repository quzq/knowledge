
# React向けチャートライブラリRechartsで円グラフ

📊グラフの種類

[AreaChart](http://recharts.org/en-US/api/AreaChart):面グラフ。折れ線グラフに基づき、定量データを表示したグラフ。
[BarChart](http://recharts.org/en-US/api/BarChart):棒グラフ。四角い棒の長さで何らかの値を表現するグラフ。
[LineChart](http://recharts.org/en-US/api/LineChart):折れ線グラフ。散布図の一種であり、プロットされた点を直線でつないだグラフ。
[ComposedChart](http://recharts.org/en-US/api/ComposedChart):折れ線グラフ、面グラフ、棒グラフで構成されるグラフ。線のような単一のタイプのチャートを描画する場合は、LineChartを使う。
[PieChart](http://recharts.org/en-US/api/PieChart):円グラフ。丸い図形を扇形に分割し、何らかの構成比率を表したグラフ。
[RadarChart](http://recharts.org/en-US/api/RadarChart):レーダーチャート。複数の項目の大きさを一見して比較することのできるグラフ。
[RadialBarChart](http://recharts.org/en-US/api/RadialBarChart):円弧で表現する棒グラフ。
[ScatterChart](http://recharts.org/en-US/api/ScatterChart):散布図（分布図）。縦軸、横軸に2項目の量や大きさ等を対応させ、データを点でプロットしたグラフ。
[FunnelChart](http://recharts.org/en-US/api/FunnelChart):ファンネルチャート。販売プロセスの段階を表し、各段階の潜在的な収益額を示すためによく使用されるグラフ。
[Treemap](http://recharts.org/en-US/api/Treemap):ツリーマップ。ネストされた長方形のセットとして階層状のツリー構造のデータを表示したもの。

`<PieChart />`
  プロパティ: width, height, margin 
  イベント: onClick, onMouseEnter, onMouseLeave

  親コンポーネント
    `<ResponsiveContainer />`:

  子コンポーネント
    `<PolarAngleAxis />`:
    `<PolarRadiusAxis />`:
    `<PolarGrid />`:
    `<Legend />`:
    `<Tooltip />`:
    `<Pie />`:
    `<Customized />`:


`<Pie />`
  - プロパティ

    - cx: 中心のx座標。パーセンテージを設定した場合、最終値はコンテナ幅のパーセンテージを乗算することにより取得されます。
    - cy: 中心のy座標。パーセンテージを設定すると、コンテナの高さのパーセンテージを乗算することで最終的な値が取得されます。
    - innerRadius: すべてのセクターの内半径。パーセンテージを設定した場合、最終値は、幅、高さ、cx、cyで計算されるmaxRadiusのパーセンテージを乗算することにより取得されます。
    - outerRadius: すべてのセクターの外半径。パーセンテージを設定した場合、最終値は、幅、高さ、cx、cyで計算されるmaxRadiusのパーセンテージを乗算することにより取得されます。
    - startAngle: 最初のセクターの開始角度。
    - endAngle: 最後のセクターの終了角度。これは、startAngleと等しくないはずです。
    - minAngle: 各非ゼロデータの最小角度。
    - paddingAngle: 2つのセクター間の角度。
    - nameKey: 各セクターの名前のキー。
    - dataKey: 各セクターの値のキー。
    - legendType: 凡例のアイコンのタイプ。`'none'`に設定すると、凡例項目はレンダリングされません。`'line'` `'square'` `'rect` `'circle'` `'cross'` `'diamond'` `'square'` `'star'` `'triangle'` `'wye'` `'none'`
    - label:trueの時、値がラベルとして表示されます。デフォルトはfalse。
    - labelLine:trueの場合、セクターとラベルを結ぶ線を描画。
    - data:各要素がオブジェクトであるソースデータ。
    - activeIndex:Pieのアクティブセクターのインデックス。このオプションは、マウスイベントハンドラーで変更できます。
    - activeSphape:アクティブなセクターの形状。
    - isAnimationActive:trueに設定すると、パイのアニメーションが有効化。
    - animationBegin:アニメーションをいつ開始するかを指定します。このオプションの単位はmsです。
    - animetionDuration:アニメーションの継続時間を指定します。このオプションの単位はmsです。
    - animationEasing: イージング関数のタイプ。`'ease'` `'ease-in'` `'ease-out'` `'ease-in-out'` `'linear'` `Function` 



`<Legend />`
- プロパティ
  - width, height:
  - verticalAlign:
  - iconSize:
  - iconType: 

`<ToolTip />`
- プロパティ
  - separator:


