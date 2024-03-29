/* tslint:disable */
/* eslint-disable */
/**
 * skoop/v1/api.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1HelloType } from './V1HelloType';
import {
    V1HelloTypeFromJSON,
    V1HelloTypeFromJSONTyped,
    V1HelloTypeToJSON,
} from './V1HelloType';

/**
 * 
 * @export
 * @interface V1Hello
 */
export interface V1Hello {
    /**
     * 
     * @type {string}
     * @memberof V1Hello
     */
    id?: string;
    /**
     * 
     * @type {Date}
     * @memberof V1Hello
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof V1Hello
     */
    updatedAt?: Date;
    /**
     * 
     * @type {V1HelloType}
     * @memberof V1Hello
     */
    helloType?: V1HelloType;
    /**
     * 
     * @type {string}
     * @memberof V1Hello
     */
    personName?: string;
}

/**
 * Check if a given object implements the V1Hello interface.
 */
export function instanceOfV1Hello(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1HelloFromJSON(json: any): V1Hello {
    return V1HelloFromJSONTyped(json, false);
}

export function V1HelloFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Hello {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': !exists(json, 'id') ? undefined : json['id'],
        'createdAt': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updatedAt': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
        'helloType': !exists(json, 'hello_type') ? undefined : V1HelloTypeFromJSON(json['hello_type']),
        'personName': !exists(json, 'person_name') ? undefined : json['person_name'],
    };
}

export function V1HelloToJSON(value?: V1Hello | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'created_at': value.createdAt === undefined ? undefined : (value.createdAt.toISOString()),
        'updated_at': value.updatedAt === undefined ? undefined : (value.updatedAt.toISOString()),
        'hello_type': V1HelloTypeToJSON(value.helloType),
        'person_name': value.personName,
    };
}

