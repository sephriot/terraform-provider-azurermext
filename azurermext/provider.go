package azurermext

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/pkg/errors"
	"os"
)

const(
	SubscriptionID          = "AZURE_SUBSCRIPTION_ID"
	TenantID                = "AZURE_TENANT_ID"
	ClientID                = "AZURE_CLIENT_ID"
	ClientSecret            = "AZURE_CLIENT_SECRET"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"subscription_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(SubscriptionID, ""),
			},

			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(ClientID, ""),
			},

			"tenant_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(TenantID, ""),
			},
			// Client Secret specific fields
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(ClientSecret, ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"azurermext_mysql_firewall_rule_rm": resourceArmMySqlFirewallRuleRm(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {

		os.Setenv(ClientID, d.Get("client_id").(string))
		os.Setenv(ClientSecret, d.Get("client_secret").(string))
		os.Setenv(TenantID, d.Get("tenant_id").(string))
		authorizer, err := auth.NewAuthorizerFromEnvironment()

		if err != nil {
			return nil, errors.Wrap(err, "Can't initialize authorizer")
		}

		sess := AzureSession{
			SubscriptionID: d.Get("subscription_id").(string),
			Authorizer:     authorizer,
		}

		return &sess, nil
	}
}