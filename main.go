//go:generate docker run --rm -w /local -v ${PWD}:/local openapitools/openapi-generator-cli:latest generate  -i ./pkg/api/controllers/swagger.yaml -g go -o ./client --git-user-id=numary --git-repo-id=ledger -p packageVersion=latest -p isGoSubmodule=true -p packageName=client -t ./sdk/templates/go
package main

import (
	"github.com/numary/ledger/cmd"
)

func main() {
	cmd.Execute()
}
