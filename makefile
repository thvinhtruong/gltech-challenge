generate:
	touch config.env
	echo "DB_HOST=localhost \nDB_PORT=3306 \nDB_USER=root \nDB_PASS=root \nDB_NAME=todolist" >> config.env;
run:
	go run cmd/main.go;
