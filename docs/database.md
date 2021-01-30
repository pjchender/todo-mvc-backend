# Database for DEV

## 使用 homebrew

安裝 postgreSQL

```bash
$ brew install postgresql
$ brew services start postgresql
```

## 建立使用者

```bash
$ psql postgres  # 登入 DB
```

透過 SQL 建立名為 `postgres` 的 superuser（預設的 superuser 會是系統的使用者）：

```sql
-- create superuser of "postgres"
CREATE ROLE postgres LOGIN SUPERUSER CREATEDB CREATEROLE REPLICATION BYPASSRLS;
GRANT ALL ON ALL TABLES IN SCHEMA "public" TO postgres;
```

## 建立 Database

```bash
$ createdb todo_mvc -O postgres -E utf8
```

