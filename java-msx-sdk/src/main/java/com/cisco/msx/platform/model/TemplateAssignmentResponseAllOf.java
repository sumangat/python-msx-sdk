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

/**
 * TemplateAssignmentResponseAllOf
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class TemplateAssignmentResponseAllOf {
  /**
   * Gets or Sets responseStatus
   */
  @JsonAdapter(ResponseStatusEnum.Adapter.class)
  public enum ResponseStatusEnum {
    NEW("NEW"),
    
    OK("OK"),
    
    CONFLICT("CONFLICT"),
    
    UNASSIGNED("UNASSIGNED"),
    
    ERROR("ERROR");

    private String value;

    ResponseStatusEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static ResponseStatusEnum fromValue(String value) {
      for (ResponseStatusEnum b : ResponseStatusEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<ResponseStatusEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final ResponseStatusEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public ResponseStatusEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return ResponseStatusEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_RESPONSE_STATUS = "responseStatus";
  @SerializedName(SERIALIZED_NAME_RESPONSE_STATUS)
  private ResponseStatusEnum responseStatus;


   /**
   * Get responseStatus
   * @return responseStatus
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ResponseStatusEnum getResponseStatus() {
    return responseStatus;
  }




  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    TemplateAssignmentResponseAllOf templateAssignmentResponseAllOf = (TemplateAssignmentResponseAllOf) o;
    return Objects.equals(this.responseStatus, templateAssignmentResponseAllOf.responseStatus);
  }

  @Override
  public int hashCode() {
    return Objects.hash(responseStatus);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class TemplateAssignmentResponseAllOf {\n");
    sb.append("    responseStatus: ").append(toIndentedString(responseStatus)).append("\n");
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

