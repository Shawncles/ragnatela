package blockchain

type BlockChain interface {
		// chain related
	//GetBlockHash <height>
	//GetBlock <hash>
	//SubmitTransaction <tx hex code>
		// wallet related
	ListBalances(address string)
	//(balance float64, err error)
}

	//type(
	//	Chain struct{
	//		Chain		string `json:"chain"`
	//		Address		string `json:"address"`
	//		Name
	//	}
	//	Name interface {
	//		GetName() string
	//	}
	//)
