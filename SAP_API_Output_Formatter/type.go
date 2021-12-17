package sap_api_output_formatter

type SalesContract struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	SalesContract string `json:"sales_contract"`
	Deleted       bool   `json:"deleted"`
}    
    
type Header struct {
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
    ToHeaderPartner                string      `json:"to_Partner"`
    ToItem                         string      `json:"to_Item"`	
}

type Item struct {
	SalesContract                  string      `json:"SalesContract"`
	SalesContractItem              string      `json:"SalesContractItem"`
	SalesContractItemCategory      string      `json:"SalesContractItemCategory"`
	SalesContractItemText          string      `json:"SalesContractItemText"`
	PurchaseOrderByCustomer        string      `json:"PurchaseOrderByCustomer"`
	Material                       string      `json:"Material"`
	MaterialByCustomer             string      `json:"MaterialByCustomer"`
	PricingDate                    string      `json:"PricingDate"`
	RequestedQuantity              string      `json:"RequestedQuantity"`
	RequestedQuantityUnit          string      `json:"RequestedQuantityUnit"`
	ItemGrossWeight                string      `json:"ItemGrossWeight"`
	ItemNetWeight                  string      `json:"ItemNetWeight"`
	ItemWeightUnit                 string      `json:"ItemWeightUnit"`
	ItemVolume                     string      `json:"ItemVolume"`
	ItemVolumeUnit                 string      `json:"ItemVolumeUnit"`
	TransactionCurrency            string      `json:"TransactionCurrency"`
	NetAmount                      string      `json:"NetAmount"`
	MaterialGroup                  string      `json:"MaterialGroup"`
	MaterialPricingGroup           string      `json:"MaterialPricingGroup"`
	Batch                          string      `json:"Batch"`
	ProductionPlant                string      `json:"ProductionPlant"`
	StorageLocation                string      `json:"StorageLocation"`
	ShippingPoint                  string      `json:"ShippingPoint"`
	IncotermsClassification        string      `json:"IncotermsClassification"`
	CustomerPaymentTerms           string      `json:"CustomerPaymentTerms"`
	SalesDocumentRjcnReason        string      `json:"SalesDocumentRjcnReason"`
	ItemBillingBlockReason         string      `json:"ItemBillingBlockReason"`
	WBSElement                     string      `json:"WBSElement"`
	ProfitCenter                   string      `json:"ProfitCenter"`
	ReferenceSDDocument            string      `json:"ReferenceSDDocument"`
	ReferenceSDDocumentItem        string      `json:"ReferenceSDDocumentItem"`
	SDProcessStatus                string      `json:"SDProcessStatus"`
	SalesContractValidityStartDate string      `json:"SalesContractValidityStartDate"`
	SalesContractValidityEndDate   string      `json:"SalesContractValidityEndDate"`
	SalesContractSignedDate        string      `json:"SalesContractSignedDate"`
	ToItemPricingElement           string      `json:"to_PricingElement"`
}

type ToHeaderPartner struct {
	SalesContract               string `json:"SalesContract"`
	PartnerFunction             string `json:"PartnerFunction"`
	Customer                    string `json:"Customer"`
	Supplier                    string `json:"Supplier"`
}

type ToItem struct {
	SalesContract                  string      `json:"SalesContract"`
	SalesContractItem              string      `json:"SalesContractItem"`
	SalesContractItemCategory      string      `json:"SalesContractItemCategory"`
	SalesContractItemText          string      `json:"SalesContractItemText"`
	PurchaseOrderByCustomer        string      `json:"PurchaseOrderByCustomer"`
	Material                       string      `json:"Material"`
	MaterialByCustomer             string      `json:"MaterialByCustomer"`
	PricingDate                    string      `json:"PricingDate"`
	RequestedQuantity              string      `json:"RequestedQuantity"`
	RequestedQuantityUnit          string      `json:"RequestedQuantityUnit"`
	ItemGrossWeight                string      `json:"ItemGrossWeight"`
	ItemNetWeight                  string      `json:"ItemNetWeight"`
	ItemWeightUnit                 string      `json:"ItemWeightUnit"`
	ItemVolume                     string      `json:"ItemVolume"`
	ItemVolumeUnit                 string      `json:"ItemVolumeUnit"`
	TransactionCurrency            string      `json:"TransactionCurrency"`
	NetAmount                      string      `json:"NetAmount"`
	MaterialGroup                  string      `json:"MaterialGroup"`
	MaterialPricingGroup           string      `json:"MaterialPricingGroup"`
	Batch                          string      `json:"Batch"`
	ProductionPlant                string      `json:"ProductionPlant"`
	StorageLocation                string      `json:"StorageLocation"`
	ShippingPoint                  string      `json:"ShippingPoint"`
	IncotermsClassification        string      `json:"IncotermsClassification"`
	CustomerPaymentTerms           string      `json:"CustomerPaymentTerms"`
	SalesDocumentRjcnReason        string      `json:"SalesDocumentRjcnReason"`
	ItemBillingBlockReason         string      `json:"ItemBillingBlockReason"`
	WBSElement                     string      `json:"WBSElement"`
	ProfitCenter                   string      `json:"ProfitCenter"`
	ReferenceSDDocument            string      `json:"ReferenceSDDocument"`
	ReferenceSDDocumentItem        string      `json:"ReferenceSDDocumentItem"`
	SDProcessStatus                string      `json:"SDProcessStatus"`
	SalesContractValidityStartDate string      `json:"SalesContractValidityStartDate"`
	SalesContractValidityEndDate   string      `json:"SalesContractValidityEndDate"`
	SalesContractSignedDate        string      `json:"SalesContractSignedDate"`
	ToItemPricingElement           string      `json:"to_PricingElement"`
}

type ToItemPricingElement struct {
	SalesContract                 string `json:"SalesContract"`
	SalesContractItem             string `json:"SalesContractItem"`
	PricingProcedureStep          string `json:"PricingProcedureStep"`
	PricingProcedureCounter       string `json:"PricingProcedureCounter"`
	ConditionType                 string `json:"ConditionType"`
	PricingDateTime               string `json:"PricingDateTime"`
	ConditionCalculationType      string `json:"ConditionCalculationType"`
	ConditionBaseValue            string `json:"ConditionBaseValue"`
	ConditionRateValue            string `json:"ConditionRateValue"`
	ConditionCurrency             string `json:"ConditionCurrency"`
	ConditionQuantity             string `json:"ConditionQuantity"`
	ConditionQuantityUnit         string `json:"ConditionQuantityUnit"`
	ConditionCategory             string `json:"ConditionCategory"`
	PricingScaleType              string `json:"PricingScaleType"`
	ConditionRecord               string `json:"ConditionRecord"`
	ConditionSequentialNumber     string `json:"ConditionSequentialNumber"`
	TaxCode                       string `json:"TaxCode"`
	ConditionAmount               string `json:"ConditionAmount"`
	TransactionCurrency           string `json:"TransactionCurrency"`
	PricingScaleBasis             string `json:"PricingScaleBasis"`
	ConditionScaleBasisValue      string `json:"ConditionScaleBasisValue"`
	ConditionScaleBasisUnit       string `json:"ConditionScaleBasisUnit"`
	ConditionScaleBasisCurrency   string `json:"ConditionScaleBasisCurrency"`
	ConditionIsManuallyChanged    bool   `json:"ConditionIsManuallyChanged"`
}
