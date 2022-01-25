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
import com.cisco.msx.platform.model.LegacyAbsoluteConfig;
import com.cisco.msx.platform.model.LegacyRelativeConfig;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * LegacyScheduleConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class LegacyScheduleConfig {
  public static final String SERIALIZED_NAME_METHOD = "method";
  @SerializedName(SERIALIZED_NAME_METHOD)
  private String method;

  public static final String SERIALIZED_NAME_RELATIVE = "relative";
  @SerializedName(SERIALIZED_NAME_RELATIVE)
  private LegacyRelativeConfig relative;

  public static final String SERIALIZED_NAME_ABSOLUTE = "absolute";
  @SerializedName(SERIALIZED_NAME_ABSOLUTE)
  private LegacyAbsoluteConfig absolute;


  public LegacyScheduleConfig method(String method) {
    
    this.method = method;
    return this;
  }

   /**
   * Get method
   * @return method
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMethod() {
    return method;
  }


  public void setMethod(String method) {
    this.method = method;
  }


  public LegacyScheduleConfig relative(LegacyRelativeConfig relative) {
    
    this.relative = relative;
    return this;
  }

   /**
   * Get relative
   * @return relative
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public LegacyRelativeConfig getRelative() {
    return relative;
  }


  public void setRelative(LegacyRelativeConfig relative) {
    this.relative = relative;
  }


  public LegacyScheduleConfig absolute(LegacyAbsoluteConfig absolute) {
    
    this.absolute = absolute;
    return this;
  }

   /**
   * Get absolute
   * @return absolute
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public LegacyAbsoluteConfig getAbsolute() {
    return absolute;
  }


  public void setAbsolute(LegacyAbsoluteConfig absolute) {
    this.absolute = absolute;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    LegacyScheduleConfig legacyScheduleConfig = (LegacyScheduleConfig) o;
    return Objects.equals(this.method, legacyScheduleConfig.method) &&
        Objects.equals(this.relative, legacyScheduleConfig.relative) &&
        Objects.equals(this.absolute, legacyScheduleConfig.absolute);
  }

  @Override
  public int hashCode() {
    return Objects.hash(method, relative, absolute);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class LegacyScheduleConfig {\n");
    sb.append("    method: ").append(toIndentedString(method)).append("\n");
    sb.append("    relative: ").append(toIndentedString(relative)).append("\n");
    sb.append("    absolute: ").append(toIndentedString(absolute)).append("\n");
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

