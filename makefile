run:
	go run main.go

mock:
	mockgen -source=database/interface.go -destination=database/mocks/db_mock.go -package=mocks
