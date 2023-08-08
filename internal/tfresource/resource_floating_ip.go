package tfresource

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
)

func ResourceFloatingIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceFloatingIPCreate,
		Read:   resourceFloatingIPRead,
		Update: resourceFloatingIPUpdate,
		Delete: resourceFloatingIPDelete,
		Schema: map[string]*schema.Schema{
			"projectid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version": { // only supporting IPv4 and IPv6
				Type:     schema.TypeInt,
				Optional: true,
			},
			"endpoint": { // a bundle of services to be exposed on this Floating IP.
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packetid": { // the packet id of the given container application
							Type:     schema.TypeString,
							Required: true,
						},
						"protocol": { // the packet id of the given container application
							Type:     schema.TypeString,
							Required: true,
							Default:  "TCP",
						},
						"containerport": { // the port on which the service is listening internally
							Type:     schema.TypeInt,
							Required: true,
						},
						"targetport": { // the port on which the service should be listening publicly
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceFloatingIPCreate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	version := data.Get("version").(int)
	endpointsRaw := data.Get("endpoint").([]interface{})

	endpoints := make([]swifdog.FloatingIPEndpoint, 0)
	for _, raw := range endpointsRaw {
		d := raw.(map[string]interface{})

		packetId := d["packetid"].(string)
		protocol := d["protocol"].(string)
		containerPort := d["containerport"].(int)
		targetPort := d["targetport"].(int)

		endpoints = append(endpoints, swifdog.FloatingIPEndpoint{
			PacketId:      packetId,
			Protocol:      protocol,
			ContainerPort: containerPort,
			TargetPort:    targetPort,
		})
	}

	floatingIp, err := client.CreateFloatingIP(data.Get("projectid").(string), &swifdog.FloatingIP{
		Version:   version,
		Endpoints: endpoints,
	})
	if err != nil {
		return err
	}

	data.SetId(floatingIp.ID)
	return resourceFloatingIPRead(data, i)
}

func resourceFloatingIPRead(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	floatingIP, err := client.GetFloatingIP(data.Get("projectid").(string), data.Id())
	if err != nil {
		return err
	}

	_ = data.Set("ip", floatingIP.IP)
	_ = data.Set("version", floatingIP.Version)
	_ = data.Set("endpoints", floatingIP.Endpoints)
	data.SetId(floatingIP.ID)
	return nil
}

func resourceFloatingIPUpdate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	version := data.Get("version").(int)
	endpointsRaw := data.Get("endpoint").([]interface{})

	endpoints := make([]swifdog.FloatingIPEndpoint, 0)
	for _, raw := range endpointsRaw {
		d := raw.(map[string]interface{})

		packetId := d["packetid"].(string)
		protocol := d["protocol"].(string)
		containerPort := d["containerport"].(int)
		targetPort := d["targetport"].(int)

		endpoints = append(endpoints, swifdog.FloatingIPEndpoint{
			PacketId:      packetId,
			Protocol:      protocol,
			ContainerPort: containerPort,
			TargetPort:    targetPort,
		})
	}

	floatingIp, err := client.PatchFloatingIP(data.Get("projectid").(string), data.Id(), &swifdog.FloatingIP{
		Version:   version,
		Endpoints: endpoints,
	})
	if err != nil {
		return err
	}

	data.SetId(floatingIp.ID)
	return resourceFloatingIPRead(data, i)
}

func resourceFloatingIPDelete(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	err := client.DeleteFloatingIPById(data.Get("projectid").(string), data.Id())
	if err != nil {
		return err
	}

	return nil
}
