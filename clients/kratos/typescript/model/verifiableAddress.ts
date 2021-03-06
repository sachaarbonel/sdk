/**
 * Ory Kratos
 * Welcome to the ORY Kratos HTTP API documentation!
 *
 * The version of the OpenAPI document: v0.4.6-alpha.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from '../api';

export class VerifiableAddress {
    'expiresAt': Date;
    'id': string;
    'value': string;
    'verified': boolean;
    'verifiedAt'?: Date;
    'via': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "expiresAt",
            "baseName": "expires_at",
            "type": "Date"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "value",
            "baseName": "value",
            "type": "string"
        },
        {
            "name": "verified",
            "baseName": "verified",
            "type": "boolean"
        },
        {
            "name": "verifiedAt",
            "baseName": "verified_at",
            "type": "Date"
        },
        {
            "name": "via",
            "baseName": "via",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return VerifiableAddress.attributeTypeMap;
    }
}

