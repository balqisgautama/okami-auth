set GeneralConfiguration=D:/okami/playground/auth/config/
set HOST=http://localhost
set PORT=7001
set VERSION=1.0.0
set RESOURCE_ID=oauth
set PREFIX_PATH=okami/auth
set DB_CONNECTION=user=postgres password=bg1603 dbname=okami sslmode=disable host=localhost port=5432 TimeZone=Asia/Jakarta
set DB_SCHEMA=oauth
set DB_VIEW_CONNECTION=user=postgres password=bg1603 dbname=okami sslmode=disable host=localhost port=5432
set DB_VIEW_SCHEMA=oauth
set JWT_KEY=5892fbbd-c740-4cb8-bf01-4c0ec3d72001

go run main.go development