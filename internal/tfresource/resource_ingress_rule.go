package tfresource

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
)

func ResourceIngressRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceIngressRuleCreate,
		Read:   resourceIngressRuleRead,
		Update: resourceIngressRuleUpdate,
		Delete: resourceIngressRuleDelete,
		Schema: map[string]*schema.Schema{
			"projectid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"automanagessl": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"path": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Required: true,
						},
						"packetid": {
							Type:     schema.TypeString,
							Required: true,
						},
						"containerport": {
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceIngressRuleCreate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	hostname := data.Get("hostname").(string)
	automanagessl := data.Get("automanagessl").(bool)
	pathsRaw := data.Get("path").([]interface{})

	paths := make([]swifdog.IngressRulePath, 0)
	for _, raw := range pathsRaw {
		d := raw.(map[string]interface{})

		path := d["path"].(string)
		packetid := d["packetid"].(string)
		containerport := d["containerport"].(int)

		paths = append(paths, swifdog.IngressRulePath{
			Path:          path,
			PacketId:      packetid,
			ContainerPort: containerport,
		})
	}

	project, err := client.CreateIngressRule(data.Get("projectid").(string), &swifdog.IngressRule{
		Hostname:      hostname,
		AutoManageSSL: automanagessl,
		PathRules:     paths,
	})
	if err != nil {
		return err
	}

	data.SetId(project.ID)
	return resourceIngressRuleRead(data, i)
}

func resourceIngressRuleRead(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	rule, err := client.GetIngressRule(data.Get("projectid").(string), data.Id())
	if err != nil {
		return err
	}

	_ = data.Set("hostname", rule.Hostname)
	_ = data.Set("path", rule.PathRules)
	data.SetId(rule.ID)
	return nil
}

func resourceIngressRuleUpdate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	hostname := data.Get("hostname").(string)
	automanagessl := data.Get("automanagessl").(bool)
	pathsRaw := data.Get("path").([]interface{})

	paths := make([]swifdog.IngressRulePath, 0)
	for _, raw := range pathsRaw {
		d := raw.(map[string]interface{})

		path := d["path"].(string)
		packetid := d["packetid"].(string)
		containerport := d["containerport"].(int)

		paths = append(paths, swifdog.IngressRulePath{
			Path:          path,
			PacketId:      packetid,
			ContainerPort: containerport,
		})
	}

	project, err := client.PatchIngressRule(data.Get("projectid").(string), data.Id(), &swifdog.IngressRule{
		Hostname:      hostname,
		AutoManageSSL: automanagessl,
		PathRules:     paths,
	})
	if err != nil {
		return err
	}

	data.SetId(project.ID)
	return resourceIngressRuleRead(data, i)
}

func resourceIngressRuleDelete(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	err := client.DeleteIngressRuleById(data.Get("projectid").(string), data.Id())
	if err != nil {
		return err
	}

	return nil
}
