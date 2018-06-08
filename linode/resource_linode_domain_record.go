package linode

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLinodeDomainRecord() *schema.Resource {
	return &schema.Resource{
		Create: createLinodeDomainRecord,
		Read:   readLinodeDomainRecord,
		Update: updateLinodeDomainRecord,
		Delete: deleteLinodeDomainRecord,
		Schema: map[string]*schema.Schema{
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"portocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"ttl_sec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func fillDomainRecord(d *schema.ResourceData) *DomainRecord {
	res := &DomainRecord{}

	if value, ok := d.GetOk("weight"); ok {
		weight := value.(int)
		res.Weight = &weight
	}

	if value, ok := d.GetOk("name"); ok {
		name := value.(string)
		res.Name = &name
	}

	if value, ok := d.GetOk("target"); ok {
		target := value.(string)
		res.Target = &target
	}

	if value, ok := d.GetOk("priority"); ok {
		priority := value.(int)
		res.Priority = &priority
	}

	res.Type = d.Get("type").(string)

	if value, ok := d.GetOk("service"); ok {
		service := value.(string)
		res.Service = &service
	}

	if value, ok := d.GetOk("protocol"); ok {
		protocol := value.(string)
		res.Protocol = &protocol
	}

	if value, ok := d.GetOk("ttl_sec"); ok {
		ttlSec := value.(int)
		res.TTLSec = &ttlSec
	}

	return res
}

func fillResourceData(r *DomainRecord, d *schema.ResourceData) {
	d.SetId(fmt.Sprintf("%d", *r.ID))

	if r.Weight != nil {
		d.Set("weight", *r.Weight)
	}

	if r.Name != nil {
		d.Set("name", *r.Name)
	}

	if r.Priority != nil {
		d.Set("target", *r.Target)
	}

	if r.Priority != nil {
		d.Set("priority", *r.Priority)
	}

	if r.Port != nil {
		d.Set("port", *r.Port)
	}

	if r.Service != nil {
		d.Set("service", *r.Service)
	}

	if r.Protocol != nil {
		d.Set("protocol", *r.Protocol)
	}

	if r.TTLSec != nil {
		d.Set("ttl_sec", *r.TTLSec)
	}
}

func createLinodeDomainRecord(d *schema.ResourceData, meta interface{}) error {
	client := meta.(LinodeClient)

	domainID := d.Get("domain_id").(string)

	domainRecord := fillDomainRecord(d)

	res := &DomainRecord{}

	// https://developers.linode.com/api/v4#operation/createDomainRecord
	if err := client.Request("POST", fmt.Sprintf("domains/%s/records", domainID), domainRecord, res); err != nil {
		return err
	}

	fillResourceData(res, d)

	return nil
}

func readLinodeDomainRecord(d *schema.ResourceData, meta interface{}) error {
	client := meta.(LinodeClient)

	domainID := d.Get("domain_id").(string)

	recordID := d.Id()

	res := &DomainRecord{}

	// https://developers.linode.com/api/v4#operation/createDomainRecord
	if err := client.Request("GET", fmt.Sprintf("domains/%s/records/%s", domainID, recordID), nil, res); err != nil {
		return err
	}

	fillResourceData(res, d)

	return nil
}

func updateLinodeDomainRecord(d *schema.ResourceData, meta interface{}) error {
	return errors.New("Not Implemented")
}

func deleteLinodeDomainRecord(d *schema.ResourceData, meta interface{}) error {
	return errors.New("Not Implemented")
}
