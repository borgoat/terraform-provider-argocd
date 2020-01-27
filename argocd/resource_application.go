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
