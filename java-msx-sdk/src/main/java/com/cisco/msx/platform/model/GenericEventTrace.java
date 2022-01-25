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
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.UUID;

/**
 * GenericEventTrace
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class GenericEventTrace {
  public static final String SERIALIZED_NAME_TRACE_ID = "traceId";
  @SerializedName(SERIALIZED_NAME_TRACE_ID)
  private UUID traceId;

  public static final String SERIALIZED_NAME_SPAN_ID = "spanId";
  @SerializedName(SERIALIZED_NAME_SPAN_ID)
  private UUID spanId;

  public static final String SERIALIZED_NAME_PARENT_ID = "parentId";
  @SerializedName(SERIALIZED_NAME_PARENT_ID)
  private UUID parentId;


  public GenericEventTrace traceId(UUID traceId) {
    
    this.traceId = traceId;
    return this;
  }

   /**
   * Get traceId
   * @return traceId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public UUID getTraceId() {
    return traceId;
  }


  public void setTraceId(UUID traceId) {
    this.traceId = traceId;
  }


  public GenericEventTrace spanId(UUID spanId) {
    
    this.spanId = spanId;
    return this;
  }

   /**
   * Get spanId
   * @return spanId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public UUID getSpanId() {
    return spanId;
  }


  public void setSpanId(UUID spanId) {
    this.spanId = spanId;
  }


  public GenericEventTrace parentId(UUID parentId) {
    
    this.parentId = parentId;
    return this;
  }

   /**
   * Get parentId
   * @return parentId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public UUID getParentId() {
    return parentId;
  }


  public void setParentId(UUID parentId) {
    this.parentId = parentId;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GenericEventTrace genericEventTrace = (GenericEventTrace) o;
    return Objects.equals(this.traceId, genericEventTrace.traceId) &&
        Objects.equals(this.spanId, genericEventTrace.spanId) &&
        Objects.equals(this.parentId, genericEventTrace.parentId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(traceId, spanId, parentId);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GenericEventTrace {\n");
    sb.append("    traceId: ").append(toIndentedString(traceId)).append("\n");
    sb.append("    spanId: ").append(toIndentedString(spanId)).append("\n");
    sb.append("    parentId: ").append(toIndentedString(parentId)).append("\n");
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

