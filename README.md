# Todo MVC Backend

## 專案建立過程

- 建立資料夾結構
- 建立 database pkg
- 建立 router pkg
- 建立 config pkg
- 建立 `internal/setup` pkg
- 建立 global pkg

## 備註

- 暫時沒有建立 `pkg/app`，未來若有需要再加入。因此在回傳 response 時，不用呼叫 `app.NewResponse()`，也不需要呼叫 `app.BindAndValue()`

## 參考

- [Go 語言編程之旅：一起用 Go 做項目](https://github.com/go-programming-tour-book/blog-service) @ github
- [gotify/server](https://github.com/gotify/server)
- [xinliangnote/go-gin-api](https://github.com/xinliangnote/go-gin-api) @ github

