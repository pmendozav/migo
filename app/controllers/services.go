package controllers

var (
	cuser *CUser = &CUser{}
	clogin *CLogin = &CLogin{}
	cpartner *CPartner = &CPartner{}
	cconnectlandedcost *CConnectLandedCost = &CConnectLandedCost{}
)

func RegisterUsers() {
	clogin.RegisterServices()
	cuser.RegisterServices()
	cpartner.RegisterServices()
	cconnectlandedcost.RegisterServices()
}