package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudSslDescribeHostTkeInstanceListDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSslDescribeHostTkeInstanceListDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.tencentcloud_ssl_describe_host_tke_instance_list.describe_host_tke_instance_list")),
			},
		},
	})
}

const testAccSslDescribeHostTkeInstanceListDataSource = `

data "tencentcloud_ssl_describe_host_tke_instance_list" "describe_host_tke_instance_list" {
  certificate_id = "8u8DII0l"
}
`
