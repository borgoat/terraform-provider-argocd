package argocd

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCreate,
		Read:   resourceApplicationRead,
		Update: resourceApplicationUpdate,
		Delete: resourceApplicationDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"finalizers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"source": &schema.Schema{
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repo_url": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"path": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceApplicationCreate(d *schema.ResourceData, m interface{}) error {
	return resourceApplicationRead(d, m)
}

func resourceApplicationRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceApplicationRead(d, m)
}

func resourceApplicationDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
