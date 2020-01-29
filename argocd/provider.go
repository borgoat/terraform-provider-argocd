package argocd

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/argoproj/argo-cd/pkg/apiclient"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_token": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"argocd_application": resourceApplication(),
		},
	}

	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return p
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	m := &Meta{}

	opts := apiclient.ClientOptions{}

	client, err := apiclient.NewClient(&opts)
	if err != nil {
		return nil, fmt.Errorf("could not create argo client: %w", err)
	}

	m.Argo = client

	return m, nil
}

type Meta struct {
	Argo apiclient.Client
}
