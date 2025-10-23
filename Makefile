include .env
export

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: help
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]
	
# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
.PHONY: run/serv
## run/serv: run the cmd/sso application
run/serv:
	go run ./cmd/sso

