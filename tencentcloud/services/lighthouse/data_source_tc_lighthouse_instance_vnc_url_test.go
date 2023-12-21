package lighthouse_test

import (
	"testing"

	tcacctest "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudLighthouseInstanceVncUrlDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { tcacctest.AccPreCheckCommon(t, tcacctest.ACCOUNT_TYPE_PREPAY) },
		Providers: tcacctest.AccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccLighthouseInstanceVncUrlDataSource,
				Check:  resource.ComposeTestCheckFunc(tcacctest.AccCheckTencentCloudDataSourceID("data.tencentcloud_lighthouse_instance_vnc_url.instance_vnc_url")),
			},
		},
	})
}

const testAccLighthouseInstanceVncUrlDataSource = tcacctest.DefaultLighthoustVariables + `

data "tencentcloud_lighthouse_instance_vnc_url" "instance_vnc_url" {
  instance_id = var.lighthouse_id
}
`
