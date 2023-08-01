/*
Provides a resource to create a mysql ro_group_load_operation

Example Usage

```hcl
data "tencentcloud_availability_zones_by_product" "zones" {
  product = "cdb"
}

data "tencentcloud_mysql_instance" "example" {
  mysql_id = tencentcloud_mysql_instance.example.id

  depends_on = [tencentcloud_mysql_readonly_instance.example]
}

resource "tencentcloud_vpc" "vpc" {
  name       = "vpc-mysql"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloud_subnet" "subnet" {
  availability_zone = data.tencentcloud_availability_zones_by_product.zones.zones.1.name
  name              = "subnet-mysql"
  vpc_id            = tencentcloud_vpc.vpc.id
  cidr_block        = "10.0.0.0/16"
  is_multicast      = false
}

resource "tencentcloud_security_group" "security_group" {
  name        = "sg-mysql"
  description = "mysql test"
}

resource "tencentcloud_mysql_instance" "example" {
  internet_service  = 1
  engine_version    = "5.7"
  charge_type       = "POSTPAID"
  root_password     = "PassWord123"
  slave_deploy_mode = 0
  availability_zone = data.tencentcloud_availability_zones_by_product.zones.zones.1.name
  slave_sync_mode   = 1
  instance_name     = "tf-example-mysql"
  mem_size          = 4000
  volume_size       = 200
  vpc_id            = tencentcloud_vpc.vpc.id
  subnet_id         = tencentcloud_subnet.subnet.id
  intranet_port     = 3306
  security_groups   = [tencentcloud_security_group.security_group.id]

  tags = {
    name = "test"
  }

  parameters = {
    character_set_server = "utf8"
    max_connections      = "1000"
  }
}

resource "tencentcloud_mysql_readonly_instance" "example" {
  master_instance_id = tencentcloud_mysql_instance.example.id
  instance_name      = "tf-mysql"
  mem_size           = 2000
  volume_size        = 200
  vpc_id             = tencentcloud_vpc.vpc.id
  subnet_id          = tencentcloud_subnet.subnet.id
  intranet_port      = 3306
  security_groups    = [tencentcloud_security_group.security_group.id]

  tags = {
    createBy = "terraform"
  }
}

resource "tencentcloud_mysql_ro_group_load_operation" "ro_group_load_operation" {
  ro_group_id = data.tencentcloud_mysql_instance.example.instance_list.0.ro_groups.0.group_id
}
```

*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	mysql "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudMysqlRoGroupLoadOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudMysqlRoGroupLoadOperationCreate,
		Read:   resourceTencentCloudMysqlRoGroupLoadOperationRead,
		Delete: resourceTencentCloudMysqlRoGroupLoadOperationDelete,

		Schema: map[string]*schema.Schema{
			"ro_group_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The ID of the RO group, in the format: cdbrg-c1nl9rpv.",
			},
		},
	}
}

func resourceTencentCloudMysqlRoGroupLoadOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mysql_ro_group_load_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request   = mysql.NewBalanceRoGroupLoadRequest()
		roGroupId string
	)
	if v, ok := d.GetOk("ro_group_id"); ok {
		roGroupId = v.(string)
		request.RoGroupId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseMysqlClient().BalanceRoGroupLoad(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate mysql roGroupLoadOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(roGroupId)

	return resourceTencentCloudMysqlRoGroupLoadOperationRead(d, meta)
}

func resourceTencentCloudMysqlRoGroupLoadOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mysql_ro_group_load_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudMysqlRoGroupLoadOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_mysql_ro_group_load_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
