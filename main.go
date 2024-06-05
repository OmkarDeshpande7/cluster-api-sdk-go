package main

import (
	"fmt"

	"github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure"
	capa "github.com/OmkarDeshpande7/cluster-api-sdk-go/infrastructure/capa"
)

func main() {
	var i infrastructure.InfraProvider
	i = &capa.AWSProvider{
		Client: nil,
	}
	fmt.Println(i)
}
