install-mockgen:
	@go install go.uber.org/mock/mockgen@latest
mockgen:
	@mockgen -destination=application/mocks/application.go -source=application/product.go