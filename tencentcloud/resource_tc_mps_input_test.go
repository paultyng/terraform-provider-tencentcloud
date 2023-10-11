package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudMpsInputResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMpsInput,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloud_mps_input.input", "id")),
			},
			{
				ResourceName:      "tencentcloud_mps_input.input",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccMpsInput = `

resource "tencentcloud_mps_input" "input" {
  flow_id = ""
  input_group {
		input_name = ""
		protocol = ""
		description = ""
		allow_ip_list = 
		srt_settings {
			mode = ""
			stream_id = ""
			latency = 
			recv_latency = 
			peer_latency = 
			peer_idle_timeout = 
			passphrase = ""
			pb_key_len = 
			source_addresses {
				ip = ""
				port = 
			}
		}
		rtp_settings {
			fec = ""
			idle_timeout = 
		}
		fail_over = ""
		rtmp_pull_settings {
			source_addresses {
				tc_url = ""
				stream_key = ""
			}
		}
		rtsp_pull_settings {
			source_addresses {
				url = ""
			}
		}
		hls_pull_settings {
			source_addresses {
				url = ""
			}
		}
		resilient_stream {
			enable = 
			buffer_time = 
		}

  }
}

`
