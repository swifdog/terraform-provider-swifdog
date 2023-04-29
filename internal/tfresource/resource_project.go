package tfresource

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
)

func ResourceProject() *schema.Resource {
	return &schema.Resource{
		Description: "A project can be used as a space for sub resources.",
		Create:      resourceProjectCreate,
		Read:        resourceProjectRead,
		Update:      resourceProjectUpdate,
		Delete:      resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
			},
		},
	}
}

func resourceProjectCreate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	name := data.Get("name").(string)
	description := data.Get("description").(string)

	project, err := client.CreateProject(&swifdog.CreateOrPatchProjectRequest{
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}

	data.SetId(project.ID)
	return resourceProjectRead(data, i)
}

func resourceProjectRead(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	project, err := client.GetProject(data.Id())
	if err != nil {
		return err
	}

	_ = data.Set("name", project.Name)
	_ = data.Set("description", project.Description)
	data.SetId(project.ID)
	return nil
}

func resourceProjectUpdate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	name := data.Get("name").(string)
	description := data.Get("description").(string)

	project, err := client.PatchProject(data.Id(), &swifdog.CreateOrPatchProjectRequest{
		Name:        name,
		Description: description,
	})
	if err != nil {
		return err
	}

	data.SetId(project.ID)
	return resourceProjectRead(data, i)
}

func resourceProjectDelete(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	err := client.DeleteProjectById(data.Id())
	if err != nil {
		return err
	}

	return nil
}
