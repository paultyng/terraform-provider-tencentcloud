package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudDnspodModifyRecordGroupOperationResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_PREPAY) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspodModifyRecordGroupOperation,
				Check:  resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloud_dnspod_modify_record_group_operation.modify_record_group", "domain", "iac-tf.cloud"),
					resource.TestCheckResourceAttr("tencentcloud_dnspod_modify_record_group_operation.modify_record_group", "group_id", 1),
					resource.TestCheckResourceAttr("tencentcloud_dnspod_modify_record_group_operation.modify_record_group", "record_id", "234"),
					resource.TestCheckResourceAttr("tencentcloud_dnspod_modify_record_group_operation.modify_record_group", "domain_id", 123),
				),
			},
		},
	})
}

const testAccDnspodModifyRecordGroupOperation = `

resource "tencentcloud_dnspod_modify_record_group_operation" "modify_record_group" {
  domain = "dnspod.cn"
  group_id = 1
  record_id = "234|345"
  domain_id = 123
}

`
