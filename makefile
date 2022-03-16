build:
	./build.sh
up:
	docker-compose up --build -d
down:
	docker-compose down 
clean:
	./clean.sh
console:
	./aql-console.sh
