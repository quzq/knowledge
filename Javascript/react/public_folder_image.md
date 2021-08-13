React にて、publicフォルダ内の画像を参照する

# tl;dr
`process.env.PUBLIC_URL`もしくは`window.location.origin`から`public`フォルダのパスを取得できます。

# サンプル
🚀以下のサンプルは、`create-react-app`直後の`App.js`を編集したものです。

```Javascript:App.js
import React from 'react';

function App() {
  return (
    <div className="App">
        <img src={`${process.env.PUBLIC_URL}/logo192.png`} />
        <img src={`${window.location.origin}/logo192.png`} />
    </div>
  );
}

export default App;
```

# 参考
📖[ReactJS and images in public folder](https://stackoverflow.com/questions/47196800/reactjs-and-images-in-public-folder)