generate:
	touch config.env
	echo "DB_HOST=localhost \n DB_PORT=3306 \n DB_USER=root \n DB_PASS=root \n DB_NAME=todolist" >> config.env;
run:
	go run cmd/main.go;
