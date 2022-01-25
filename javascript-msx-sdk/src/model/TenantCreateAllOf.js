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

import ApiClient from '../ApiClient';

/**
 * The TenantCreateAllOf model module.
 * @module model/TenantCreateAllOf
 * @version 1.0.9
 */
class TenantCreateAllOf {
    /**
     * Constructs a new <code>TenantCreateAllOf</code>.
     * @alias module:model/TenantCreateAllOf
     */
    constructor() { 
        
        TenantCreateAllOf.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TenantCreateAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TenantCreateAllOf} obj Optional instance to populate.
     * @return {module:model/TenantCreateAllOf} The populated <code>TenantCreateAllOf</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TenantCreateAllOf();

            if (data.hasOwnProperty('parentId')) {
                obj['parentId'] = ApiClient.convertToType(data['parentId'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} parentId
 */
TenantCreateAllOf.prototype['parentId'] = undefined;

/**
 * @member {String} externalId
 */
TenantCreateAllOf.prototype['externalId'] = undefined;






export default TenantCreateAllOf;

