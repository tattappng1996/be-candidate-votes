swaggen:
	swag init -g ./routes/swagger.go -o ./docs/v1

mockgen:
	mockgen -package=mocks -source=services/v1/service/srv.go -destination=mocks/services/service_v1_mock.go
	mockgen -package=mocks -source=adapter/repositories/tb/tb_repo.go -destination=mocks/repositories/tb_repo_mock.go
	mockgen -package=mocks -source=adapter/repositories/thirdparties/repo.go -destination=mocks/repositories/thirdparties_repo_mock.go

srv-test:
	go test -v ./services/v1/service/

srv-coverage:
	go test -coverprofile=./services/v1/service/services_coverage.out ./services/v1/service/ -v
	go tool cover -html=./services/v1/service/services_coverage.out