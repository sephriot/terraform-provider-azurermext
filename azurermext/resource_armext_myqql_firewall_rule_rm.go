package azurermext

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceArmMySqlFirewallRuleRm() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmMySqlFirewallRuleRmUpdate,
		Read:   resourceArmMySqlFirewallRuleRmRead,
		Update: resourceArmMySqlFirewallRuleRmUpdate,
		Delete: resourceArmMySqlFirewallRuleRmDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"server_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// Because FirewallRule can be created during current terraform apply Update method will always be executed
func resourceArmMySqlFirewallRuleRmRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceArmMySqlFirewallRuleRmUpdate(d *schema.ResourceData, m interface{}) error {

	sess := m.(*AzureSession)
	client := mysql.NewFirewallRulesClient(sess.SubscriptionID)
	client.Authorizer = sess.Authorizer
	rg := d.Get("resource_group_name").(string)
	serverName := d.Get("server_name").(string)
	frName := d.Get("name").(string)

	resp, err := client.Get(context.Background(), rg, serverName, frName)
	if err != nil {
		if resourceNotFound(resp.Response) {
			return nil
		}
		return fmt.Errorf("Error issuing get request for MySQL Firewall Rule, %s", err)
	}
	if resp.ID == nil {
		return fmt.Errorf("Cannot read MySQL Firewall Rule, %s", err)
	}

	future, err := client.Delete(context.Background(), rg, serverName, frName)
	if err != nil {
		return fmt.Errorf("Error issuing delete request for MySQL Firewall Rule, %s", err)
	}

	err = future.WaitForCompletionRef(context.Background(), client.Client)
	if err != nil {
		return fmt.Errorf("Error waiting for delete request for MySQL Firewall Rule, %s", err)
	}

	return resourceArmMySqlFirewallRuleRmRead(d, m)
}

func resourceArmMySqlFirewallRuleRmDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}