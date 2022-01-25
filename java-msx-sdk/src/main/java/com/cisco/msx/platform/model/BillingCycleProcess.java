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
 * BillingCycleProcess
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class BillingCycleProcess {
  public static final String SERIALIZED_NAME_NEXT_BILLED_ON = "nextBilledOn";
  @SerializedName(SERIALIZED_NAME_NEXT_BILLED_ON)
  private String nextBilledOn;


  public BillingCycleProcess nextBilledOn(String nextBilledOn) {
    
    this.nextBilledOn = nextBilledOn;
    return this;
  }

   /**
   * Get nextBilledOn
   * @return nextBilledOn
  **/
  @ApiModelProperty(example = "2020-09-18T18:37:33.810Z", required = true, value = "")

  public String getNextBilledOn() {
    return nextBilledOn;
  }


  public void setNextBilledOn(String nextBilledOn) {
    this.nextBilledOn = nextBilledOn;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    BillingCycleProcess billingCycleProcess = (BillingCycleProcess) o;
    return Objects.equals(this.nextBilledOn, billingCycleProcess.nextBilledOn);
  }

  @Override
  public int hashCode() {
    return Objects.hash(nextBilledOn);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class BillingCycleProcess {\n");
    sb.append("    nextBilledOn: ").append(toIndentedString(nextBilledOn)).append("\n");
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

