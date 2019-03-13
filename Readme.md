# Azurermext


### Motivation
During my everyday work my team and I had to execute custom SQL query on MySql database located on Azure.
Because the database required whitelisting of particular IP address to access it, we had to create firewall rule.
After the script was executed, whitelisted IP should have been removed from the list.
Azurerm provider does not allow resources removal.

I've decided to create custom provider for both enterprise use and for fun.

I know well that Terraform is a tool for infrastructure provisioning, not destroying it.
This custom provider was created mostly out of curiosity and necessity for single use case.

### Usage

This provider is mostly based on azurerm provider.
It uses service principal for authorization purposes.

First compile the provider by running:
```
go build .
```

Then copy it to example directory, and fill up the providers blocks.
```$xslt
cp terraform-provider-azurermext example/
cd example/
```

Now initialize providers in example directory and execute apply
```$xslt
terraform init
terraform apply
```

That's all, firewall rule should be created on your Azure MySQL database first using azurerm provider, and then deleted by azurermext provider.

 