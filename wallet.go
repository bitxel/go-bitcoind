package bitcoind

type WalletInfo struct {
	WalletName            string  `json:"walletname"`
	WalletVersion         int32   `json:"walletversion"`
	Balance               float64 `json:"balance"`
	UnconfirmedBalance    float64 `json:"unconfirmed_balance"`
	ImmatureBalance       float64 `json:"immature_balance"`
	TxCount               int32   `json:"txcount"`
	KeyPoolOldest         int32   `json:"keypoololdest"`
	KeyPoolSize           int32   `json:"keypoolsize"`
	KeyPoolSizeHdInternal int32   `json:"keypoolsize_hd_internal"`
	PayTxFee              float64 `json:"paytxfee"`
	HdMasterKeyId         string  `json:"hdmasterkeyid"`
}
