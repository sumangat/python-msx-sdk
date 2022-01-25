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
import TemplateApplication from './TemplateApplication';

/**
 * The TemplateApplicationsPageAllOf model module.
 * @module model/TemplateApplicationsPageAllOf
 * @version 1.0.9
 */
class TemplateApplicationsPageAllOf {
    /**
     * Constructs a new <code>TemplateApplicationsPageAllOf</code>.
     * @alias module:model/TemplateApplicationsPageAllOf
     */
    constructor() { 
        
        TemplateApplicationsPageAllOf.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>TemplateApplicationsPageAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/TemplateApplicationsPageAllOf} obj Optional instance to populate.
     * @return {module:model/TemplateApplicationsPageAllOf} The populated <code>TemplateApplicationsPageAllOf</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new TemplateApplicationsPageAllOf();

            if (data.hasOwnProperty('contents')) {
                obj['contents'] = ApiClient.convertToType(data['contents'], [TemplateApplication]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/TemplateApplication>} contents
 */
TemplateApplicationsPageAllOf.prototype['contents'] = undefined;






export default TemplateApplicationsPageAllOf;

