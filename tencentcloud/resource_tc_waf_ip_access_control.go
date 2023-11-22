/*
Provides a resource to create a waf ip_access_control

Example Usage

```hcl
resource "tencentcloud_waf_ip_access_control" "example" {
  instance_id = "waf_2kxtlbky00b3b4qz"
  domain      = "www.demo.com"
  edition     = "sparta-waf"
  item {
    ip   = "1.1.1.1"
    note = "desc info."
    action = 40
    valid_ts = "2019571199"
  }

  item {
    ip   = "2.2.2.2"
    note = "desc info."
    action = 42
    valid_ts = "2019571199"
  }

  item {
    ip   = "3.3.3.3"
    note = "desc info."
    action = 40
    valid_ts = "1680570420"
  }
}
```

Import

waf ip_access_control can be imported using the id, e.g.

```
terraform import tencentcloud_waf_ip_access_control.example waf_2kxtlbky00b3b4qz#www.demo.com#sparta-waf
```
*/
package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	waf "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/waf/v20180125"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudWafIpAccessControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudWafIpAccessControlCreate,
		Read:   resourceTencentCloudWafIpAccessControlRead,
		Update: resourceTencentCloudWafIpAccessControlUpdate,
		Delete: resourceTencentCloudWafIpAccessControlDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Waf instance Id.",
			},
			"domain": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},
			"edition": {
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(EDITION_TYPE),
				Description:  "Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.",
			},
			"items": {
				Required:    true,
				Type:        schema.TypeSet,
				Description: "Ip parameter list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IP address.",
						},
						"note": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Note info.",
						},
						"action": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validateAllowedIntValue([]int{40, 42}),
							Description:  "Action value 40 is whitelist, 42 is blacklist.",
						},
						"valid_ts": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Effective date, with a second level timestamp value. For example, 1680570420 represents 2023-04-04 09:07:00; 2019571199 means permanently effective.",
						},
						"valid_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Valid status.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudWafIpAccessControlCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_waf_ip_access_control.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = waf.NewUpsertIpAccessControlRequest()
		response   = waf.NewUpsertIpAccessControlResponse()
		instanceId string
		domain     string
		edition    string
	)

	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
		instanceId = v.(string)
	}

	if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
		domain = v.(string)
	}

	if v, ok := d.GetOk("edition"); ok {
		request.Edition = helper.String(v.(string))
		edition = v.(string)
	}

	if v, ok := d.GetOk("items"); ok {
		for _, item := range v.(*schema.Set).List() {
			dMap := item.(map[string]interface{})
			tmpMap := make(map[string]interface{})
			if v, ok := dMap["ip"]; ok {
				tmpMap["ip"] = v.(string)
			}

			tmpMap["source"] = "custom"

			if v, ok := dMap["note"]; ok {
				tmpMap["note"] = v.(string)
			}

			if v, ok := dMap["action"]; ok {
				tmpMap["action"] = v.(int)
			}

			if v, ok := dMap["valid_ts"]; ok {
				tmpMap["valid_ts"] = v.(int)
			}

			tmpByte, _ := json.Marshal(tmpMap)
			tmpStr := string(tmpByte)
			request.Items = append(request.Items, &tmpStr)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseWafClient().UpsertIpAccessControl(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response.Ids == nil || len(result.Response.Ids) == 0 {
			e = fmt.Errorf("waf ipAccessControl not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create waf ipAccessControl failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{instanceId, domain, edition}, FILED_SP))

	return resourceTencentCloudWafIpAccessControlRead(d, meta)
}

func resourceTencentCloudWafIpAccessControlRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_waf_ip_access_control.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = WafService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	instanceId := idSplit[0]
	domain := idSplit[1]
	edition := idSplit[2]

	ipAccessControlList, err := service.DescribeWafIpAccessControlById(ctx, domain)
	if err != nil {
		return err
	}

	if ipAccessControlList == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `WafIpAccessControl` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("instance_id", instanceId)
	_ = d.Set("domain", domain)
	_ = d.Set("edition", edition)

	if ipAccessControlList != nil {
		tmpList := make([]map[string]interface{}, 0, len(ipAccessControlList))
		for _, item := range ipAccessControlList {
			ipAccessControlMap := map[string]interface{}{}

			if item.Ip != nil {
				ipAccessControlMap["ip"] = item.Ip
			}

			if item.Source != nil {
				ipAccessControlMap["source"] = item.Source
			}

			if item.Note != nil {
				ipAccessControlMap["note"] = item.Note
			}

			if item.ActionType != nil {
				ipAccessControlMap["action"] = item.ActionType
			}

			if item.ValidTs != nil {
				ipAccessControlMap["valid_ts"] = item.ValidTs
			}

			if item.ValidStatus != nil {
				ipAccessControlMap["valid_status"] = item.ValidStatus
			}

			if item.Id != nil {
				ipAccessControlMap["id"] = item.Id
			}

			tmpList = append(tmpList, ipAccessControlMap)
		}

		_ = d.Set("items", tmpList)
	}

	return nil
}

func resourceTencentCloudWafIpAccessControlUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_waf_ip_access_control.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = waf.NewUpsertIpAccessControlRequest()
	)

	immutableArgs := []string{"instance_id", "domain", "edition"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	instanceId := idSplit[0]
	domain := idSplit[1]
	edition := idSplit[2]

	if v, ok := d.GetOk("items"); ok {
		for _, item := range v.(*schema.Set).List() {
			dMap := item.(map[string]interface{})
			tmpMap := make(map[string]interface{})
			if v, ok := dMap["ip"]; ok {
				tmpMap["ip"] = v.(string)
			}

			tmpMap["source"] = "custom"

			if v, ok := dMap["note"]; ok {
				tmpMap["note"] = v.(string)
			}

			if v, ok := dMap["action"]; ok {
				tmpMap["action"] = v.(int)
			}

			if v, ok := dMap["valid_ts"]; ok {
				tmpMap["valid_ts"] = v.(int)
			}

			tmpByte, _ := json.Marshal(tmpMap)
			tmpStr := string(tmpByte)
			request.Items = append(request.Items, &tmpStr)
		}
	}

	request.InstanceId = &instanceId
	request.Domain = &domain
	request.Edition = &edition

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseWafClient().UpsertIpAccessControl(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update waf ipAccessControl failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudWafIpAccessControlRead(d, meta)
}

func resourceTencentCloudWafIpAccessControlDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_waf_ip_access_control.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = WafService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[1]

	if err := service.DeleteWafIpAccessControlById(ctx, domain); err != nil {
		return err
	}

	return nil
}
