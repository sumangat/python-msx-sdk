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
import BillingEvent from './BillingEvent';

/**
 * The BillingEventsPageAllOf model module.
 * @module model/BillingEventsPageAllOf
 * @version 1.0.9
 */
class BillingEventsPageAllOf {
    /**
     * Constructs a new <code>BillingEventsPageAllOf</code>.
     * @alias module:model/BillingEventsPageAllOf
     */
    constructor() { 
        
        BillingEventsPageAllOf.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>BillingEventsPageAllOf</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/BillingEventsPageAllOf} obj Optional instance to populate.
     * @return {module:model/BillingEventsPageAllOf} The populated <code>BillingEventsPageAllOf</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new BillingEventsPageAllOf();

            if (data.hasOwnProperty('contents')) {
                obj['contents'] = ApiClient.convertToType(data['contents'], [BillingEvent]);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/BillingEvent>} contents
 */
BillingEventsPageAllOf.prototype['contents'] = undefined;






export default BillingEventsPageAllOf;

