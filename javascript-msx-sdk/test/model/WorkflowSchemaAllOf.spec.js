/**
 * MSX SDK
 * MSX SDK client.
 *
 * The version of the OpenAPI document: 1.0.9
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.JavascriptMsxSdk);
  }
}(this, function(expect, JavascriptMsxSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('WorkflowSchemaAllOf', function() {
    it('should create an instance of WorkflowSchemaAllOf', function() {
      // uncomment below and update the code to test WorkflowSchemaAllOf
      //var instane = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be.a(JavascriptMsxSdk.WorkflowSchemaAllOf);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property schemaId (base name: "schema_id")', function() {
      // uncomment below and update the code to test the property schemaId
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property title (base name: "title")', function() {
      // uncomment below and update the code to test the property title
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property baseType (base name: "base_type")', function() {
      // uncomment below and update the code to test the property baseType
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property version (base name: "version")', function() {
      // uncomment below and update the code to test the property version
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property invisible (base name: "invisible")', function() {
      // uncomment below and update the code to test the property invisible
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property inherits (base name: "inherits")', function() {
      // uncomment below and update the code to test the property inherits
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property accessMeta (base name: "access_meta")', function() {
      // uncomment below and update the code to test the property accessMeta
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property variableSchema (base name: "variable_schema")', function() {
      // uncomment below and update the code to test the property variableSchema
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property propertySchema (base name: "property_schema")', function() {
      // uncomment below and update the code to test the property propertySchema
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property outputSchema (base name: "output_schema")', function() {
      // uncomment below and update the code to test the property outputSchema
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property viewConfig (base name: "view_config")', function() {
      // uncomment below and update the code to test the property viewConfig
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

    it('should have the property attributes (base name: "attributes")', function() {
      // uncomment below and update the code to test the property attributes
      //var instance = new JavascriptMsxSdk.WorkflowSchemaAllOf();
      //expect(instance).to.be();
    });

  });

}));
