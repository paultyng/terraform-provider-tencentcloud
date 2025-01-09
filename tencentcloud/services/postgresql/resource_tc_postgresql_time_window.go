// Code generated by iacg; DO NOT EDIT.
package postgresql

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	postgresv20170312 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/postgres/v20170312"
	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func ResourceTencentCloudPostgresqlTimeWindow() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudPostgresqlTimeWindowCreate,
		Read:   resourceTencentCloudPostgresqlTimeWindowRead,
		Update: resourceTencentCloudPostgresqlTimeWindowUpdate,
		Delete: resourceTencentCloudPostgresqlTimeWindowDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"db_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance id.",
			},

			"maintain_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Maintenance start time. Time zone is UTC+8.",
			},

			"maintain_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maintenance duration, Unit: hours.",
			},

			"maintain_week_days": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Maintenance cycle.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceTencentCloudPostgresqlTimeWindowCreate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_postgresql_time_window.create")()
	defer tccommon.InconsistentCheck(d, meta)()

	var (
		dBInstanceId string
	)
	if v, ok := d.GetOk("db_instance_id"); ok {
		dBInstanceId = v.(string)
	}

	d.SetId(dBInstanceId)

	return resourceTencentCloudPostgresqlTimeWindowUpdate(d, meta)
}

func resourceTencentCloudPostgresqlTimeWindowRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_postgresql_time_window.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := PostgresqlService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	dBInstanceId := d.Id()

	_ = d.Set("db_instance_id", dBInstanceId)

	respData, err := service.DescribePostgresqlTimeWindowById(ctx, dBInstanceId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `postgresql_time_window` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if respData.DBInstanceId != nil {
		_ = d.Set("db_instance_id", respData.DBInstanceId)
	}

	if respData.MaintainStartTime != nil {
		_ = d.Set("maintain_start_time", respData.MaintainStartTime)
	}

	if respData.MaintainDuration != nil {
		_ = d.Set("maintain_duration", respData.MaintainDuration)
	}

	if respData.MaintainWeekDays != nil {
		_ = d.Set("maintain_week_days", respData.MaintainWeekDays)
	}

	return nil
}

func resourceTencentCloudPostgresqlTimeWindowUpdate(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_postgresql_time_window.update")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(tccommon.ContextNil)

	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	dBInstanceId := d.Id()

	needChange := false
	mutableArgs := []string{"maintain_start_time", "maintain_duration", "maintain_week_days"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := postgresv20170312.NewModifyMaintainTimeWindowRequest()

		request.DBInstanceId = helper.String(dBInstanceId)

		if v, ok := d.GetOk("maintain_start_time"); ok {
			request.MaintainStartTime = helper.String(v.(string))
		}

		if v, ok := d.GetOkExists("maintain_duration"); ok {
			request.MaintainDuration = helper.IntUint64(v.(int))
		}

		if v, ok := d.GetOk("maintain_week_days"); ok {
			maintainWeekDaysSet := v.(*schema.Set).List()
			for i := range maintainWeekDaysSet {
				maintainWeekDays := maintainWeekDaysSet[i].(string)
				request.MaintainWeekDays = append(request.MaintainWeekDays, helper.String(maintainWeekDays))
			}
		}

		reqErr := resource.Retry(tccommon.WriteRetryTimeout, func() *resource.RetryError {
			result, e := meta.(tccommon.ProviderMeta).GetAPIV3Conn().UsePostgresV20170312Client().ModifyMaintainTimeWindowWithContext(ctx, request)
			if e != nil {
				return tccommon.RetryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if reqErr != nil {
			log.Printf("[CRITAL]%s update postgresql time window failed, reason:%+v", logId, reqErr)
			return reqErr
		}
	}

	return resourceTencentCloudPostgresqlTimeWindowRead(d, meta)
}

func resourceTencentCloudPostgresqlTimeWindowDelete(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("resource.tencentcloud_postgresql_time_window.delete")()
	defer tccommon.InconsistentCheck(d, meta)()

	return nil
}
