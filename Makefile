BINARY=main

engine:
	go build -o ${BINARY}

start:
	go run app/api/main.go

migrate:
	go run migrations/main.go

seed:
	go run seeders/main.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t login-go:v1 .

run:
	docker-compose up --build -d

stop:
	docker-compose down