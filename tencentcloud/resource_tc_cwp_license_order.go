/*
Provides a resource to create a cwp license_order

Example Usage

```hcl
resource "tencentcloud_cwp_license_order" "example" {
  alias        = "tf_example"
  license_type = 0
  license_num  = 1
  region_id    = 1
  project_id   = 0
}
```

Import

cwp license_order can be imported using the id, e.g.

```
terraform import tencentcloud_cwp_license_order.license_order license_order_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cwp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cwp/v20180228"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudCwpLicenseOrder() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCwpLicenseOrderCreate,
		Read:   resourceTencentCloudCwpLicenseOrderRead,
		Update: resourceTencentCloudCwpLicenseOrderUpdate,
		Delete: resourceTencentCloudCwpLicenseOrderDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"alias": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Resource alias.",
			},
			"license_type": {
				Optional:     true,
				Type:         schema.TypeInt,
				Default:      LICENSE_TYPE_0,
				ValidateFunc: validateAllowedIntValue(LICENSE_TYPE),
				Description:  "LicenseType, 0 CWP Pro - Pay as you go, 1 CWP Pro - Monthly subscription, 2 CWP Ultimate - Monthly subscription. Default is 0.",
			},
			"license_num": {
				Optional:    true,
				Type:        schema.TypeInt,
				Default:     1,
				Description: "License quantity, Quantity to be purchased.Default is 1.",
			},
			"region_id": {
				Optional:     true,
				Type:         schema.TypeInt,
				Default:      REGION_ID_1,
				ValidateFunc: validateAllowedIntValue(REGION_ID),
				Description:  "Purchase order region, only 1 Guangzhou, 9 Singapore is supported here. Guangzhou is recommended. Singapore is whitelisted. Default is 1.",
			},
			"project_id": {
				Optional:    true,
				Type:        schema.TypeInt,
				Default:     0,
				Description: "Project ID. Default is 0.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tags of the license order.",
			},
		},
	}
}

func resourceTencentCloudCwpLicenseOrderCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cwp_license_order.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = cwp.NewCreateLicenseOrderRequest()
		response   = cwp.NewCreateLicenseOrderResponse()
		resourceId string
	)

	if v, ok := d.GetOkExists("license_type"); ok {
		request.LicenseType = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("license_num"); ok {
		request.LicenseNum = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("region_id"); ok {
		request.RegionId = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("project_id"); ok {
		request.ProjectId = helper.IntUint64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().CreateLicenseOrder(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || len(result.Response.ResourceIds) != 1 {
			e = fmt.Errorf("cwp licenseOrder not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cwp licenseOrder failed, reason:%+v", logId, err)
		return err
	}

	resourceId = *response.Response.ResourceIds[0]
	d.SetId(resourceId)

	//if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
	//	tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	//	region := meta.(*TencentCloudClient).apiV3Conn.Region
	//	resourceName := fmt.Sprintf("qcs::cwp::uin/:apiAppId/%s", region, apiAppId)
	//	if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
	//		return err
	//	}
	//}

	// set alias
	if v, ok := d.GetOk("alias"); ok {
		aliasRequest := cwp.NewModifyLicenseOrderRequest()
		aliasRequest.Alias = helper.String(v.(string))
		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().ModifyLicenseOrder(aliasRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s set cwp licenseOrder alias failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCwpLicenseOrderRead(d, meta)
}

func resourceTencentCloudCwpLicenseOrderRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cwp_license_order.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
		resourceId = d.Id()
	)

	licenseOrder, err := service.DescribeCwpLicenseOrderById(ctx, resourceId)
	if err != nil {
		return err
	}

	if licenseOrder == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CwpLicenseOrder` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if licenseOrder.Alias != nil {
		_ = d.Set("alias", licenseOrder.Alias)
	}

	if licenseOrder.LicenseType != nil {
		_ = d.Set("license_type", licenseOrder.LicenseType)
	}

	if licenseOrder.LicenseCnt != nil {
		_ = d.Set("license_num", licenseOrder.LicenseCnt)
	}

	if licenseOrder.ProjectId != nil {
		_ = d.Set("project_id", licenseOrder.ProjectId)
	}

	return nil
}

func resourceTencentCloudCwpLicenseOrderUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cwp_license_order.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = cwp.NewModifyLicenseOrderRequest()
		resourceId = d.Id()
	)

	immutableArgs := []string{"license_type", "region_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	request.ResourceId = &resourceId

	if d.HasChange("alias") {
		if v, ok := d.GetOk("alias"); ok {
			request.Alias = helper.String(v.(string))
		}
	}

	if d.HasChange("license_num") {
		if v, ok := d.GetOkExists("license_num"); ok {
			request.InquireNum = helper.IntUint64(v.(int))
		}
	}

	if d.HasChange("project_id") {
		if v, ok := d.GetOkExists("project_id"); ok {
			request.ProjectId = helper.IntUint64(v.(int))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().ModifyLicenseOrder(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update cwp licenseOrder failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCwpLicenseOrderRead(d, meta)
}

func resourceTencentCloudCwpLicenseOrderDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cwp_license_order.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		service     = CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
		resourceId  = d.Id()
		licenseType *uint64
	)

	if v, ok := d.GetOkExists("license_type"); ok {
		licenseType = helper.IntUint64(v.(int))
	}

	if err := service.DeleteCwpLicenseOrderById(ctx, resourceId, licenseType); err != nil {
		return err
	}

	return nil
}
