package models

type ConnectLandedCost struct {
	Id int64 `json:"id"`
	Partner_key string `json:"partner_key"`
	Language string `json:"language"`
	Items string `json:"items"`
	Shipmentorigincountry string `json:"shipmentorigincountry"`
	Shipmentdestinationcountry string `json:"shipmentdestinationcountry"`
	Domesticshippingcost string `json:"domesticshippingcost"`
} 