package waf_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	tcacctest "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/acctest"
)

// go test -i; go test -test.run TestAccTencentCloudNeedFixWafClbIpAccessControlResource_basic -v
func TestAccTencentCloudNeedFixWafClbIpAccessControlResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			tcacctest.AccPreCheck(t)
		},
		Providers: tcacctest.AccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccWafClbIpAccessControl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "id"),
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "instance_id"),
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "domain"),
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "action_type"),
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "note"),
				),
			},
			{
				Config: testAccWafClbIpAccessControlUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "id"),
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "instance_id"),
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "domain"),
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "action_type"),
					resource.TestCheckResourceAttrSet("tencentcloud_waf_clb_ip_access_control.example", "note"),
				),
			},
			{
				ResourceName:      "tencentcloud_waf_clb_ip_access_control.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccWafClbIpAccessControl = `
resource "tencentcloud_waf_clb_ip_access_control" "example" {
  instance_id = "waf_2kxtlbky11bbcr4b"
  domain      = "example.com"
  action_type = 40
  note        = "note."

  ip_list = [
    "10.0.0.10",
    "172.0.0.16",
    "192.168.0.30"
  ]

  job_type = "TimedJob"

  job_date_time {
    time_t_zone = "UTC+8"

    timed {
      end_date_time   = 0
      start_date_time = 0
    }
  }
}
`

const testAccWafClbIpAccessControlUpdate = `
resource "tencentcloud_waf_clb_ip_access_control" "example" {
  instance_id = "waf_2kxtlbky11bbcr4b"
  domain      = "example.com"
  action_type = 40
  note        = "note update."

  ip_list = [
    "10.0.0.10",
    "172.0.0.16",
    "192.168.0.30",
	"168.10.10.10"
  ]

  job_type = "TimedJob"

  job_date_time {
    time_t_zone = "UTC+8"

    timed {
      end_date_time   = 0
      start_date_time = 0
    }
  }
}
`
