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


/**
 * 
 * @export
 */
export const V1HelloType = {
    Generic: 'HELLO_GENERIC',
    Specific: 'HELLO_SPECIFIC'
} as const;
export type V1HelloType = typeof V1HelloType[keyof typeof V1HelloType];


export function V1HelloTypeFromJSON(json: any): V1HelloType {
    return V1HelloTypeFromJSONTyped(json, false);
}

export function V1HelloTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1HelloType {
    return json as V1HelloType;
}

export function V1HelloTypeToJSON(value?: V1HelloType | null): any {
    return value as any;
}

