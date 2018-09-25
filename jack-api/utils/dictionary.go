package utils


type JKString string

const (
	JKOrderAccepted    JKString = "Commande validée. Retrait disponible a %v"
	JKOrderPreparing   JKString = "Votre commande est en preparation, elle sera prête a %v"
	JKOrderReady    JKString = "Votre commande est prete"
	JKBusinessOrderCanceled   JKString = "Votre commande a été annullée"
	JKOrderRejected   JKString = "Votre commande a été refusée"
	JKClientCanceledOrder    JKString = "Le client a annulé sa commande"
)
