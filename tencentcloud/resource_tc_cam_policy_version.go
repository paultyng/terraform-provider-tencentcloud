/*
Provides a resource to create a cam policy_version

Example Usage

```hcl
resource "tencentcloud_cam_policy_version" "policy_version" {
  policy_id = 171173780
  policy_document = jsonencode({
    "version": "2.0",
    "statement": [
      {
        "effect": "allow",
        "action": [
          "sts:AssumeRole"
        ],
        "resource": [
          "*"
        ]
      },
      {
        "effect": "allow",
        "action": [
          "cos:PutObject"
        ],
        "resource": [
          "*"
        ]
      },
      {
        "effect": "deny",
        "action": [
          "aa:*"
        ],
        "resource": [
          "*"
        ]
      },
      {
        "effect": "deny",
        "action": [
          "aa:*"
        ],
        "resource": [
          "*"
        ]
      }
    ]
  })
  set_as_default = "false"
}
```

Import

cam policy_version can be imported using the id, e.g.

```
terraform import tencentcloud_cam_policy_version.policy_version policy_version_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cam "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam/v20190116"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudCamPolicyVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamPolicyVersionCreate,
		Read:   resourceTencentCloudCamPolicyVersionRead,
		Update: resourceTencentCloudCamPolicyVersionUpdate,
		Delete: resourceTencentCloudCamPolicyVersionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policy_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "Strategy ID.",
			},

			"policy_document": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Strategic text information.",
			},

			"set_as_default": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeBool,
				Description: "Whether to set as a version of the current strategy.",
			},

			"policy_version": {
				Computed:    true,
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Strategic version detailsNote: This field may return NULL, indicating that the valid value cannot be obtained.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"version_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Strategic version numberNote: This field may return NULL, indicating that the valid value cannot be obtained.",
						},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Strategic version creation timeNote: This field may return NULL, indicating that the valid value cannot be obtained.",
						},
						"is_default_version": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether it is an effective version.0 means not, 1 means yesNote: This field may return NULL, indicating that the valid value cannot be obtained.",
						},
						"document": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Strategic grammar textNote: This field may return NULL, indicating that the valid value cannot be obtained.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudCamPolicyVersionCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cam_policy_version.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request   = cam.NewCreatePolicyVersionRequest()
		response  = cam.NewCreatePolicyVersionResponse()
		policyId  string
		versionId string
	)
	if v, ok := d.GetOkExists("policy_id"); ok {
		policyId = helper.IntToStr(v.(int))
		request.PolicyId = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("policy_document"); ok {
		request.PolicyDocument = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("set_as_default"); ok {
		request.SetAsDefault = helper.Bool(v.(bool))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().CreatePolicyVersion(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cam policyVersion failed, reason:%+v", logId, err)
		return err
	}
	if response == nil || response.Response == nil || response.Response.VersionId == nil {
		return fmt.Errorf("CAM policy version is null")
	}
	versionId = helper.UInt64ToStr(*response.Response.VersionId)
	d.SetId(policyId + FILED_SP + versionId)

	return resourceTencentCloudCamPolicyVersionRead(d, meta)
}

func resourceTencentCloudCamPolicyVersionRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cam_policy_version.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	policyId := idSplit[0]
	versionId := idSplit[1]

	policyVersion, err := service.DescribeCamPolicyVersionById(ctx, helper.StrToUInt64(policyId), helper.StrToUInt64(versionId))
	if err != nil {
		return err
	}

	if policyVersion == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CamPolicyVersion` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("policy_id", helper.StrToInt64(policyId))

	if policyVersion != nil {
		policyVersionMap := map[string]interface{}{}

		if policyVersion.VersionId != nil {
			policyVersionMap["version_id"] = policyVersion.VersionId
		}

		if policyVersion.CreateDate != nil {
			policyVersionMap["create_date"] = policyVersion.CreateDate
		}

		if policyVersion.IsDefaultVersion != nil {
			policyVersionMap["is_default_version"] = policyVersion.IsDefaultVersion
		}

		if policyVersion.Document != nil {
			policyVersionMap["document"] = policyVersion.Document
		}

		_ = d.Set("policy_version", []interface{}{policyVersionMap})

	}

	return nil
}

func resourceTencentCloudCamPolicyVersionDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cam_policy_version.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CamService{client: meta.(*TencentCloudClient).apiV3Conn}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	policyId := idSplit[0]
	versionId := idSplit[1]

	if err := service.DeleteCamPolicyVersionById(ctx, helper.StrToUInt64(policyId), helper.StrToUInt64(versionId)); err != nil {
		return err
	}

	return nil
}
func resourceTencentCloudCamPolicyVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_cam_policy_version.update")()
	defer inconsistentCheck(d, meta)()

	return resourceTencentCloudCamPolicyVersionRead(d, meta)
}
