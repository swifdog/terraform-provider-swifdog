package tfdata

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
)

func DataProject() *schema.Resource {
	return &schema.Resource{
		Read: dataProjectRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataProjectRead(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	prjs, err := client.ListProjects()
	if err != nil {
		return err
	}

	for _, prj := range prjs {
		if prj.Name == data.Get("name") {
			data.SetId(prj.ID)
			_ = data.Set("name", prj.Name)
			_ = data.Set("description", prj.Description)
			return nil
		}
	}

	return errors.New("There is no project with the given name")
}
