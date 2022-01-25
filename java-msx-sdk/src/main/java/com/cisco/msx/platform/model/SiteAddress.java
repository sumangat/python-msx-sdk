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
 * SiteAddress
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2022-01-21T12:39:36.897411-07:00[America/Edmonton]")
public class SiteAddress {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_COMPANY = "company";
  @SerializedName(SERIALIZED_NAME_COMPANY)
  private String company;

  public static final String SERIALIZED_NAME_FORMATTED_ADDRESS = "formattedAddress";
  @SerializedName(SERIALIZED_NAME_FORMATTED_ADDRESS)
  private String formattedAddress;

  public static final String SERIALIZED_NAME_ADDRESS1 = "address1";
  @SerializedName(SERIALIZED_NAME_ADDRESS1)
  private String address1;

  public static final String SERIALIZED_NAME_ADDRESS2 = "address2";
  @SerializedName(SERIALIZED_NAME_ADDRESS2)
  private String address2;

  public static final String SERIALIZED_NAME_CITY = "city";
  @SerializedName(SERIALIZED_NAME_CITY)
  private String city;

  public static final String SERIALIZED_NAME_STATE = "state";
  @SerializedName(SERIALIZED_NAME_STATE)
  private String state;

  public static final String SERIALIZED_NAME_COUNTRY = "country";
  @SerializedName(SERIALIZED_NAME_COUNTRY)
  private String country;

  public static final String SERIALIZED_NAME_POST_CODE = "postCode";
  @SerializedName(SERIALIZED_NAME_POST_CODE)
  private String postCode;


  public SiteAddress name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public SiteAddress company(String company) {
    
    this.company = company;
    return this;
  }

   /**
   * Get company
   * @return company
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCompany() {
    return company;
  }


  public void setCompany(String company) {
    this.company = company;
  }


  public SiteAddress formattedAddress(String formattedAddress) {
    
    this.formattedAddress = formattedAddress;
    return this;
  }

   /**
   * Get formattedAddress
   * @return formattedAddress
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getFormattedAddress() {
    return formattedAddress;
  }


  public void setFormattedAddress(String formattedAddress) {
    this.formattedAddress = formattedAddress;
  }


  public SiteAddress address1(String address1) {
    
    this.address1 = address1;
    return this;
  }

   /**
   * Get address1
   * @return address1
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAddress1() {
    return address1;
  }


  public void setAddress1(String address1) {
    this.address1 = address1;
  }


  public SiteAddress address2(String address2) {
    
    this.address2 = address2;
    return this;
  }

   /**
   * Get address2
   * @return address2
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAddress2() {
    return address2;
  }


  public void setAddress2(String address2) {
    this.address2 = address2;
  }


  public SiteAddress city(String city) {
    
    this.city = city;
    return this;
  }

   /**
   * Get city
   * @return city
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCity() {
    return city;
  }


  public void setCity(String city) {
    this.city = city;
  }


  public SiteAddress state(String state) {
    
    this.state = state;
    return this;
  }

   /**
   * Get state
   * @return state
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getState() {
    return state;
  }


  public void setState(String state) {
    this.state = state;
  }


  public SiteAddress country(String country) {
    
    this.country = country;
    return this;
  }

   /**
   * Get country
   * @return country
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCountry() {
    return country;
  }


  public void setCountry(String country) {
    this.country = country;
  }


  public SiteAddress postCode(String postCode) {
    
    this.postCode = postCode;
    return this;
  }

   /**
   * Get postCode
   * @return postCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPostCode() {
    return postCode;
  }


  public void setPostCode(String postCode) {
    this.postCode = postCode;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    SiteAddress siteAddress = (SiteAddress) o;
    return Objects.equals(this.name, siteAddress.name) &&
        Objects.equals(this.company, siteAddress.company) &&
        Objects.equals(this.formattedAddress, siteAddress.formattedAddress) &&
        Objects.equals(this.address1, siteAddress.address1) &&
        Objects.equals(this.address2, siteAddress.address2) &&
        Objects.equals(this.city, siteAddress.city) &&
        Objects.equals(this.state, siteAddress.state) &&
        Objects.equals(this.country, siteAddress.country) &&
        Objects.equals(this.postCode, siteAddress.postCode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, company, formattedAddress, address1, address2, city, state, country, postCode);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class SiteAddress {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    company: ").append(toIndentedString(company)).append("\n");
    sb.append("    formattedAddress: ").append(toIndentedString(formattedAddress)).append("\n");
    sb.append("    address1: ").append(toIndentedString(address1)).append("\n");
    sb.append("    address2: ").append(toIndentedString(address2)).append("\n");
    sb.append("    city: ").append(toIndentedString(city)).append("\n");
    sb.append("    state: ").append(toIndentedString(state)).append("\n");
    sb.append("    country: ").append(toIndentedString(country)).append("\n");
    sb.append("    postCode: ").append(toIndentedString(postCode)).append("\n");
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

