// Package vendorOrders provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package vendorOrders

import (
	"time"
)

// Defines values for ImportDetailsInternationalCommercialTerms.
const (
	ImportDetailsInternationalCommercialTermsCarriageAndInsurancePaidTo ImportDetailsInternationalCommercialTerms = "CarriageAndInsurancePaidTo"

	ImportDetailsInternationalCommercialTermsCarriagePaidTo ImportDetailsInternationalCommercialTerms = "CarriagePaidTo"

	ImportDetailsInternationalCommercialTermsCostAndFreight ImportDetailsInternationalCommercialTerms = "CostAndFreight"

	ImportDetailsInternationalCommercialTermsCostInsuranceAndFreight ImportDetailsInternationalCommercialTerms = "CostInsuranceAndFreight"

	ImportDetailsInternationalCommercialTermsDeliverDutyPaid ImportDetailsInternationalCommercialTerms = "DeliverDutyPaid"

	ImportDetailsInternationalCommercialTermsDeliveredAtPlace ImportDetailsInternationalCommercialTerms = "DeliveredAtPlace"

	ImportDetailsInternationalCommercialTermsDeliveredAtTerminal ImportDetailsInternationalCommercialTerms = "DeliveredAtTerminal"

	ImportDetailsInternationalCommercialTermsExWorks ImportDetailsInternationalCommercialTerms = "ExWorks"

	ImportDetailsInternationalCommercialTermsFreeAlongSideShip ImportDetailsInternationalCommercialTerms = "FreeAlongSideShip"

	ImportDetailsInternationalCommercialTermsFreeCarrier ImportDetailsInternationalCommercialTerms = "FreeCarrier"

	ImportDetailsInternationalCommercialTermsFreeOnBoard ImportDetailsInternationalCommercialTerms = "FreeOnBoard"
)

// Defines values for ImportDetailsMethodOfPayment.
const (
	ImportDetailsMethodOfPaymentCollectOnDelivery ImportDetailsMethodOfPayment = "CollectOnDelivery"

	ImportDetailsMethodOfPaymentDefinedByBuyerAndSeller ImportDetailsMethodOfPayment = "DefinedByBuyerAndSeller"

	ImportDetailsMethodOfPaymentFOBPortOfCall ImportDetailsMethodOfPayment = "FOBPortOfCall"

	ImportDetailsMethodOfPaymentPaidByBuyer ImportDetailsMethodOfPayment = "PaidByBuyer"

	ImportDetailsMethodOfPaymentPaidBySeller ImportDetailsMethodOfPayment = "PaidBySeller"

	ImportDetailsMethodOfPaymentPrepaidBySeller ImportDetailsMethodOfPayment = "PrepaidBySeller"
)

// Defines values for ItemQuantityUnitOfMeasure.
const (
	ItemQuantityUnitOfMeasureCases ItemQuantityUnitOfMeasure = "Cases"

	ItemQuantityUnitOfMeasureEaches ItemQuantityUnitOfMeasure = "Eaches"
)

// Defines values for OrderPurchaseOrderState.
const (
	OrderPurchaseOrderStateAcknowledged OrderPurchaseOrderState = "Acknowledged"

	OrderPurchaseOrderStateClosed OrderPurchaseOrderState = "Closed"

	OrderPurchaseOrderStateNew OrderPurchaseOrderState = "New"
)

// Defines values for OrderDetailsPaymentMethod.
const (
	OrderDetailsPaymentMethodConsignment OrderDetailsPaymentMethod = "Consignment"

	OrderDetailsPaymentMethodCreditCard OrderDetailsPaymentMethod = "CreditCard"

	OrderDetailsPaymentMethodInvoice OrderDetailsPaymentMethod = "Invoice"

	OrderDetailsPaymentMethodPrepaid OrderDetailsPaymentMethod = "Prepaid"
)

// Defines values for OrderDetailsPurchaseOrderType.
const (
	OrderDetailsPurchaseOrderTypeConsignedOrder OrderDetailsPurchaseOrderType = "ConsignedOrder"

	OrderDetailsPurchaseOrderTypeNewProductIntroduction OrderDetailsPurchaseOrderType = "NewProductIntroduction"

	OrderDetailsPurchaseOrderTypeRegularOrder OrderDetailsPurchaseOrderType = "RegularOrder"

	OrderDetailsPurchaseOrderTypeRushOrder OrderDetailsPurchaseOrderType = "RushOrder"
)

// Defines values for OrderItemAcknowledgementAcknowledgementCode.
const (
	OrderItemAcknowledgementAcknowledgementCodeAccepted OrderItemAcknowledgementAcknowledgementCode = "Accepted"

	OrderItemAcknowledgementAcknowledgementCodeBackordered OrderItemAcknowledgementAcknowledgementCode = "Backordered"

	OrderItemAcknowledgementAcknowledgementCodeRejected OrderItemAcknowledgementAcknowledgementCode = "Rejected"
)

// Defines values for OrderItemAcknowledgementRejectionReason.
const (
	OrderItemAcknowledgementRejectionReasonInvalidProductIdentifier OrderItemAcknowledgementRejectionReason = "InvalidProductIdentifier"

	OrderItemAcknowledgementRejectionReasonObsoleteProduct OrderItemAcknowledgementRejectionReason = "ObsoleteProduct"

	OrderItemAcknowledgementRejectionReasonTemporarilyUnavailable OrderItemAcknowledgementRejectionReason = "TemporarilyUnavailable"
)

// Defines values for OrderItemStatusAcknowledgementStatusConfirmationStatus.
const (
	OrderItemStatusAcknowledgementStatusConfirmationStatusACCEPTED OrderItemStatusAcknowledgementStatusConfirmationStatus = "ACCEPTED"

	OrderItemStatusAcknowledgementStatusConfirmationStatusPARTIALLYACCEPTED OrderItemStatusAcknowledgementStatusConfirmationStatus = "PARTIALLY_ACCEPTED"

	OrderItemStatusAcknowledgementStatusConfirmationStatusREJECTED OrderItemStatusAcknowledgementStatusConfirmationStatus = "REJECTED"

	OrderItemStatusAcknowledgementStatusConfirmationStatusUNCONFIRMED OrderItemStatusAcknowledgementStatusConfirmationStatus = "UNCONFIRMED"
)

// Defines values for OrderStatusPurchaseOrderStatus.
const (
	OrderStatusPurchaseOrderStatusCLOSED OrderStatusPurchaseOrderStatus = "CLOSED"

	OrderStatusPurchaseOrderStatusOPEN OrderStatusPurchaseOrderStatus = "OPEN"
)

// Defines values for TaxRegistrationDetailsTaxRegistrationType.
const (
	TaxRegistrationDetailsTaxRegistrationTypeGST TaxRegistrationDetailsTaxRegistrationType = "GST"

	TaxRegistrationDetailsTaxRegistrationTypeVAT TaxRegistrationDetailsTaxRegistrationType = "VAT"
)

// Details of item quantity ordered
type AcknowledgementStatusDetails struct {

	// Details of quantity ordered.
	AcceptedQuantity *ItemQuantity `json:"acceptedQuantity,omitempty"`

	// The date when the line item was confirmed by vendor. Must be in ISO-8601 date/time format.
	AcknowledgementDate *time.Time `json:"acknowledgementDate,omitempty"`

	// Details of quantity ordered.
	RejectedQuantity *ItemQuantity `json:"rejectedQuantity,omitempty"`
}

// Address of the party.
type Address struct {

	// First line of the address.
	AddressLine1 string `json:"addressLine1"`

	// Additional address information, if required.
	AddressLine2 *string `json:"addressLine2,omitempty"`

	// Additional address information, if required.
	AddressLine3 *string `json:"addressLine3,omitempty"`

	// The city where the person, business or institution is located.
	City *string `json:"city,omitempty"`

	// The two digit country code. In ISO 3166-1 alpha-2 format.
	CountryCode string `json:"countryCode"`

	// The county where person, business or institution is located.
	County *string `json:"county,omitempty"`

	// The district where person, business or institution is located.
	District *string `json:"district,omitempty"`

	// The name of the person, business or institution at that address.
	Name string `json:"name"`

	// The phone number of the person, business or institution located at that address.
	Phone *string `json:"phone,omitempty"`

	// The postal code of that address. It conatins a series of letters or digits or both, sometimes including spaces or punctuation.
	PostalCode *string `json:"postalCode,omitempty"`

	// The state or region where person, business or institution is located.
	StateOrRegion *string `json:"stateOrRegion,omitempty"`
}

// Defines a date time interval according to ISO8601. Interval is separated by double hyphen (--).
type DateTimeInterval string

// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation. <br>**Pattern** : `^-?(0|([1-9]\d*))(\.\d+)?([eE][+-]?\d+)?$`.
type Decimal string

// Error response returned when the request is unsuccessful.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition.
	Message string `json:"message"`
}

// A list of error responses returned when a request is unsuccessful.
type ErrorList []Error

// The response schema for the getPurchaseOrder operation.
type GetPurchaseOrderResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList `json:"errors,omitempty"`
	Payload *Order     `json:"payload,omitempty"`
}

// The response schema for the getPurchaseOrders operation.
type GetPurchaseOrdersResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList `json:"errors,omitempty"`
	Payload *OrderList `json:"payload,omitempty"`
}

// The response schema for the getPurchaseOrdersStatus operation.
type GetPurchaseOrdersStatusResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList       `json:"errors,omitempty"`
	Payload *OrderListStatus `json:"payload,omitempty"`
}

// Import details for an import order.
type ImportDetails struct {

	// Types and numbers of container(s) for import purchase orders. Can be a comma-separated list if the shipment has multiple containers. HC signifies a high-capacity container. Free-text field, limited to 64 characters. The format will be a comma-delimited list containing values of the type: $NUMBER_OF_CONTAINERS_OF_THIS_TYPE-$CONTAINER_TYPE. The list of values for the container type is: 40'(40-foot container), 40'HC (40-foot high-capacity container), 45', 45'HC, 30', 30'HC, 20', 20'HC.
	ImportContainers *string `json:"importContainers,omitempty"`

	// Incoterms (International Commercial Terms) are used to divide transaction costs and responsibilities between buyer and seller and reflect state-of-the-art transportation practices. This is for import purchase orders only.
	InternationalCommercialTerms *ImportDetailsInternationalCommercialTerms `json:"internationalCommercialTerms,omitempty"`

	// If the recipient requests, contains the shipment method of payment. This is for import PO's only.
	MethodOfPayment *ImportDetailsMethodOfPayment `json:"methodOfPayment,omitempty"`

	// The port where goods on an import purchase order must be delivered by the vendor. This should only be specified when the internationalCommercialTerms is FOB.
	PortOfDelivery *string `json:"portOfDelivery,omitempty"`

	// Special instructions regarding the shipment. This field is for import purchase orders.
	ShippingInstructions *string `json:"shippingInstructions,omitempty"`
}

// Incoterms (International Commercial Terms) are used to divide transaction costs and responsibilities between buyer and seller and reflect state-of-the-art transportation practices. This is for import purchase orders only.
type ImportDetailsInternationalCommercialTerms string

// If the recipient requests, contains the shipment method of payment. This is for import PO's only.
type ImportDetailsMethodOfPayment string

// Details of quantity ordered.
type ItemQuantity struct {

	// Acknowledged quantity. This value should not be zero.
	Amount *int `json:"amount,omitempty"`

	// Unit of measure for the acknowledged quantity.
	UnitOfMeasure *ItemQuantityUnitOfMeasure `json:"unitOfMeasure,omitempty"`

	// The case size, in the event that we ordered using cases.
	UnitSize *int `json:"unitSize,omitempty"`
}

// Unit of measure for the acknowledged quantity.
type ItemQuantityUnitOfMeasure string

// Detailed description of items order status.
type ItemStatus []OrderItemStatus

// An amount of money, including units in the form of currency.
type Money struct {

	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation. <br>**Pattern** : `^-?(0|([1-9]\d*))(\.\d+)?([eE][+-]?\d+)?$`.
	Amount *Decimal `json:"amount,omitempty"`

	// Three digit currency code in ISO 4217 format. String of length 3.
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// Order defines model for Order.
type Order struct {

	// Details of an order.
	OrderDetails *OrderDetails `json:"orderDetails,omitempty"`

	// The purchase order number for this order. Formatting Notes: 8-character alpha-numeric code.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`

	// This field will contain the current state of the purchase order.
	PurchaseOrderState OrderPurchaseOrderState `json:"purchaseOrderState"`
}

// This field will contain the current state of the purchase order.
type OrderPurchaseOrderState string

// OrderAcknowledgement defines model for OrderAcknowledgement.
type OrderAcknowledgement struct {

	// The date and time when the purchase order is acknowledged, in ISO-8601 date/time format.
	AcknowledgementDate time.Time `json:"acknowledgementDate"`

	// A list of the items being acknowledged with associated details.
	Items []OrderAcknowledgementItem `json:"items"`

	// The purchase order number. Formatting Notes: 8-character alpha-numeric code.
	PurchaseOrderNumber string              `json:"purchaseOrderNumber"`
	SellingParty        PartyIdentification `json:"sellingParty"`
}

// Details of the item being acknowledged.
type OrderAcknowledgementItem struct {

	// Amazon Standard Identification Number (ASIN) of an item.
	AmazonProductIdentifier *string `json:"amazonProductIdentifier,omitempty"`

	// The discount multiplier that should be applied to the price if a vendor sells books with a list price. This is a multiplier factor to arrive at a final discounted price. A multiplier of .90 would be the factor if a 10% discount is given.
	DiscountMultiplier *string `json:"discountMultiplier,omitempty"`

	// This is used to indicate acknowledged quantity.
	ItemAcknowledgements []OrderItemAcknowledgement `json:"itemAcknowledgements"`

	// Line item sequence number for the item.
	ItemSequenceNumber *string `json:"itemSequenceNumber,omitempty"`

	// An amount of money, including units in the form of currency.
	ListPrice *Money `json:"listPrice,omitempty"`

	// An amount of money, including units in the form of currency.
	NetCost Money `json:"netCost"`

	// Details of quantity ordered.
	OrderedQuantity ItemQuantity `json:"orderedQuantity"`

	// The vendor selected product identification of the item. Should be the same as was sent in the purchase order.
	VendorProductIdentifier *string `json:"vendorProductIdentifier,omitempty"`
}

// Details of an order.
type OrderDetails struct {
	BillToParty *PartyIdentification `json:"billToParty,omitempty"`
	BuyingParty *PartyIdentification `json:"buyingParty,omitempty"`

	// If requested by the recipient, this field will contain a promotional/deal number. The discount code line is optional. It is used to obtain a price discount on items on the order.
	DealCode *string `json:"dealCode,omitempty"`

	// Defines a date time interval according to ISO8601. Interval is separated by double hyphen (--).
	DeliveryWindow *DateTimeInterval `json:"deliveryWindow,omitempty"`

	// Import details for an import order.
	ImportDetails *ImportDetails `json:"importDetails,omitempty"`

	// A list of items in this purchase order.
	Items []OrderItem `json:"items"`

	// Payment method used.
	PaymentMethod *OrderDetailsPaymentMethod `json:"paymentMethod,omitempty"`

	// The date when purchase order was last changed by Amazon after the order was placed. This date will be greater than 'purchaseOrderDate'. This means the PO data was changed on that date and vendors are required to fulfill the  updated PO. The PO changes can be related to Item Quantity, Ship to Location, Ship Window etc. This field will not be present in orders that have not changed after creation. Must be in ISO-8601 date/time format.
	PurchaseOrderChangedDate *time.Time `json:"purchaseOrderChangedDate,omitempty"`

	// The date the purchase order was placed. Must be in ISO-8601 date/time format.
	PurchaseOrderDate time.Time `json:"purchaseOrderDate"`

	// The date when current purchase order state was changed. Current purchase order state is available in the field 'purchaseOrderState'. Must be in ISO-8601 date/time format.
	PurchaseOrderStateChangedDate time.Time `json:"purchaseOrderStateChangedDate"`

	// Type of purchase order.
	PurchaseOrderType *OrderDetailsPurchaseOrderType `json:"purchaseOrderType,omitempty"`
	SellingParty      *PartyIdentification           `json:"sellingParty,omitempty"`
	ShipToParty       *PartyIdentification           `json:"shipToParty,omitempty"`

	// Defines a date time interval according to ISO8601. Interval is separated by double hyphen (--).
	ShipWindow *DateTimeInterval `json:"shipWindow,omitempty"`
}

// Payment method used.
type OrderDetailsPaymentMethod string

// Type of purchase order.
type OrderDetailsPurchaseOrderType string

// OrderItem defines model for OrderItem.
type OrderItem struct {

	// Amazon Standard Identification Number (ASIN) of an item.
	AmazonProductIdentifier *string `json:"amazonProductIdentifier,omitempty"`

	// When true, we will accept backorder confirmations for this item.
	IsBackOrderAllowed bool `json:"isBackOrderAllowed"`

	// Numbering of the item on the purchase order. The first item will be 1, the second 2, and so on.
	ItemSequenceNumber string `json:"itemSequenceNumber"`

	// An amount of money, including units in the form of currency.
	ListPrice *Money `json:"listPrice,omitempty"`

	// An amount of money, including units in the form of currency.
	NetCost *Money `json:"netCost,omitempty"`

	// Details of quantity ordered.
	OrderedQuantity ItemQuantity `json:"orderedQuantity"`

	// The vendor selected product identification of the item.
	VendorProductIdentifier *string `json:"vendorProductIdentifier,omitempty"`
}

// OrderItemAcknowledgement defines model for OrderItemAcknowledgement.
type OrderItemAcknowledgement struct {

	// Details of quantity ordered.
	AcknowledgedQuantity ItemQuantity `json:"acknowledgedQuantity"`

	// This indicates the acknowledgement code.
	AcknowledgementCode OrderItemAcknowledgementAcknowledgementCode `json:"acknowledgementCode"`

	// Indicates the reason for rejection.
	RejectionReason *OrderItemAcknowledgementRejectionReason `json:"rejectionReason,omitempty"`

	// Estimated delivery date per line item. Must be in ISO-8601 date/time format.
	ScheduledDeliveryDate *time.Time `json:"scheduledDeliveryDate,omitempty"`

	// Estimated ship date per line item. Must be in ISO-8601 date/time format.
	ScheduledShipDate *time.Time `json:"scheduledShipDate,omitempty"`
}

// This indicates the acknowledgement code.
type OrderItemAcknowledgementAcknowledgementCode string

// Indicates the reason for rejection.
type OrderItemAcknowledgementRejectionReason string

// OrderItemStatus defines model for OrderItemStatus.
type OrderItemStatus struct {

	// Acknowledgement status information.
	AcknowledgementStatus *struct {

		// Details of quantity ordered.
		AcceptedQuantity *ItemQuantity `json:"acceptedQuantity,omitempty"`

		// Details of item quantity confirmed.
		AcknowledgementStatusDetails *[]AcknowledgementStatusDetails `json:"acknowledgementStatusDetails,omitempty"`

		// Confirmation status of line item.
		ConfirmationStatus *OrderItemStatusAcknowledgementStatusConfirmationStatus `json:"confirmationStatus,omitempty"`

		// Details of quantity ordered.
		RejectedQuantity *ItemQuantity `json:"rejectedQuantity,omitempty"`
	} `json:"acknowledgementStatus,omitempty"`

	// Buyer's Standard Identification Number (ASIN) of an item.
	BuyerProductIdentifier *string `json:"buyerProductIdentifier,omitempty"`

	// Numbering of the item on the purchase order. The first item will be 1, the second 2, and so on.
	ItemSequenceNumber string `json:"itemSequenceNumber"`

	// An amount of money, including units in the form of currency.
	ListPrice *Money `json:"listPrice,omitempty"`

	// An amount of money, including units in the form of currency.
	NetCost *Money `json:"netCost,omitempty"`

	// Ordered quantity information.
	OrderedQuantity *struct {

		// Details of quantity ordered.
		OrderedQuantity *ItemQuantity `json:"orderedQuantity,omitempty"`

		// Details of item quantity ordered.
		OrderedQuantityDetails *[]OrderedQuantityDetails `json:"orderedQuantityDetails,omitempty"`
	} `json:"orderedQuantity,omitempty"`

	// The vendor selected product identification of the item.
	VendorProductIdentifier *string `json:"vendorProductIdentifier,omitempty"`
}

// Confirmation status of line item.
type OrderItemStatusAcknowledgementStatusConfirmationStatus string

// OrderList defines model for OrderList.
type OrderList struct {
	Orders     *[]Order    `json:"orders,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// OrderListStatus defines model for OrderListStatus.
type OrderListStatus struct {
	OrdersStatus *[]OrderStatus `json:"ordersStatus,omitempty"`
	Pagination   *Pagination    `json:"pagination,omitempty"`
}

// Current status of a purchase order.
type OrderStatus struct {

	// Detailed description of items order status.
	ItemStatus ItemStatus `json:"itemStatus"`

	// The date when the purchase order was last updated. Must be in ISO-8601 date/time format.
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`

	// The date the purchase order was placed. Must be in ISO-8601 date/time format.
	PurchaseOrderDate time.Time `json:"purchaseOrderDate"`

	// The buyer's purchase order number for this order. Formatting Notes: 8-character alpha-numeric code.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`

	// The status of the buyer's purchase order for this order.
	PurchaseOrderStatus OrderStatusPurchaseOrderStatus `json:"purchaseOrderStatus"`
	SellingParty        PartyIdentification            `json:"sellingParty"`
	ShipToParty         PartyIdentification            `json:"shipToParty"`
}

// The status of the buyer's purchase order for this order.
type OrderStatusPurchaseOrderStatus string

// Details of item quantity ordered
type OrderedQuantityDetails struct {

	// Details of quantity ordered.
	CancelledQuantity *ItemQuantity `json:"cancelledQuantity,omitempty"`

	// Details of quantity ordered.
	OrderedQuantity *ItemQuantity `json:"orderedQuantity,omitempty"`

	// The date when the line item quantity was updated by buyer. Must be in ISO-8601 date/time format.
	UpdatedDate *time.Time `json:"updatedDate,omitempty"`
}

// Pagination defines model for Pagination.
type Pagination struct {

	// A generated string used to pass information to your next request. If NextToken is returned, pass the value of NextToken to the next request. If NextToken is not returned, there are no more purchase order items to return.
	NextToken *string `json:"nextToken,omitempty"`
}

// PartyIdentification defines model for PartyIdentification.
type PartyIdentification struct {

	// Address of the party.
	Address *Address `json:"address,omitempty"`

	// Assigned identification for the party. For example, warehouse code or vendor code. Please refer to specific party for more details.
	PartyId string `json:"partyId"`

	// Tax registration details of the entity.
	TaxInfo *TaxRegistrationDetails `json:"taxInfo,omitempty"`
}

// The request schema for the submitAcknowledgment operation.
type SubmitAcknowledgementRequest struct {
	Acknowledgements *[]OrderAcknowledgement `json:"acknowledgements,omitempty"`
}

// The response schema for the submitAcknowledgement operation
type SubmitAcknowledgementResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList     `json:"errors,omitempty"`
	Payload *TransactionId `json:"payload,omitempty"`
}

// Tax registration details of the entity.
type TaxRegistrationDetails struct {

	// Tax registration number for the entity. For example, VAT ID.
	TaxRegistrationNumber string `json:"taxRegistrationNumber"`

	// Tax registration type for the entity.
	TaxRegistrationType TaxRegistrationDetailsTaxRegistrationType `json:"taxRegistrationType"`
}

// Tax registration type for the entity.
type TaxRegistrationDetailsTaxRegistrationType string

// TransactionId defines model for TransactionId.
type TransactionId struct {

	// GUID assigned by Amazon to identify this transaction. This value can be used with the Transaction Status API to return the status of this transaction.
	TransactionId *string `json:"transactionId,omitempty"`
}

// SubmitAcknowledgementJSONBody defines parameters for SubmitAcknowledgement.
type SubmitAcknowledgementJSONBody SubmitAcknowledgementRequest

// GetPurchaseOrdersParams defines parameters for GetPurchaseOrders.
type GetPurchaseOrdersParams struct {

	// The limit to the number of records returned. Default value is 100 records.
	Limit *int64 `json:"limit,omitempty"`

	// Purchase orders that became available after this time will be included in the result. Must be in ISO-8601 date/time format.
	CreatedAfter *time.Time `json:"createdAfter,omitempty"`

	// Purchase orders that became available before this time will be included in the result. Must be in ISO-8601 date/time format.
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`

	// Sort in ascending or descending order by purchase order creation date.
	SortOrder *GetPurchaseOrdersParamsSortOrder `json:"sortOrder,omitempty"`

	// Used for pagination when there is more purchase orders than the specified result size limit. The token value is returned in the previous API call
	NextToken *string `json:"nextToken,omitempty"`

	// When true, returns purchase orders with complete details. Otherwise, only purchase order numbers are returned. Default value is true.
	IncludeDetails *string `json:"includeDetails,omitempty"`

	// Purchase orders that changed after this timestamp will be included in the result. Must be in ISO-8601 date/time format.
	ChangedAfter *time.Time `json:"changedAfter,omitempty"`

	// Purchase orders that changed before this timestamp will be included in the result. Must be in ISO-8601 date/time format.
	ChangedBefore *time.Time `json:"changedBefore,omitempty"`

	// Current state of the purchase order item. If this value is Cancelled, this API will return purchase orders which have one or more items cancelled by Amazon with updated item quantity as zero.
	PoItemState *GetPurchaseOrdersParamsPoItemState `json:"poItemState,omitempty"`

	// When true, returns purchase orders which were modified after the order was placed. Vendors are required to pull the changed purchase order and fulfill the updated purchase order and not the original one. Default value is false.
	IsPOChanged *string `json:"isPOChanged,omitempty"`

	// Filters purchase orders based on the purchase order state.
	PurchaseOrderState *GetPurchaseOrdersParamsPurchaseOrderState `json:"purchaseOrderState,omitempty"`

	// Filters purchase orders based on the specified ordering vendor code. This value should be same as 'sellingParty.partyId' in the purchase order. If not included in the filter, all purchase orders for all of the vendor codes that exist in the vendor group used to authorize the API client application are returned.
	OrderingVendorCode *string `json:"orderingVendorCode,omitempty"`
}

// GetPurchaseOrdersParamsSortOrder defines parameters for GetPurchaseOrders.
type GetPurchaseOrdersParamsSortOrder string

// GetPurchaseOrdersParamsPoItemState defines parameters for GetPurchaseOrders.
type GetPurchaseOrdersParamsPoItemState string

// GetPurchaseOrdersParamsPurchaseOrderState defines parameters for GetPurchaseOrders.
type GetPurchaseOrdersParamsPurchaseOrderState string

// GetPurchaseOrdersStatusParams defines parameters for GetPurchaseOrdersStatus.
type GetPurchaseOrdersStatusParams struct {

	// The limit to the number of records returned. Default value is 100 records.
	Limit *int64 `json:"limit,omitempty"`

	// Sort in ascending or descending order by purchase order creation date.
	SortOrder *GetPurchaseOrdersStatusParamsSortOrder `json:"sortOrder,omitempty"`

	// Used for pagination when there are more purchase orders than the specified result size limit.
	NextToken *string `json:"nextToken,omitempty"`

	// Purchase orders that became available after this timestamp will be included in the result. Must be in ISO-8601 date/time format.
	CreatedAfter *time.Time `json:"createdAfter,omitempty"`

	// Purchase orders that became available before this timestamp will be included in the result. Must be in ISO-8601 date/time format.
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`

	// Purchase orders for which the last purchase order update happened after this timestamp will be included in the result. Must be in ISO-8601 date/time format.
	UpdatedAfter *time.Time `json:"updatedAfter,omitempty"`

	// Purchase orders for which the last purchase order update happened before this timestamp will be included in the result. Must be in ISO-8601 date/time format.
	UpdatedBefore *time.Time `json:"updatedBefore,omitempty"`

	// Provides purchase order status for the specified purchase order number.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`

	// Filters purchase orders based on the specified purchase order status. If not included in filter, this will return purchase orders for all statuses.
	PurchaseOrderStatus *GetPurchaseOrdersStatusParamsPurchaseOrderStatus `json:"purchaseOrderStatus,omitempty"`

	// Filters purchase orders based on the specified purchase order item status. If not included in filter, purchase orders for all statuses are included.
	ItemConfirmationStatus *GetPurchaseOrdersStatusParamsItemConfirmationStatus `json:"itemConfirmationStatus,omitempty"`

	// Filters purchase orders based on the specified ordering vendor code. This value should be same as 'sellingParty.partyId' in the purchase order. If not included in filter, all purchase orders for all the vendor codes that exist in the vendor group used to authorize API client application are returned.
	OrderingVendorCode *string `json:"orderingVendorCode,omitempty"`

	// Filters purchase orders for a specific buyer's Fulfillment Center/warehouse by providing ship to location id here. This value should be same as 'shipToParty.partyId' in the purchase order. If not included in filter, this will return purchase orders for all the buyer's warehouses used for vendor group purchase orders.
	ShipToPartyId *string `json:"shipToPartyId,omitempty"`
}

// GetPurchaseOrdersStatusParamsSortOrder defines parameters for GetPurchaseOrdersStatus.
type GetPurchaseOrdersStatusParamsSortOrder string

// GetPurchaseOrdersStatusParamsPurchaseOrderStatus defines parameters for GetPurchaseOrdersStatus.
type GetPurchaseOrdersStatusParamsPurchaseOrderStatus string

// GetPurchaseOrdersStatusParamsItemConfirmationStatus defines parameters for GetPurchaseOrdersStatus.
type GetPurchaseOrdersStatusParamsItemConfirmationStatus string

// SubmitAcknowledgementJSONRequestBody defines body for SubmitAcknowledgement for application/json ContentType.
type SubmitAcknowledgementJSONRequestBody SubmitAcknowledgementJSONBody