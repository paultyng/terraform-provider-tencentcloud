package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	// go test -v ./tencentcloud -sweep=ap-guangzhou -sweep-run=tencentcloud_teo_zone
	resource.AddTestSweepers("tencentcloud_teo_zone", &resource.Sweeper{
		Name: "tencentcloud_teo_zone",
		F:    testSweepZone,
	})
}

func testSweepZone(region string) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cli, _ := sharedClientForRegion(region)
	client := cli.(*TencentCloudClient).apiV3Conn
	service := TeoService{client}

	zoneId := clusterPrometheusId

	zone, err := service.DescribeTeoZone(ctx, zoneId)
	if err != nil {
		return err
	}

	if zone == nil {
		return nil
	}

	err = service.DeleteTeoZoneById(ctx, zoneId)
	if err != nil {
		return err
	}

	return nil
}

// go test -i; go test -test.run TestAccTencentCloudTeoZone_basic -v
func TestAccTencentCloudTeoZone_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_PRIVATE) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckZoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTeoZone,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneExists("tencentcloud_teo_zone.basic"),
					resource.TestCheckResourceAttr("tencentcloud_teo_zone.basic", "zone_name", "tf-teo.xyz"),
					resource.TestCheckResourceAttr("tencentcloud_teo_zone.basic", "plan_type", "ent_with_bot"),
					resource.TestCheckResourceAttr("tencentcloud_teo_zone.basic", "type", "full"),
					resource.TestCheckResourceAttr("tencentcloud_teo_zone.basic", "paused", "false"),
					resource.TestCheckResourceAttr("tencentcloud_teo_zone.basic", "cname_speed_up", "enabled"),
					resource.TestCheckResourceAttr("tencentcloud_teo_zone.basic", "vanity_name_servers.#", "1"),
					resource.TestCheckResourceAttr("tencentcloud_teo_zone.basic", "vanity_name_servers.0.switch", "on"),
					//resource.TestCheckResourceAttr("tencentcloud_teo_zone.basic", "vanity_name_servers.0.servers", ""),
				),
			},
			//{
			//	ResourceName:      "tencentcloud_teo_zone.basic",
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//},
		},
	})
}

func testAccCheckZoneDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TeoService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloud_teo_zone" {
			continue
		}

		agents, err := service.DescribeTeoZone(ctx, rs.Primary.ID)
		if agents != nil {
			return fmt.Errorf("zone %s still exists", rs.Primary.ID)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func testAccCheckZoneExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := TeoService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		agents, err := service.DescribeTeoZone(ctx, rs.Primary.ID)
		if agents == nil {
			return fmt.Errorf("zone %s is not found", rs.Primary.ID)
		}
		if err != nil {
			return err
		}

		return nil
	}
}

const testAccTeoZoneVar = `
variable "zone_name" {
  default = "` + defaultZoneName + `"
}`

const testAccTeoZone = testAccTeoZoneVar + `

resource "tencentcloud_teo_zone" "basic" {
  zone_name      = var.zone_name
  plan_type      = "ent_with_bot"
  type           = "full"
  paused         = false
  cname_speed_up = "enabled"

  vanity_name_servers {
    switch  = "on"
    servers = ["ns5.tf-teo-test.com", "ns6.tencent-tf-teo-tset.com"]
  }

  tags = {
    "TestKey" = "TestValue"
  }
}

`
