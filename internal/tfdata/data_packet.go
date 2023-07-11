package tfdata

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
)

func DataPacket() *schema.Resource {
	return &schema.Resource{
		Read: dataPacketRead,
		Schema: map[string]*schema.Schema{
			"projectid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataPacketRead(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	packets, err := client.ListPackets(data.Get("projectid").(string))
	if err != nil {
		return err
	}

	for _, p := range packets {
		if p.Name == data.Get("name") {
			data.SetId(p.ID)
			_ = data.Set("name", p.Name)
			_ = data.Set("image", p.Image)
			_ = data.Set("registryCredentialId", p.RegistryCredentialId)
			_ = data.Set("environmentVariables", p.EnvironmentVariables)
			_ = data.Set("mountedVolumes", p.VolumeMounts)
			_ = data.Set("internalPorts", p.InternalPorts)
			return nil
		}
	}

	return errors.New("There is no packet with the given name")
}
