// Code generated by iacg; DO NOT EDIT.
package cvm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func DataSourceTencentCloudCvmImageSharePermission() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCvmImageSharePermissionRead,
		Schema: map[string]*schema.Schema{
			"image_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the image to be shared.",
			},

			"share_permission_set": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information on image sharing.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time when an image was shared.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the account with which the image is shared.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCvmImageSharePermissionRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("data_source.tencentcloud_cvm_image_share_permission.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(nil)
	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := CvmService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("image_id"); ok {
		paramMap["ImageId"] = helper.String(v.(string))
	}

	var respData []*cvm.SharePermission
	err := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCvmImageSharePermissionByFilter(ctx, paramMap)
		if e != nil {
			return tccommon.RetryError(e)
		}
		respData = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(respData))
	sharePermissionSetList := make([]map[string]interface{}, 0, len(respData))
	if respData != nil {
		for _, sharePermissionSet := range respData {
			sharePermissionSetMap := map[string]interface{}{}

			if sharePermissionSet.CreatedTime != nil {
				sharePermissionSetMap["created_time"] = sharePermissionSet.CreatedTime
			}

			var accountId string
			if sharePermissionSet.AccountId != nil {
				sharePermissionSetMap["account_id"] = sharePermissionSet.AccountId
				accountId = *sharePermissionSet.AccountId
			}

			ids = append(ids, accountId)
			sharePermissionSetList = append(sharePermissionSetList, sharePermissionSetMap)
		}

		_ = d.Set("share_permission_set", sharePermissionSetList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := tccommon.WriteToFile(output.(string), sharePermissionSetList); e != nil {
			return e
		}
	}

	return nil
}
