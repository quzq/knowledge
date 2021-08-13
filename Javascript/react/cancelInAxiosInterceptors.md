axiosのinterceptorsで、requestのキャンセルを行う


# tl;dr
`interceptors.request.use()`内で`axios.Cancel()`をthrowすることでリクエストを直前でキャンセルできます。

# 🚀サンプル
認証トークンがlocalStorageに保存されており、トークンがなければAPIリクエストを投げたくないというケースを考えてみます。

認証トークンを取得する関数は以下です。
```
const getToken = () => window.localStorage.getItem("token")
```





https://stackoverflow.com/questions/50461746/axios-how-to-cancel-request-inside-request-interceptor-properly




```
const axios = require('axios');


axios.interceptors.request.use(config => {
  if (!getToken()) {
    console.log("interceptors: no access token");
  } else {
    config.headers.Authorization = "Bearer " + getToken().accessToken;
    return config;
  }
});
```

```
async function getUser() {
  try {
    const response = await axios.get('/user?ID=12345');
    console.log(response);
  } catch (error) {
    console.error(error);
  }
}
```


```



// Add a request interceptor
axios.interceptors.request.use(function (config) {
    // Do something before request is sent
    return config;
  }, function (error) {
    // Do something with request error
    return Promise.reject(error);
  });

```