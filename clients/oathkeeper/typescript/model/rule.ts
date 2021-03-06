/**
 * ORY Oathkeeper
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * The version of the OpenAPI document: Latest
 * Contact: hi@ory.am
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from '../api';
import { RuleHandler } from './ruleHandler';
import { RuleMatch } from './ruleMatch';
import { Upstream } from './upstream';

export class Rule {
    /**
    * Authenticators is a list of authentication handlers that will try and authenticate the provided credentials. Authenticators are checked iteratively from index 0 to n and if the first authenticator to return a positive result will be the one used.  If you want the rule to first check a specific authenticator  before \"falling back\" to others, have that authenticator as the first item in the array.
    */
    'authenticators'?: Array<RuleHandler>;
    'authorizer'?: RuleHandler;
    /**
    * Description is a human readable description of this rule.
    */
    'description'?: string;
    /**
    * ID is the unique id of the rule. It can be at most 190 characters long, but the layout of the ID is up to you. You will need this ID later on to update or delete the rule.
    */
    'id'?: string;
    'match'?: RuleMatch;
    /**
    * Mutators is a list of mutation handlers that transform the HTTP request. A common use case is generating a new set of credentials (e.g. JWT) which then will be forwarded to the upstream server.  Mutations are performed iteratively from index 0 to n and should all succeed in order for the HTTP request to be forwarded.
    */
    'mutators'?: Array<RuleHandler>;
    'upstream'?: Upstream;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "authenticators",
            "baseName": "authenticators",
            "type": "Array<RuleHandler>"
        },
        {
            "name": "authorizer",
            "baseName": "authorizer",
            "type": "RuleHandler"
        },
        {
            "name": "description",
            "baseName": "description",
            "type": "string"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "match",
            "baseName": "match",
            "type": "RuleMatch"
        },
        {
            "name": "mutators",
            "baseName": "mutators",
            "type": "Array<RuleHandler>"
        },
        {
            "name": "upstream",
            "baseName": "upstream",
            "type": "Upstream"
        }    ];

    static getAttributeTypeMap() {
        return Rule.attributeTypeMap;
    }
}

