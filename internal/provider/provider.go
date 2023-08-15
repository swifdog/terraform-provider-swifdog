package provider

import (
	"errors"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/swifdog/go-swifdog/swifdog"
	"github.com/swifdog/terraform-provider-swifdog/internal/tfdata"
	_ "github.com/swifdog/terraform-provider-swifdog/internal/tfdata"
	"github.com/swifdog/terraform-provider-swifdog/internal/tfresource"
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
			"swifdog_floatingip":        tfresource.ResourceFloatingIP(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"swifdog_project":             tfdata.DataProject(),
			"swifdog_persistent_volume":   tfdata.DataPersistentVolume(),
			"swifdog_packet":              tfdata.DataPacket(),
			"swifdog_registry_credential": tfdata.DataRegistryCredential(),
		},
		ConfigureFunc: configure,
	}
}

func returnClient(client *swifdog.Client, err error) (interface{}, error) {
	apiEndpoint := os.Getenv("SWIFDOG_API_ENDPOINT")

	if apiEndpoint != "" {
		client.WithEndpoint(apiEndpoint)
	}

	return client, err
}

func configure(d *schema.ResourceData) (interface{}, error) {
	apiToken := d.Get("api_token")

	if apiToken != "" {
		client, err := swifdog.NewBearerTokenClient(apiToken.(string))
		return returnClient(client, err)
	}

	email := d.Get("email")
	password := d.Get("password")

	if email != "" || password != "" {
		client, err := swifdog.NewBasicClient(email.(string), password.(string))
		return returnClient(client, err)
	}

	return nil, errors.New("Please provide credentials like basic or token authentication.")
}
