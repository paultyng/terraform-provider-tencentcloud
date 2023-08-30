package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTencentCloudGaapSecurityPolices_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_PREPAY) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccDataSourceTencentCloudGaapSecurityPolicesBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloud_gaap_security_policies.foo"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_gaap_security_policies.foo", "proxy_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloud_gaap_security_policies.foo", "status"),
				),
			},
		},
	})
}

var TestAccDataSourceTencentCloudGaapSecurityPolicesBasic = fmt.Sprintf(`

data tencentcloud_gaap_security_policies "foo" {
  id = "%s"
}
`, defaultGaapSecurityPolicyId)
