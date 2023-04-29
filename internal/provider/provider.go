package provider

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/terraform-provider-swifdog/internal/tfdata"
	_ "github.com/swifdog/terraform-provider-swifdog/internal/tfdata"
	"github.com/swifdog/terraform-provider-swifdog/internal/tfresource"
	"github.com/swifdog/go-swifdog/swifdog"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SWIFDOG_EMAIL", nil),
				Description: "The Email when using basic auth",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SWIFDOG_PASSWORD", nil),
				Description: "The Password when using basic auth",
			},
			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SWIFDOG_API_TOKEN", nil),
				Description: "The API token to access the swifdog API using bearer token auth",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"swifdog_project":           tfresource.ResourceProject(),
			"swifdog_persistent_volume": tfresource.ResourcePersistentVolume(),
			"swifdog_packet":            tfresource.ResourcePacket(),
			"swifdog_ingress_rule":      tfresource.ResourceIngressRule(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"swifdog_project":           tfdata.DataProject(),
			"swifdog_persistent_volume": tfdata.DataPersistentVolume(),
			"swifdog_packet":            tfdata.DataPacket(),
		},
		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	apiToken := d.Get("api_token")

	if apiToken != "" {
		return swifdog.NewBearerTokenClient(apiToken.(string))
	}

	email := d.Get("email")
	password := d.Get("password")

	if email != "" || password != "" {
		return swifdog.NewBasicClient(email.(string), password.(string))
	}

	return nil, errors.New("Credentials not found.")
}
