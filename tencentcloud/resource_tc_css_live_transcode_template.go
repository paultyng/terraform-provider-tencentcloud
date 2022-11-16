/*
Provides a resource to create a css live_transcode_template

Example Usage

```hcl
resource "tencentcloud_css_live_transcode_template" "live_transcode_template" {
  template_name = ""
  video_bitrate = ""
  acodec = ""
  audio_bitrate = ""
  vcodec = ""
  description = ""
  need_video = ""
  width = ""
  need_audio = ""
  height = ""
  fps = ""
  gop = ""
  rotate = ""
  profile = ""
  bitrate_to_orig = ""
  height_to_orig = ""
  fps_to_orig = ""
  ai_trans_code = ""
  adapt_bitrate_percent = ""
  short_edge_as_height = ""
  d_r_m_type = ""
  d_r_m_tracks = ""
}

```
Import

css live_transcode_template can be imported using the id, e.g.
```
$ terraform import tencentcloud_css_live_transcode_template.live_transcode_template liveTranscodeTemplate_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	css "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func resourceTencentCloudCssLiveTranscodeTemplate() *schema.Resource {
	return &schema.Resource{
		Read:   resourceTencentCloudCssLiveTranscodeTemplateRead,
		Create: resourceTencentCloudCssLiveTranscodeTemplateCreate,
		Update: resourceTencentCloudCssLiveTranscodeTemplateUpdate,
		Delete: resourceTencentCloudCssLiveTranscodeTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "template name, only support 0-9 and a-z.",
			},

			"video_bitrate": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "video bitrate, 0 for origin, range 0kbps - 8000kbps.",
			},

			"acodec": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "default acc, not support now.",
			},

			"audio_bitrate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "default 0, range 0 - 500.",
			},

			"vcodec": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "video codec, default origin, support h264/h265/origin.",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "template desc.",
			},

			"need_video": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "keep video or not, default 1 for yes, 0 for no.",
			},

			"width": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "template width, default 0, range 0 - 3000, must be pow of 2.",
			},

			"need_audio": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "keep audio or not, default 1 for yes, 0 for no.",
			},

			"height": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "template height, default 0, range 0 - 3000, must be pow of 2, needed while AiTransCode = 1.",
			},

			"fps": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "video fps, default 0, range 0 - 60.",
			},

			"gop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "gop of the video, second, default origin of the video, range 2 - 6.",
			},

			"rotate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "roate degree, default 0, support 0/90/180/270.",
			},

			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "quality of the video, default baseline, support baseline/main/high.",
			},

			"bitrate_to_orig": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "base on origin bitrate if origin bitrate is lower than the setting bitrate. default 0, 1 for yes, 0 for no.",
			},

			"height_to_orig": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "base on origin height if origin height is lower than the setting height. default 0, 1 for yes, 0 for no.",
			},

			"fps_to_orig": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "base on origin fps if origin fps is lower than the setting fps. default 0, 1 for yes, 0 for no.",
			},

			"ai_trans_code": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "enable high speed mode, default 0, 1 for enable, 0 for no.",
			},

			"adapt_bitrate_percent": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "high speed mode adapt bitrate, support 0 - 0.5.",
			},

			"short_edge_as_height": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "let the short edge as the height.",
			},

			"d_r_m_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "DRM type, support fairplay/normalaes/widevine.",
			},

			"d_r_m_tracks": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "DRM tracks, support AUDIO/SD/HD/UHD1/UHD2.",
			},
		},
	}
}

func resourceTencentCloudCssLiveTranscodeTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_live_transcode_template.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = css.NewCreateLiveTranscodeTemplateRequest()
		response   *css.CreateLiveTranscodeTemplateResponse
		templateId string
	)

	if v, ok := d.GetOk("template_name"); ok {

		request.TemplateName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("video_bitrate"); ok {
		request.VideoBitrate = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("acodec"); ok {

		request.Acodec = helper.String(v.(string))
	}

	if v, ok := d.GetOk("audio_bitrate"); ok {
		request.AudioBitrate = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("vcodec"); ok {

		request.Vcodec = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {

		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOk("need_video"); ok {
		request.NeedVideo = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("width"); ok {
		request.Width = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("need_audio"); ok {
		request.NeedAudio = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("height"); ok {
		request.Height = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("fps"); ok {
		request.Fps = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("gop"); ok {
		request.Gop = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("rotate"); ok {
		request.Rotate = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("profile"); ok {

		request.Profile = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bitrate_to_orig"); ok {
		request.BitrateToOrig = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("height_to_orig"); ok {
		request.HeightToOrig = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("fps_to_orig"); ok {
		request.FpsToOrig = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("ai_trans_code"); ok {
		request.AiTransCode = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("adapt_bitrate_percent"); ok {
		request.AdaptBitratePercent = helper.Float64(v.(float64))
	}

	if v, ok := d.GetOk("short_edge_as_height"); ok {
		request.ShortEdgeAsHeight = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("d_r_m_type"); ok {
		request.DRMType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("d_r_m_tracks"); ok {
		request.DRMTracks = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCssClient().CreateLiveTranscodeTemplate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create css liveTranscodeTemplate failed, reason:%+v", logId, err)
		return err
	}

	templateId = helper.Int64ToStr(*response.Response.TemplateId)

	d.SetId(templateId)
	return resourceTencentCloudCssLiveTranscodeTemplateRead(d, meta)
}

func resourceTencentCloudCssLiveTranscodeTemplateRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_live_transcode_template.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CssService{client: meta.(*TencentCloudClient).apiV3Conn}

	templateId := d.Id()

	liveTranscodeTemplate, err := service.DescribeCssLiveTranscodeTemplate(ctx, helper.Int64(helper.StrToInt64(templateId)))

	if err != nil {
		return err
	}

	if liveTranscodeTemplate == nil {
		d.SetId("")
		return fmt.Errorf("resource `liveTranscodeTemplate` %s does not exist", templateId)
	}

	if liveTranscodeTemplate.TemplateName != nil {
		_ = d.Set("template_name", liveTranscodeTemplate.TemplateName)
	}

	if liveTranscodeTemplate.VideoBitrate != nil {
		_ = d.Set("video_bitrate", liveTranscodeTemplate.VideoBitrate)
	}

	if liveTranscodeTemplate.Acodec != nil {
		_ = d.Set("acodec", liveTranscodeTemplate.Acodec)
	}

	if liveTranscodeTemplate.AudioBitrate != nil {
		_ = d.Set("audio_bitrate", liveTranscodeTemplate.AudioBitrate)
	}

	if liveTranscodeTemplate.Vcodec != nil {
		_ = d.Set("vcodec", liveTranscodeTemplate.Vcodec)
	}

	if liveTranscodeTemplate.Description != nil {
		_ = d.Set("description", liveTranscodeTemplate.Description)
	}

	if liveTranscodeTemplate.NeedVideo != nil {
		_ = d.Set("need_video", liveTranscodeTemplate.NeedVideo)
	}

	if liveTranscodeTemplate.Width != nil {
		_ = d.Set("width", liveTranscodeTemplate.Width)
	}

	if liveTranscodeTemplate.NeedAudio != nil {
		_ = d.Set("need_audio", liveTranscodeTemplate.NeedAudio)
	}

	if liveTranscodeTemplate.Height != nil {
		_ = d.Set("height", liveTranscodeTemplate.Height)
	}

	if liveTranscodeTemplate.Fps != nil {
		_ = d.Set("fps", liveTranscodeTemplate.Fps)
	}

	if liveTranscodeTemplate.Gop != nil {
		_ = d.Set("gop", liveTranscodeTemplate.Gop)
	}

	if liveTranscodeTemplate.Rotate != nil {
		_ = d.Set("rotate", liveTranscodeTemplate.Rotate)
	}

	if liveTranscodeTemplate.Profile != nil {
		_ = d.Set("profile", liveTranscodeTemplate.Profile)
	}

	if liveTranscodeTemplate.BitrateToOrig != nil {
		_ = d.Set("bitrate_to_orig", liveTranscodeTemplate.BitrateToOrig)
	}

	if liveTranscodeTemplate.HeightToOrig != nil {
		_ = d.Set("height_to_orig", liveTranscodeTemplate.HeightToOrig)
	}

	if liveTranscodeTemplate.FpsToOrig != nil {
		_ = d.Set("fps_to_orig", liveTranscodeTemplate.FpsToOrig)
	}

	if liveTranscodeTemplate.AiTransCode != nil {
		_ = d.Set("ai_trans_code", liveTranscodeTemplate.AiTransCode)
	}

	if liveTranscodeTemplate.AdaptBitratePercent != nil {
		_ = d.Set("adapt_bitrate_percent", liveTranscodeTemplate.AdaptBitratePercent)
	}

	if liveTranscodeTemplate.ShortEdgeAsHeight != nil {
		_ = d.Set("short_edge_as_height", liveTranscodeTemplate.ShortEdgeAsHeight)
	}

	if liveTranscodeTemplate.DRMType != nil {
		_ = d.Set("d_r_m_type", liveTranscodeTemplate.DRMType)
	}

	if liveTranscodeTemplate.DRMTracks != nil {
		_ = d.Set("d_r_m_tracks", liveTranscodeTemplate.DRMTracks)
	}

	return nil
}

func resourceTencentCloudCssLiveTranscodeTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_live_transcode_template.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.TODO(), logIdKey, logId)

	request := css.NewModifyLiveTranscodeTemplateRequest()

	request.TemplateId = helper.Int64(helper.StrToInt64(d.Id()))

	if d.HasChange("template_name") {

		return fmt.Errorf("`template_name` do not support change now.")

	}

	if d.HasChange("video_bitrate") {
		if v, ok := d.GetOk("video_bitrate"); ok {
			request.VideoBitrate = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("acodec") {
		if v, ok := d.GetOk("acodec"); ok {
			request.Acodec = helper.String(v.(string))
		}
	}

	if d.HasChange("audio_bitrate") {
		return fmt.Errorf("`audio_bitrate` do not support change now.")
	}

	if d.HasChange("vcodec") {
		if v, ok := d.GetOk("vcodec"); ok {
			request.Vcodec = helper.String(v.(string))
		}
	}

	if d.HasChange("description") {
		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}
	}

	if d.HasChange("need_video") {
		if v, ok := d.GetOk("need_video"); ok {
			request.NeedVideo = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("width") {
		if v, ok := d.GetOk("width"); ok {
			request.Width = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("need_audio") {
		if v, ok := d.GetOk("need_audio"); ok {
			request.NeedAudio = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("height") {
		if v, ok := d.GetOk("height"); ok {
			request.Height = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("fps") {
		if v, ok := d.GetOk("fps"); ok {
			request.Fps = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("gop") {
		if v, ok := d.GetOk("gop"); ok {
			request.Gop = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("rotate") {
		if v, ok := d.GetOk("rotate"); ok {
			request.Rotate = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("profile") {
		if v, ok := d.GetOk("profile"); ok {
			request.Profile = helper.String(v.(string))
		}
	}

	if d.HasChange("bitrate_to_orig") {
		if v, ok := d.GetOk("bitrate_to_orig"); ok {
			request.BitrateToOrig = helper.IntInt64(v.(int))
		}

	}

	if d.HasChange("height_to_orig") {
		if v, ok := d.GetOk("height_to_orig"); ok {
			request.HeightToOrig = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("fps_to_orig") {
		if v, ok := d.GetOk("fps_to_orig"); ok {
			request.FpsToOrig = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("ai_trans_code") {
		return fmt.Errorf("`ai_trans_code` do not support change now.")
	}

	if d.HasChange("adapt_bitrate_percent") {
		if v, ok := d.GetOk("adapt_bitrate_percent"); ok {
			request.AdaptBitratePercent = helper.Float64(v.(float64))
		}
	}

	if d.HasChange("short_edge_as_height") {
		if v, ok := d.GetOk("short_edge_as_height"); ok {
			request.ShortEdgeAsHeight = helper.IntInt64(v.(int))
		}

	}

	if d.HasChange("d_r_m_type") {
		if v, ok := d.GetOk("d_r_m_type"); ok {
			request.DRMType = helper.String(v.(string))
		}

	}

	if d.HasChange("d_r_m_tracks") {
		if v, ok := d.GetOk("d_r_m_tracks"); ok {
			request.DRMTracks = helper.String(v.(string))
		}

	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCssClient().ModifyLiveTranscodeTemplate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create css liveTranscodeTemplate failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCssLiveTranscodeTemplateRead(d, meta)
}

func resourceTencentCloudCssLiveTranscodeTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloud_css_live_transcode_template.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CssService{client: meta.(*TencentCloudClient).apiV3Conn}

	templateId := d.Id()

	if err := service.DeleteCssLiveTranscodeTemplateById(ctx, helper.Int64(helper.StrToInt64(templateId))); err != nil {
		return err
	}

	return nil
}
