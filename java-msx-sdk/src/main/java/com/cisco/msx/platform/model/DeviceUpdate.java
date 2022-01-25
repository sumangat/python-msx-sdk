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
import com.cisco.msx.platform.model.DeviceComplianceState;
import com.cisco.msx.platform.model.DevicePatch;
import com.cisco.msx.platform.model.DeviceUpdateAllOf;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * DeviceUpdate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class DeviceUpdate {
  public static final String SERIALIZED_NAME_SERVICE_TYPE = "serviceType";
  @SerializedName(SERIALIZED_NAME_SERVICE_TYPE)
  private String serviceType;

  public static final String SERIALIZED_NAME_TAGS = "tags";
  @SerializedName(SERIALIZED_NAME_TAGS)
  private Map<String, String> tags = null;

  public static final String SERIALIZED_NAME_MANAGED = "managed";
  @SerializedName(SERIALIZED_NAME_MANAGED)
  private Boolean managed = false;

  public static final String SERIALIZED_NAME_ONBOARD_TYPE = "onboardType";
  @SerializedName(SERIALIZED_NAME_ONBOARD_TYPE)
  private String onboardType;

  public static final String SERIALIZED_NAME_ONBOARD_INFORMATION = "onboardInformation";
  @SerializedName(SERIALIZED_NAME_ONBOARD_INFORMATION)
  private Map<String, Object> onboardInformation = null;

  public static final String SERIALIZED_NAME_ATTRIBUTES = "attributes";
  @SerializedName(SERIALIZED_NAME_ATTRIBUTES)
  private Map<String, Object> attributes = null;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_MODEL = "model";
  @SerializedName(SERIALIZED_NAME_MODEL)
  private String model;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_SUB_TYPE = "subType";
  @SerializedName(SERIALIZED_NAME_SUB_TYPE)
  private String subType;

  public static final String SERIALIZED_NAME_SERIAL_KEY = "serialKey";
  @SerializedName(SERIALIZED_NAME_SERIAL_KEY)
  private String serialKey;

  public static final String SERIALIZED_NAME_VERSION = "version";
  @SerializedName(SERIALIZED_NAME_VERSION)
  private String version;

  public static final String SERIALIZED_NAME_COMPLIANCE_STATE = "complianceState";
  @SerializedName(SERIALIZED_NAME_COMPLIANCE_STATE)
  private DeviceComplianceState complianceState;


  public DeviceUpdate serviceType(String serviceType) {
    
    this.serviceType = serviceType;
    return this;
  }

   /**
   * Get serviceType
   * @return serviceType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServiceType() {
    return serviceType;
  }


  public void setServiceType(String serviceType) {
    this.serviceType = serviceType;
  }


  public DeviceUpdate tags(Map<String, String> tags) {
    
    this.tags = tags;
    return this;
  }

  public DeviceUpdate putTagsItem(String key, String tagsItem) {
    if (this.tags == null) {
      this.tags = new HashMap<>();
    }
    this.tags.put(key, tagsItem);
    return this;
  }

   /**
   * Get tags
   * @return tags
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, String> getTags() {
    return tags;
  }


  public void setTags(Map<String, String> tags) {
    this.tags = tags;
  }


  public DeviceUpdate managed(Boolean managed) {
    
    this.managed = managed;
    return this;
  }

   /**
   * Get managed
   * @return managed
  **/
  @ApiModelProperty(required = true, value = "")

  public Boolean getManaged() {
    return managed;
  }


  public void setManaged(Boolean managed) {
    this.managed = managed;
  }


  public DeviceUpdate onboardType(String onboardType) {
    
    this.onboardType = onboardType;
    return this;
  }

   /**
   * Get onboardType
   * @return onboardType
  **/
  @ApiModelProperty(required = true, value = "")

  public String getOnboardType() {
    return onboardType;
  }


  public void setOnboardType(String onboardType) {
    this.onboardType = onboardType;
  }


  public DeviceUpdate onboardInformation(Map<String, Object> onboardInformation) {
    
    this.onboardInformation = onboardInformation;
    return this;
  }

  public DeviceUpdate putOnboardInformationItem(String key, Object onboardInformationItem) {
    if (this.onboardInformation == null) {
      this.onboardInformation = new HashMap<>();
    }
    this.onboardInformation.put(key, onboardInformationItem);
    return this;
  }

   /**
   * Get onboardInformation
   * @return onboardInformation
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, Object> getOnboardInformation() {
    return onboardInformation;
  }


  public void setOnboardInformation(Map<String, Object> onboardInformation) {
    this.onboardInformation = onboardInformation;
  }


  public DeviceUpdate attributes(Map<String, Object> attributes) {
    
    this.attributes = attributes;
    return this;
  }

  public DeviceUpdate putAttributesItem(String key, Object attributesItem) {
    if (this.attributes == null) {
      this.attributes = new HashMap<>();
    }
    this.attributes.put(key, attributesItem);
    return this;
  }

   /**
   * Get attributes
   * @return attributes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Map<String, Object> getAttributes() {
    return attributes;
  }


  public void setAttributes(Map<String, Object> attributes) {
    this.attributes = attributes;
  }


  public DeviceUpdate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public DeviceUpdate model(String model) {
    
    this.model = model;
    return this;
  }

   /**
   * Get model
   * @return model
  **/
  @ApiModelProperty(required = true, value = "")

  public String getModel() {
    return model;
  }


  public void setModel(String model) {
    this.model = model;
  }


  public DeviceUpdate type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @ApiModelProperty(required = true, value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public DeviceUpdate subType(String subType) {
    
    this.subType = subType;
    return this;
  }

   /**
   * Get subType
   * @return subType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSubType() {
    return subType;
  }


  public void setSubType(String subType) {
    this.subType = subType;
  }


  public DeviceUpdate serialKey(String serialKey) {
    
    this.serialKey = serialKey;
    return this;
  }

   /**
   * Get serialKey
   * @return serialKey
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSerialKey() {
    return serialKey;
  }


  public void setSerialKey(String serialKey) {
    this.serialKey = serialKey;
  }


  public DeviceUpdate version(String version) {
    
    this.version = version;
    return this;
  }

   /**
   * Get version
   * @return version
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVersion() {
    return version;
  }


  public void setVersion(String version) {
    this.version = version;
  }


  public DeviceUpdate complianceState(DeviceComplianceState complianceState) {
    
    this.complianceState = complianceState;
    return this;
  }

   /**
   * Get complianceState
   * @return complianceState
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public DeviceComplianceState getComplianceState() {
    return complianceState;
  }


  public void setComplianceState(DeviceComplianceState complianceState) {
    this.complianceState = complianceState;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeviceUpdate deviceUpdate = (DeviceUpdate) o;
    return Objects.equals(this.serviceType, deviceUpdate.serviceType) &&
        Objects.equals(this.tags, deviceUpdate.tags) &&
        Objects.equals(this.managed, deviceUpdate.managed) &&
        Objects.equals(this.onboardType, deviceUpdate.onboardType) &&
        Objects.equals(this.onboardInformation, deviceUpdate.onboardInformation) &&
        Objects.equals(this.attributes, deviceUpdate.attributes) &&
        Objects.equals(this.name, deviceUpdate.name) &&
        Objects.equals(this.model, deviceUpdate.model) &&
        Objects.equals(this.type, deviceUpdate.type) &&
        Objects.equals(this.subType, deviceUpdate.subType) &&
        Objects.equals(this.serialKey, deviceUpdate.serialKey) &&
        Objects.equals(this.version, deviceUpdate.version) &&
        Objects.equals(this.complianceState, deviceUpdate.complianceState);
  }

  @Override
  public int hashCode() {
    return Objects.hash(serviceType, tags, managed, onboardType, onboardInformation, attributes, name, model, type, subType, serialKey, version, complianceState);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeviceUpdate {\n");
    sb.append("    serviceType: ").append(toIndentedString(serviceType)).append("\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
    sb.append("    managed: ").append(toIndentedString(managed)).append("\n");
    sb.append("    onboardType: ").append(toIndentedString(onboardType)).append("\n");
    sb.append("    onboardInformation: ").append(toIndentedString(onboardInformation)).append("\n");
    sb.append("    attributes: ").append(toIndentedString(attributes)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    model: ").append(toIndentedString(model)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    subType: ").append(toIndentedString(subType)).append("\n");
    sb.append("    serialKey: ").append(toIndentedString(serialKey)).append("\n");
    sb.append("    version: ").append(toIndentedString(version)).append("\n");
    sb.append("    complianceState: ").append(toIndentedString(complianceState)).append("\n");
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

