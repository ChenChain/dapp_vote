generate_sol_go_code:
	solc --abi --bin -o ./sol/vote --overwrite ./sol/vote/Vote.sol && \
	abigen --bin=./sol/vote/Vote.bin --abi=./sol/vote/Vote.abi --pkg=vote --out=./sol/vote/Vote.go


deploy_contract:
	go run ./cmd/deploy_contract.go