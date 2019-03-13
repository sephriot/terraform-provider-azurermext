package azurermext

import "github.com/Azure/go-autorest/autorest"

// AzureSession is an object representing session for subscription
type AzureSession struct {
	SubscriptionID string
	Authorizer     autorest.Authorizer
}