package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			SalesContract                  string      `json:"SalesContract"`
			SalesContractType              string      `json:"SalesContractType"`
			SalesOrganization              string      `json:"SalesOrganization"`
			DistributionChannel            string      `json:"DistributionChannel"`
			OrganizationDivision           string      `json:"OrganizationDivision"`
			SalesGroup                     string      `json:"SalesGroup"`
			SalesOffice                    string      `json:"SalesOffice"`
			SalesDistrict                  string      `json:"SalesDistrict"`
			SoldToParty                    string      `json:"SoldToParty"`
			CreationDate                   string      `json:"CreationDate"`
			LastChangeDate                 string      `json:"LastChangeDate"`
			PurchaseOrderByCustomer        string      `json:"PurchaseOrderByCustomer"`
			CustomerPurchaseOrderDate      string      `json:"CustomerPurchaseOrderDate"`
			SalesContractDate              string      `json:"SalesContractDate"`
			TotalNetAmount                 string      `json:"TotalNetAmount"`
			TransactionCurrency            string      `json:"TransactionCurrency"`
			SDDocumentReason               string      `json:"SDDocumentReason"`
			PricingDate                    string      `json:"PricingDate"`
			IncotermsClassification        string      `json:"IncotermsClassification"`
			CustomerPaymentTerms           string      `json:"CustomerPaymentTerms"`
			PaymentMethod                  string      `json:"PaymentMethod"`
			SalesContractValidityStartDate string      `json:"SalesContractValidityStartDate"`
			SalesContractValidityEndDate   string      `json:"SalesContractValidityEndDate"`
			SalesContractSignedDate        string      `json:"SalesContractSignedDate"`
			ReferenceSDDocument            string      `json:"ReferenceSDDocument"`
			ReferenceSDDocumentCategory    string      `json:"ReferenceSDDocumentCategory"`
			SalesDocApprovalStatus         string      `json:"SalesDocApprovalStatus"`
			SalesContractApprovalReason    string      `json:"SalesContractApprovalReason"`
			OverallSDProcessStatus         string      `json:"OverallSDProcessStatus"`
			TotalCreditCheckStatus         string      `json:"TotalCreditCheckStatus"`
			OverallSDDocumentRejectionSts  string      `json:"OverallSDDocumentRejectionSts"`
			ToHeaderPartner struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Partner"`
			ToItem                         struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
		} `json:"results"`
	} `json:"d"`
}
