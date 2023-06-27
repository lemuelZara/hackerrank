help:
	@echo "HackerRank Exercises"
	@echo ""
	@echo "test		Run unit tests"

test:
	@go test ./... -count 1 -cover -shuffle=on