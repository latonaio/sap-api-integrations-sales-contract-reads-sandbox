package responses

type ToHeaderPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SalesContract   string `json:"SalesContract"`
			PartnerFunction string `json:"PartnerFunction"`
			Customer        string `json:"Customer"`
			Supplier        string `json:"Supplier"`
		} `json:"results"`
	} `json:"d"`
}
