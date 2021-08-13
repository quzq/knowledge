react にて、画像リンク切れの際、imgタグを非表示にする

 
通常のhtmlにおいて、リンク切れのimgタグを非表示にできます。

```html
<img src="original.png" alt="title" onerror="this.style.display='none'"/>
```

reactではこれができないため、以下のように`onError`ハンドラを作成します。

```Javascript
<img src={`${process.env.PUBLIC_URL}/logo192_.png`} onError={e => e.target.style.display = 'none'} />
```

# サンプル
🚀以下のサンプルは、`create-react-app`直後の`App.js`を編集したものです。

```Javascript:App.js
import React from 'react';

function App() {
  return (
    <div className="App">
      <img border="1" src={`${process.env.PUBLIC_URL}/logo192.png`} />
      <img border="1" src={`${process.env.PUBLIC_URL}/logo192_.png`} />
      <img border="1" src={`${process.env.PUBLIC_URL}/logo192_.png`} onerror="this.style.display='none'" />
      <img border="1" src={`${process.env.PUBLIC_URL}/logo192_.png`} onError={e => e.target.style.display = 'none'} />
    </div>
  );
}

export default App;
```

# 参考
📖[Hide broken image link in Semantic UI React](https://stackoverflow.com/questions/49997068/hide-broken-image-link-in-semantic-ui-react)
📖リンク切れの時に代替画像を表示したいときは[こちら](https://qiita.com/m3816/items/56b9b6b6f340265cbfab)。
