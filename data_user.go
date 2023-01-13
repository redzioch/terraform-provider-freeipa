package main

import (
	"context"
	"log"
	// "strings"
	"fmt"
	// "tools"

	ipa "github.com/RomanButsiy/go-freeipa/freeipa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataFreeIPAUser() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataFreeIPAUserRead,

		Schema: map[string]*schema.Schema{
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"initials": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"home_directory": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gecos": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"login_shell": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"krb_principal_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"krb_principal_expiration": {
				Description: "Kerberos principal expiration " +
					"[RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format " +
					"(see [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) e.g., " +
					"`YYYY-MM-DDTHH:MM:SSZ`)",
				Type:     schema.TypeString,
				Optional: true,
			},
			"krb_password_expiration": {
				Description: "User password expiration " +
					"[RFC3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.8) format " +
					"(see [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) e.g., " +
					"`YYYY-MM-DDTHH:MM:SSZ`)",
				Type:     schema.TypeString,
				Optional: true,
			},
			// "userpassword": {
			// 	Type:      schema.TypeString,
			// 	Optional:  true,
			// 	Sensitive: true,
			// },
			"email_address": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"telephone_numbers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mobile_numbers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			// "random_password": {
			// 	Type:     schema.TypeBool,
			// 	Optional: true,
			// },
			"uid_number": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"gid_number": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"street_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"city": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"province": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"postal_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"organisation_unit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"job_title": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"manager": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"employee_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"employee_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"preferred_language": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ssh_public_key": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"car_license": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataFreeIPAUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Read freeipa user")

	client, err := meta.(*Config).Client()
	if err != nil {
		return diag.Errorf("Error creating freeipa identity client: %s", err)
	}

	all := true
	optArgs := ipa.UserShowOptionalArgs{
		All: &all,
	}

	if _v, ok := d.GetOkExists("name"); ok {
		v := _v.(string)
		optArgs.UID = &v
	}
	res, err := client.UserShow(&ipa.UserShowArgs{}, &optArgs)

	if err != nil {
		return diag.Errorf("Error show freeipa user: %s", err)
	}

	if err := d.Set("first_name", res.Result.Givenname); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("last_name", res.Result.Sn); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("name", res.Result.UID); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("name", res.Result.UID); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("full_name", res.Result.Cn); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("display_name", res.Result.Displayname); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("initials", res.Result.Initials); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("home_directory", res.Result.Homedirectory); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("gecos", res.Result.Gecos); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("login_shell", res.Result.Loginshell); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	// if err := d.Set("krb_principal_name", []schema.TypeString{res.Result.Krbprincipalname}); err != nil {
	// 	return diag.FromErr(fmt.Errorf("%s", err.Error()))
	// }

	// if err := d.Set("krb_principal_expiration", res.Result.Krbprincipalexpiration); err != nil {
	// 	return diag.FromErr(fmt.Errorf("%s", err.Error()))
	// }

	// if err := d.Set("krb_password_expiration", res.Result.Krbpasswordexpiration); err != nil {
	// 	return diag.FromErr(fmt.Errorf("%s", err.Error()))
	// }

	// if err := d.Set("email_address", res.Result.Mail); err != nil {
	// 	return diag.FromErr(fmt.Errorf("%s", err.Error()))
	// }

	// var telephone_numbers []string
	// telephone_numbers = append([]string{}, *res.Result.Telephonenumber...)
	// log.Printf(fmt.Sprintf("[DEBUG] telephone_numbers %s", strings.Join(telephone_numbers, ", ")))
	// if err := d.Set("telephone_numbers", *res.Result.Telephonenumber); err != nil {
	// 	return diag.FromErr(fmt.Errorf("%s", err.Error()))
	// }
	// var res_telephone_numbers []string
	res_mobile_numbers := []string{}
	if res.Result.Mobile != nil {
		res_mobile_numbers = *res.Result.Mobile
	}
	if err := d.Set("mobile_numbers", res_mobile_numbers); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	res_telephone_numbers := []string{}
	if res.Result.Telephonenumber != nil {
		res_telephone_numbers = *res.Result.Telephonenumber
	}
	if err := d.Set("telephone_numbers", res_telephone_numbers); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("uid_number", res.Result.Uidnumber); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("gid_number", res.Result.Gidnumber); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("street_address", res.Result.Street); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("city", res.Result.L); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("province", res.Result.St); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("postal_code", res.Result.Postalcode); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("organisation_unit", res.Result.Ou); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("job_title", res.Result.Title); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("manager", res.Result.Manager); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("employee_number", res.Result.Employeenumber); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("employee_type", res.Result.Employeetype); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("preferred_language", res.Result.Preferredlanguage); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	if err := d.Set("account_disabled", res.Result.Nsaccountlock); err != nil {
		return diag.FromErr(fmt.Errorf("%s", err.Error()))
	}

	// if err := d.Set("ssh_public_key", res.Result.Initials); err != nil {
	// 	return diag.FromErr(fmt.Errorf("%s", err.Error()))
	// }

	// if err := d.Set("preferred_language", res.Result.Initials); err != nil {
	// 	return diag.FromErr(fmt.Errorf("%s", err.Error()))
	// }

	d.SetId(res.Result.UID)

	log.Printf("[DEBUG] Read freeipa user %s", res.Result.UID)

	return nil
	// return dataFreeIPAUserRead(ctx, d, meta)
}
