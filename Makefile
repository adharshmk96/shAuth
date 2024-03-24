keys:
	mkdir -p .keys
	openssl ecparam -genkey -name secp521r1 -noout -out .keys/ecdsa-private.pem
	openssl ec -in .keys/ecdsa-private.pem -pubout -out .keys/ecdsa-public.pem

swagger:
	swag init

migration:
	migrate create -ext sql -dir db/migrations -seq $(filter-out $@,$(MAKECMDGOALS))
%:
	@:

pre:
	go install github.com/swaggo/swag/cmd/swag@latest
	brew install golang-migrate

init: pre keys swagger

run:
	go run main.go