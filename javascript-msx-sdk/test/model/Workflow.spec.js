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
    instance = new JavascriptMsxSdk.Workflow();
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

  describe('Workflow', function() {
    it('should create an instance of Workflow', function() {
      // uncomment below and update the code to test Workflow
      //var instane = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be.a(JavascriptMsxSdk.Workflow);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property schemaId (base name: "schema_id")', function() {
      // uncomment below and update the code to test the property schemaId
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property title (base name: "title")', function() {
      // uncomment below and update the code to test the property title
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property baseType (base name: "base_type")', function() {
      // uncomment below and update the code to test the property baseType
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property objectType (base name: "object_type")', function() {
      // uncomment below and update the code to test the property objectType
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property properties (base name: "properties")', function() {
      // uncomment below and update the code to test the property properties
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property valid (base name: "valid")', function() {
      // uncomment below and update the code to test the property valid
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property workflowValid (base name: "workflow_valid")', function() {
      // uncomment below and update the code to test the property workflowValid
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property categories (base name: "categories")', function() {
      // uncomment below and update the code to test the property categories
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property metadata (base name: "metadata")', function() {
      // uncomment below and update the code to test the property metadata
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property permissions (base name: "permissions")', function() {
      // uncomment below and update the code to test the property permissions
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property variables (base name: "variables")', function() {
      // uncomment below and update the code to test the property variables
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property actions (base name: "actions")', function() {
      // uncomment below and update the code to test the property actions
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property createdOn (base name: "created_on")', function() {
      // uncomment below and update the code to test the property createdOn
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property createdBy (base name: "created_by")', function() {
      // uncomment below and update the code to test the property createdBy
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property updatedOn (base name: "updated_on")', function() {
      // uncomment below and update the code to test the property updatedOn
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property updatedBy (base name: "updated_by")', function() {
      // uncomment below and update the code to test the property updatedBy
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property owner (base name: "owner")', function() {
      // uncomment below and update the code to test the property owner
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

    it('should have the property uniqueName (base name: "unique_name")', function() {
      // uncomment below and update the code to test the property uniqueName
      //var instance = new JavascriptMsxSdk.Workflow();
      //expect(instance).to.be();
    });

  });

}));
