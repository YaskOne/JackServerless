package utils

var JKOrderPreparing = "order_preparing"
var JKOrderReady = "order_ready"
var JKBusinessOrderCanceled = "business_order_canceled"
var JKOrderRejected = "order_rejected"
var JKClientCanceledOrder = "client_canceled_order"

var NotificationTexts = map[string]string{
	JKOrderPreparing: "Votre commande est en preparation",
	JKOrderReady: "Votre commande est prete",
	JKBusinessOrderCanceled: "Votre commande a été annullée",
	JKOrderRejected: "Votre commande a été refusée",
	JKClientCanceledOrder: "Le client a annulé sa commande",
}
