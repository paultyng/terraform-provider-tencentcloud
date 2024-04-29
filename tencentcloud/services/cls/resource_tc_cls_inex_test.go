package cls_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	tcacctest "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/acctest"
)

// go test -i; go test -test.run TestAccTencentCloudClsIndex_basic -v
func TestAccTencentCloudClsIndex_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { tcacctest.AccPreCheck(t) },
		Providers: tcacctest.AccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClsIndex,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloud_cls_index.example", "topic_id"),
					resource.TestCheckResourceAttr("tencentcloud_cls_index.example", "status", "true"),
					resource.TestCheckResourceAttr("tencentcloud_cls_index.example", "include_internal_fields", "true"),
				),
			},
			{
				ResourceName:      "tencentcloud_cls_index.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccClsIndexUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloud_cls_index.example", "topic_id"),
					resource.TestCheckResourceAttr("tencentcloud_cls_index.example", "status", "false"),
					resource.TestCheckResourceAttr("tencentcloud_cls_index.example", "include_internal_fields", "true"),
				),
			},
		},
	})
}

const testAccClsIndex = `
resource "tencentcloud_cls_logset" "example" {
  logset_name = "tf_example"
  tags        = {
    "demo" = "test"
  }
}

resource "tencentcloud_cls_topic" "example" {
  topic_name           = "tf_example"
  logset_id            = tencentcloud_cls_logset.example.id
  auto_split           = false
  max_split_partitions = 20
  partition_count      = 1
  period               = 30
  storage_type         = "hot"
  describes            = "Test Demo."
  hot_period           = 10
  tags                 = {
    "test" = "test",
  }
}

resource "tencentcloud_cls_index" "example" {
  topic_id = tencentcloud_cls_topic.example.id

  rule {
    full_text {
      case_sensitive = true
      tokenizer      = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
      contain_z_h    = true
    }

    key_value {
      case_sensitive = true
      key_values {
        key = "hello"
        value {
          contain_z_h = true
          sql_flag    = true
          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
          type        = "text"
        }
      }

      key_values {
        key = "world"
        value {
          contain_z_h = true
          sql_flag    = true
          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
          type        = "text"
        }
      }
    }

    tag {
      case_sensitive = true
      key_values {
        key = "terraform"
        value {
          contain_z_h = true
          sql_flag    = true
          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
          type        = "text"
        }
      }
    }

    dynamic_index {
      status = true
    }
  }
  status                  = true
  include_internal_fields = true
  metadata_flag           = 1
}
`

const testAccClsIndexUpdate = `
resource "tencentcloud_cls_logset" "example" {
  logset_name = "tf_example"
  tags        = {
    "demo" = "test"
  }
}

resource "tencentcloud_cls_topic" "example" {
  topic_name           = "tf_example"
  logset_id            = tencentcloud_cls_logset.example.id
  auto_split           = false
  max_split_partitions = 20
  partition_count      = 1
  period               = 30
  storage_type         = "hot"
  describes            = "Test Demo."
  hot_period           = 10
  tags                 = {
    "test" = "test",
  }
}

resource "tencentcloud_cls_index" "example" {
  topic_id = tencentcloud_cls_topic.example.id

  rule {
    full_text {
      case_sensitive = true
      tokenizer      = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
      contain_z_h    = true
    }

    key_value {
      case_sensitive = true
      key_values {
        key = "hello"
        value {
          contain_z_h = true
          sql_flag    = true
          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
          type        = "text"
        }
      }

      key_values {
        key = "world"
        value {
          contain_z_h = true
          sql_flag    = true
          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
          type        = "text"
        }
      }
    }

    tag {
      case_sensitive = true
      key_values {
        key = "terraform"
        value {
          contain_z_h = true
          sql_flag    = true
          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
          type        = "text"
        }
      }
    }

    dynamic_index {
      status = false
    }
  }
  status                  = false
  include_internal_fields = true
  metadata_flag           = 1
}
`
