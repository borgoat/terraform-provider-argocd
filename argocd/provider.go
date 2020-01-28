package argocd

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"argocd_application": resourceApplication(),
		},

		ConfigureFunc: configureFunc,
	}
}

func configureFunc(d *schema.ResourceData) (interface{}, error) {
	return nil, nil
}
