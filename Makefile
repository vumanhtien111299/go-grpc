init:
	go get entgo.io/ent
	go get github.com/go-sql-driver/mysql
	go get github.com/google/wire
	go get github.com/spf13/viper

db:
	 docker run --name mysql -e MYSQL_ROOT_PASSWORD=123@123 -e MYSQL_DATABASE=MYdb -p 3306:3306 -d mysql

ent.install:
	go install entgo.io/ent/cmd/ent@latest

ent.init:
	go run entgo.io/ent/cmd/ent init --target ent/schema User

ent.gen:
	go generate ./ent/...

wire.gen:
	wire ./cmd/...

wire.install:
	go install github.com/google/wire/cmd/wire@latest
