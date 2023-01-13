package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREEIPA_HOST", ""),
				Description: descriptions["host"],
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREEIPA_USERNAME", ""),
				Description: descriptions["username"],
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREEIPA_PASSWORD", ""),
				Description: descriptions["password"],
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["insecure"],
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"freeipa_dns_record":            resourceFreeIPADNSRecord(),
			"freeipa_dns_zone":              resourceFreeIPADNSZone(),
			"freeipa_user":                  resourceFreeIPAUser(),
			"freeipa_group":                 resourceFreeIPAGroup(),
			"freeipa_user_group_membership": resourceFreeIPAUserGroupMembership(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"freeipa_user":                  dataFreeIPAUser(),
			// "freeipa_group":                 dataFreeIPAGroup(),
			// "freeipa_user_group_membership": dataFreeIPAUserGroupMembership(),
		},

		ConfigureFunc: providerConfigure,
	}
	return provider
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"host": "The FreeIPA host",

		"username": "Username to use for connection",

		"password": "Password to use for connection",

		"insecure": "Whether to verify the server's SSL certificate",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &Config{
		Host:               d.Get("host").(string),
		Username:           d.Get("username").(string),
		Password:           d.Get("password").(string),
		InsecureSkipVerify: d.Get("insecure").(bool),
	}, nil
}
