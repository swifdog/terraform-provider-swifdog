package tfresource

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
)

func ResourcePersistentVolume() *schema.Resource {
	return &schema.Resource{
		Description: "A volume can be used to persist data when mounted into packets.",
		Create:      resourcePersistentVolumeCreate,
		Read:        resourcePersistentVolumeRead,
		Update:      resourcePersistentVolumeUpdate,
		Delete:      resourcePersistentVolumeDelete,
		Schema: map[string]*schema.Schema{
			"projectid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"capacity": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
	}
}

func resourcePersistentVolumeCreate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	name := data.Get("name").(string)
	capacity := data.Get("capacity").(string)

	project, err := client.CreatePersistentVolume(data.Get("projectid").(string), &swifdog.CreateOrPatchPersistentVolumeRequest{
		Name:     name,
		Capacity: capacity,
	})
	if err != nil {
		return err
	}

	data.SetId(project.ID)
	return resourcePersistentVolumeRead(data, i)
}

func resourcePersistentVolumeRead(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	persistentVolume, err := client.GetPersistentVolume(data.Get("projectid").(string), data.Id())
	if err != nil {
		return err
	}

	_ = data.Set("name", persistentVolume.Name)
	_ = data.Set("capacity", persistentVolume.Capacity)
	data.SetId(persistentVolume.ID)
	return nil
}

func resourcePersistentVolumeUpdate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	name := data.Get("name").(string)
	capacity := data.Get("capacity").(string)

	project, err := client.PatchPersistentVolume(data.Get("projectid").(string), data.Id(), &swifdog.CreateOrPatchPersistentVolumeRequest{
		Name:     name,
		Capacity: capacity,
	})
	if err != nil {
		return err
	}

	data.SetId(project.ID)
	return resourcePersistentVolumeRead(data, i)
}

func resourcePersistentVolumeDelete(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	err := client.DeletePersistentVolumeById(data.Get("projectid").(string), data.Id())
	if err != nil {
		return err
	}

	return nil
}
