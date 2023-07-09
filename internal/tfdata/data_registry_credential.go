package tfdata

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
)

func DataRegistryCredential() *schema.Resource {
	return &schema.Resource{
		Read: dataRegistryCredentialRead,
		Schema: map[string]*schema.Schema{
			"projectid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"registryUrl": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: false,
			},
		},
	}
}

func dataRegistryCredentialRead(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	credentials, err := client.ListRegistryCredential(data.Get("projectid").(string))
	if err != nil {
		return err
	}

	for _, c := range credentials {
		if c.RegistryURL == data.Get("registryUrl").(string) {
			if data.Get("username") == nil || c.Username == data.Get("username").(string) {
				data.SetId(c.ID)
				_ = data.Set("registryUrl", c.RegistryURL)
				_ = data.Set("username", c.Username)
			}
		}
	}

	return errors.New("There is no registry credential with the requested values.")
}
