package tfdata

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
)

func DataPersistentVolume() *schema.Resource {
	return &schema.Resource{
		Read: dataPersistentVolumeRead,
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

func dataPersistentVolumeRead(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	persistentVolumes, err := client.ListPersistentVolume(data.Get("projectid").(string))
	if err != nil {
		return err
	}

	for _, pv := range persistentVolumes {
		if pv.Name == data.Get("name") {
			data.SetId(pv.ID)
			_ = data.Set("name", pv.Name)
			_ = data.Set("capacity", pv.Capacity)
			return nil
		}
	}

	return errors.New("There is no persistent volume with the given name")
}
