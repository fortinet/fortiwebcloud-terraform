// Copyright 2021 Fortinet, Inc. All rights reserved.
// Author: Chengwei Hu (@godwindy)

// Description: Configure OpenApi Validation.

package fortiwebcloud

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOpenApiValidation() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwbCloudOpenApiValidationUpdate,
		Read:   resourceFwbCloudOpenApiValidationRead,
		Update: resourceFwbCloudOpenApiValidationUpdate,
		Delete: resourceFwbCloudOpenApiValidationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"validation_files": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceFwbCloudOpenApiValidationRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiwebCloudClient).CloudClient
	if c == nil {
		return fmt.Errorf("FortiWebCloud client did not initialize successfully!")
	}

	appName := d.Get("app_name").(string)
	appQuery := &AppQuery{AppName: appName}
	queryClient := NewAppQueryClient(c, appQuery)
	queryClient.Send()
	qApp, err := queryClient.ReadData()
	if qApp == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", appName)
		return err
	}
	app := qApp.(map[string]interface{})

	//Get Params from d
	ep_id := app["ep_id"].(string)
	d.Set("ep_id", ep_id)

	openapiValidationQuery := &OpenapiValidationQuery{EPId: ep_id}
	openapiValidation := NewOpenapiValidationQueryClient(c, openapiValidationQuery)
	openapiValidation.Send()
	qOpenapiValidation, err := openapiValidation.ReadData()
	if qOpenapiValidation == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", appName)
		return err
	}
	res := qOpenapiValidation.(map[string]interface{})
	var action string = ""
	result := res["result"].(map[string]interface{})
	configs := result["configs"].(map[string]interface{})
	_action := configs["action"].(string)
	action_list := [3]string{"alert", "alert_deny", "deny_no_log"}

	for _, x := range action_list {
		if x == _action {
			action = _action
			d.Set("action", action)
			break
		}
	}

	if action == "" {
		return fmt.Errorf("invalid action %s, only allow %v", _action, action_list)
	}

	d.SetId(appName)

	return nil
}

func resourceFwbCloudOpenApiValidationUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiwebCloudClient).CloudClient
	if c == nil {
		return fmt.Errorf("FortiWebCloud client did not initialize successfully!")
	}

	var action string = ""
	_action := d.Get("action").(string)
	action_list := [3]string{"alert", "alert_deny", "deny_no_log"}

	for _, x := range action_list {
		if x == _action {
			action = _action
			break
		}
	}

	if action == "" {
		return fmt.Errorf("invalid action %s, only allow %v", _action, action_list)
	}

	appName := d.Get("app_name").(string)

	appQuery := &AppQuery{AppName: appName}
	queryClient := NewAppQueryClient(c, appQuery)
	queryClient.Send()
	qApp, err := queryClient.ReadData()

	if qApp == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", appName)
		return err
	}
	app := qApp.(map[string]interface{})

	//Get Params from d
	ep_id := app["ep_id"].(string)

	var status bool = true
	if d.Get("enable").(bool) == false {
		status = false
	}

	validation_files := d.Get("validation_files").([]interface{})

	ofiles := []OFiles{}
	uploadFiles := make([]UploadFiles, len(validation_files))

	if validation_files != nil {
		for i, e := range validation_files {
			ofile := OFiles{OpenapiFile: filepath.Base(e.(string)), Seq: (i + 1)}
			ofiles = append(ofiles, ofile)
			uploadFiles[i].FileName = "file_" + strconv.Itoa(i+1)
			uploadFiles[i].FilePath = e.(string)
		}
	}

	openapiValidationCfg := SchemaFile{
		Status:                 status,
		Action:                 action,
		SchemaFiles:            ofiles,
	}

	var template_disable bool = false

	OpenapiValidationCreate := &OpenapiValidationCreate{
		EPId:                   ep_id,
		Template_status:        template_disable,
		OpenapiValidationCfg:   openapiValidationCfg,
	}

	openapiValidation, err := NewOpenapiValidationCreateClient(c, OpenapiValidationCreate, uploadFiles)
	if err != nil {
		log.Print(err)
		return err
	}
	err = openapiValidation.Send()
	ret, err := openapiValidation.ReadData()
	if err == nil {
		log.Print(fmt.Sprintf("ret: %v\n", ret))
		d.SetId(appName)
		return nil
	} else {
		log.Printf("[ERR] Setup Open Validation Error: %s", err)
		return err
	}
}

func resourceFwbCloudOpenApiValidationDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiwebCloudClient).CloudClient
	if c == nil {
		return fmt.Errorf("FortiWebCloud client did not initialize successfully!")
	}

	appName := d.Id()

	appQuery := &AppQuery{AppName: appName}
	queryClient := NewAppQueryClient(c, appQuery)
	queryClient.Send()
	qApp, err := queryClient.ReadData()

	if qApp == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", appName)
		return err
	}
	app := qApp.(map[string]interface{})

	//Get Params from d
	ep_id := app["ep_id"].(string)

	var status bool = false

	var action string = "alert"

	var template_disable bool = false
	
	openapiValidationCfg := SchemaFile{
		Status:                 status,
		Action:                 action,
		SchemaFiles:            []OFiles{},
	}

	OpenapiValidationCreate := &OpenapiValidationCreate{
		EPId:                   ep_id,
		Template_status:        template_disable,
		OpenapiValidationCfg:   openapiValidationCfg,
	}

	openapiValidation, err := NewOpenapiValidationCreateClient(c, OpenapiValidationCreate, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	err = openapiValidation.Send()
	ret, err := openapiValidation.ReadData()
	if err == nil {
		log.Print(fmt.Sprintf("ret: %v\n", ret))
		d.SetId("")
		return nil
	} else {
		log.Printf("[ERR] Delete Open Validation Error: %s", err)
		return err
	}
}
