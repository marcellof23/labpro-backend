migrate -path pkg/migrations -database "mysql://root:dorayaki@tcp(localhost:3306)/labpro-database?charset=utf8mb4&parseTime=True&loc=Local" -verbose up