/**
 * ORY Keto
 * Ory Keto is a cloud native access control server providing best-practice patterns (RBAC, ABAC, ACL, AWS IAM Policies, Kubernetes Roles, ...) via REST APIs.
 *
 * The version of the OpenAPI document: Latest
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from '../api';

export class OryAccessControlPolicy {
    /**
    * Actions is an array representing all the actions this ORY Access Policy applies to.
    */
    'actions'?: Array<string>;
    /**
    * Conditions represents a keyed object of conditions under which this ORY Access Policy is active.
    */
    'conditions'?: { [key: string]: object; };
    /**
    * Description is an optional, human-readable description.
    */
    'description'?: string;
    /**
    * Effect is the effect of this ORY Access Policy. It can be \"allow\" or \"deny\".
    */
    'effect'?: string;
    /**
    * ID is the unique identifier of the ORY Access Policy. It is used to query, update, and remove the ORY Access Policy.
    */
    'id'?: string;
    /**
    * Resources is an array representing all the resources this ORY Access Policy applies to.
    */
    'resources'?: Array<string>;
    /**
    * Subjects is an array representing all the subjects this ORY Access Policy applies to.
    */
    'subjects'?: Array<string>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "actions",
            "baseName": "actions",
            "type": "Array<string>"
        },
        {
            "name": "conditions",
            "baseName": "conditions",
            "type": "{ [key: string]: object; }"
        },
        {
            "name": "description",
            "baseName": "description",
            "type": "string"
        },
        {
            "name": "effect",
            "baseName": "effect",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "resources",
            "baseName": "resources",
            "type": "Array<string>"
        },
        {
            "name": "subjects",
            "baseName": "subjects",
            "type": "Array<string>"
        }    ];

    static getAttributeTypeMap() {
        return OryAccessControlPolicy.attributeTypeMap;
    }
}

