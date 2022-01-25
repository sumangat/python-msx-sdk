/*
 * MSX SDK
 * MSX SDK client.
 *
 * The version of the OpenAPI document: 1.0.9
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.cisco.msx.platform.model;

import java.util.Objects;
import java.util.Arrays;
import com.cisco.msx.platform.model.DeviceTemplateDetails;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * DeviceTemplateBatchAttachRequest
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class DeviceTemplateBatchAttachRequest {
  public static final String SERIALIZED_NAME_DEVICE_IDS = "deviceIds";
  @SerializedName(SERIALIZED_NAME_DEVICE_IDS)
  private List<String> deviceIds = null;

  public static final String SERIALIZED_NAME_TEMPLATE_DETAILS = "templateDetails";
  @SerializedName(SERIALIZED_NAME_TEMPLATE_DETAILS)
  private List<DeviceTemplateDetails> templateDetails = null;


  public DeviceTemplateBatchAttachRequest deviceIds(List<String> deviceIds) {
    
    this.deviceIds = deviceIds;
    return this;
  }

  public DeviceTemplateBatchAttachRequest addDeviceIdsItem(String deviceIdsItem) {
    if (this.deviceIds == null) {
      this.deviceIds = new ArrayList<>();
    }
    this.deviceIds.add(deviceIdsItem);
    return this;
  }

   /**
   * Get deviceIds
   * @return deviceIds
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getDeviceIds() {
    return deviceIds;
  }


  public void setDeviceIds(List<String> deviceIds) {
    this.deviceIds = deviceIds;
  }


  public DeviceTemplateBatchAttachRequest templateDetails(List<DeviceTemplateDetails> templateDetails) {
    
    this.templateDetails = templateDetails;
    return this;
  }

  public DeviceTemplateBatchAttachRequest addTemplateDetailsItem(DeviceTemplateDetails templateDetailsItem) {
    if (this.templateDetails == null) {
      this.templateDetails = new ArrayList<>();
    }
    this.templateDetails.add(templateDetailsItem);
    return this;
  }

   /**
   * Get templateDetails
   * @return templateDetails
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<DeviceTemplateDetails> getTemplateDetails() {
    return templateDetails;
  }


  public void setTemplateDetails(List<DeviceTemplateDetails> templateDetails) {
    this.templateDetails = templateDetails;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeviceTemplateBatchAttachRequest deviceTemplateBatchAttachRequest = (DeviceTemplateBatchAttachRequest) o;
    return Objects.equals(this.deviceIds, deviceTemplateBatchAttachRequest.deviceIds) &&
        Objects.equals(this.templateDetails, deviceTemplateBatchAttachRequest.templateDetails);
  }

  @Override
  public int hashCode() {
    return Objects.hash(deviceIds, templateDetails);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeviceTemplateBatchAttachRequest {\n");
    sb.append("    deviceIds: ").append(toIndentedString(deviceIds)).append("\n");
    sb.append("    templateDetails: ").append(toIndentedString(templateDetails)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

