package ethers

import "github.com/blocklords/go-core/entity"

func GetEnvironmentByChainID(chainId string) entity.Environment {
	switch chainId {
	case "1", "56", "1285", "1284", "8453", "13371", "137":
		return entity.EMain
	default:
		return entity.EBeta
	}
}
