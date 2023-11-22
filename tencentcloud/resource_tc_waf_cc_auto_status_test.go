package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudWafCcAutoStatusResource_basic -v
func TestAccTencentCloudWafCcAutoStatusResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccWafCcAutoStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloud_waf_cc_auto_status.example", "id"),
					resource.TestCheckResourceAttr("tencentcloud_waf_cc_auto_status.example", "domain", "keep.qcloudwaf.com"),
					resource.TestCheckResourceAttr("tencentcloud_waf_cc_auto_status.example", "edition", "sparta-waf"),
				),
			},
		},
	})
}

const testAccWafCcAutoStatus = `
resource "tencentcloud_waf_cc_auto_status" "example" {
  domain  = "keep.qcloudwaf.com"
  edition = "sparta-waf"
}
`
