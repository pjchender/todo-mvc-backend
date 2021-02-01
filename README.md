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

## 部署到 heorku 的筆記

- 記得要設定 heroku 上的環境變數
  ```bash
  $ heroku config:set MODE=production
  ```
- 記得要建 `Procfile` 且需要是一次新的 commit （不能用 `amend` 加入）
- 記得要啟動 `$ heroku ps:scale web=1`
- 記得要 build 成 unix 的版本才能在 heroku 上執行
  ```bash
  $ env GOOS=linux GOARCH=amd64 go build -o bin/main -v .
  ```

## 參考

- [Go 語言編程之旅：一起用 Go 做項目](https://github.com/go-programming-tour-book/blog-service) @ github
- [gotify/server](https://github.com/gotify/server)
- [xinliangnote/go-gin-api](https://github.com/xinliangnote/go-gin-api) @ github

