package api

//go:generate npx swagger-cli bundle openapi.yaml --outfile openapi.gen.yaml --type yaml
//go:generate go run github.com/ogen-go/ogen/cmd/ogen -target ../internal/generated/api -clean openapi.gen.yaml
