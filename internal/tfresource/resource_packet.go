package tfresource

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
	"strconv"
)

func ResourcePacket() *schema.Resource {
	return &schema.Resource{
		Create: resourcePacketCreate,
		Read:   resourcePacketRead,
		Update: resourcePacketUpdate,
		Delete: resourcePacketDelete,
		Schema: map[string]*schema.Schema{
			"projectid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"image": {
				Type:     schema.TypeString,
				Required: true,
			},
			"registryCredentialId": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"env": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"volume": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"volumeid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"volumename": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"mountpath": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"internalport": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"containerport": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "tcp",
						},
					},
				},
			},
		},
	}
}

func parseRegistryCredentialId(data *schema.ResourceData) *string {
	value, ok := data.GetOk("registryCredentialId")

	if !ok {
		return nil
	}

	return value.(*string)
}

func resourcePacketCreate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	name := data.Get("name").(string)
	image := data.Get("image").(string)
	registryCredentialId := parseRegistryCredentialId(data)

	// envs
	envsRaw := data.Get("env").([]interface{})
	envs := make([]swifdog.EnvironmentVariable, 0)
	for _, raw := range envsRaw {
		d := raw.(map[string]interface{})

		envkey := d["key"].(string)
		envValue := d["value"].(string)

		envs = append(envs, swifdog.EnvironmentVariable{
			Key:   envkey,
			Value: envValue,
		})
	}
	// volumes
	volumesRaw := data.Get("volume").([]interface{})
	volumes := make([]swifdog.PersistentVolumeMount, 0)
	for _, raw := range volumesRaw {
		d := raw.(map[string]interface{})

		volumeId := d["volumeid"].(string)
		volumeName := d["volumename"].(string)
		mountPath := d["mountpath"].(string)

		volumes = append(volumes, swifdog.PersistentVolumeMount{
			VolumeId:   volumeId,
			VolumeName: volumeName,
			MountPath:  mountPath,
		})
	}

	internalPortsRaw := data.Get("internalport").([]interface{})
	internalPorts := make([]string, 0)

	for _, raw := range internalPortsRaw {
		d := raw.(map[string]interface{})

		containerPort := d["containerport"].(int)
		protocol := d["protocol"].(string)

		internalPorts = append(internalPorts, strconv.Itoa(containerPort)+"/"+protocol)
	}

	createOrPatchRequest := &swifdog.CreateOrPatchPacketRequest{
		Name:                 name,
		Image:                image,
		EnvironmentVariables: envs,
		VolumeMounts:         volumes,
		InternalPorts:        internalPorts,
	}

	if registryCredentialId != nil {
		createOrPatchRequest.RegistryCredentialId = *registryCredentialId
	}

	project, err := client.CreatePacket(data.Get("projectid").(string), createOrPatchRequest)
	if err != nil {
		return err
	}

	data.SetId(project.ID)
	return resourcePacketRead(data, i)
}

func resourcePacketRead(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	packet, err := client.GetPacket(data.Get("projectid").(string), data.Id())
	if err != nil {
		return err
	}

	_ = data.Set("name", packet.Name)
	_ = data.Set("image", packet.Image)
	_ = data.Set("registryCredentialId", packet.RegistryCredentialId)
	_ = data.Set("env", packet.EnvironmentVariables)
	_ = data.Set("volume", packet.VolumeMounts)
	_ = data.Set("internalport", packet.InternalPorts)
	data.SetId(packet.ID)
	return nil
}

func resourcePacketUpdate(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	name := data.Get("name").(string)
	image := data.Get("image").(string)
	registryCredentialId := parseRegistryCredentialId(data)

	// envs
	envsRaw := data.Get("env").([]interface{})
	envs := make([]swifdog.EnvironmentVariable, 0)
	for _, raw := range envsRaw {
		d := raw.(map[string]interface{})

		envkey := d["key"].(string)
		envValue := d["value"].(string)

		envs = append(envs, swifdog.EnvironmentVariable{
			Key:   envkey,
			Value: envValue,
		})
	}
	// volumes
	volumesRaw := data.Get("volume").([]interface{})
	volumes := make([]swifdog.PersistentVolumeMount, 0)
	for _, raw := range volumesRaw {
		d := raw.(map[string]interface{})

		volumeId := d["volumeid"].(string)
		volumeName := d["volumename"].(string)
		mountPath := d["mountpath"].(string)

		volumes = append(volumes, swifdog.PersistentVolumeMount{
			VolumeId:   volumeId,
			VolumeName: volumeName,
			MountPath:  mountPath,
		})
	}
	internalPortsRaw := data.Get("internalport").([]interface{})
	internalPorts := make([]string, 0)

	for _, raw := range internalPortsRaw {
		d := raw.(map[string]interface{})

		containerPort := d["containerport"].(int)
		protocol := d["protocol"].(string)

		internalPorts = append(internalPorts, strconv.Itoa(containerPort)+"/"+protocol)
	}

	createOrPatchRequest := &swifdog.CreateOrPatchPacketRequest{
		Name:                 name,
		Image:                image,
		EnvironmentVariables: envs,
		VolumeMounts:         volumes,
		InternalPorts:        internalPorts,
	}

	if registryCredentialId != nil {
		createOrPatchRequest.RegistryCredentialId = *registryCredentialId
	}

	project, err := client.PatchPacket(data.Get("projectid").(string), data.Id(), createOrPatchRequest)
	if err != nil {
		return err
	}

	data.SetId(project.ID)
	return resourcePacketRead(data, i)
}

func resourcePacketDelete(data *schema.ResourceData, i interface{}) error {
	client := i.(*swifdog.Client)

	err := client.DeletePacketById(data.Get("projectid").(string), data.Id())
	if err != nil {
		return err
	}

	return nil
}
