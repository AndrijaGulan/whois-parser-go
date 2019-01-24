/*
 * Go module for whois info parser
 * https://www.likexian.com/
 *
 * Copyright 2014-2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package whois_parser


var (
    name_rule = map[string]string{
        "id": "domain_id",
        "roid": "domain_id",
        "domain id": "domain_id",
        "registry domain id": "domain_id",
        "domain": "domain_name",
        "domain name": "domain_name",
        "registrar id": "registrar_id",
        "registrar iana id": "registrar_id",
        "sponsoring registrar iana id": "registrar_id",
        "registrar": "registrar_name",
        "registrar name": "registrar_name",
        "sponsoring registrar": "registrar_name",
        "last updated by registrar": "registrar_name",
        "authorized agency": "registrar_name",
        "whois server": "whois_server",
        "registrar whois server": "whois_server",
        "referral url": "referral_url",
        "registrar url": "referral_url",
        "status": "domain_status",
        "state": "domain_status",
        "domain status": "domain_status",
        "registration status": "domain_status",
        "nserver": "name_servers",
        "name server": "name_servers",
        "nameservers": "name_servers",
        "name servers": "name_servers",
        "name servers information": "name_servers",
        "host name": "name_servers",
        "dnssec": "domain_dnssec",
        "domain dnssec": "domain_dnssec",
        "created": "created_date",
        "registered": "created_date",
        "create date": "created_date",
        "created on": "created_date",
        "creation date": "created_date",
        "domain registration date": "created_date",
        "registration date": "created_date",
        "domain create date": "created_date",
        "domain name commencement date": "created_date",
        "registered date": "created_date",
        "registered on": "created_date",
        "registration time": "created_date",
        "modified": "updated_date",
        "changed": "updated_date",
        "update date": "updated_date",
        "updated date": "updated_date",
        "updated on": "updated_date",
        "last update": "updated_date",
        "last updated": "updated_date",
        "last updated on": "updated_date",
        "last modified": "updated_date",
        "last updated date": "updated_date",
        "domain last updated date": "updated_date",
        "expire": "expired_date",
        "expires": "expired_date",
        "paid till": "expired_date",
        "expire date": "expired_date",
        "expired date": "expired_date",
        "expiration date": "expired_date",
        "expiration on": "expired_date",
        "registry expiry date": "expired_date",
        "registrar registration expiration date": "expired_date",
        "domain expiration date": "expired_date",
        "expiry date": "expired_date",
        "expiration time": "expired_date",
        "registry registrant id": "registrant_id",
        "registry admin id": "admin_id",
        "registry tech id": "tech_id",
        "registry bill id": "bill_id",
        "registrant c": "registrant_id",
        "registrant id": "registrant_id",
        "registrant contact id": "registrant_id",
        "registrant": "registrant_name",
        "registrant name": "registrant_name",
        "registrant contact": "registrant_name",
        "registrant contact name": "registrant_name",
        "registrant organization": "registrant_organization",
        "registrant contact organization": "registrant_organization",
        "registrant street": "registrant_street",
        "registrant contact street": "registrant_street",
        "registrant address1": "registrant_street",
        "registrant street1": "registrant_street",
        "registrant contact address1": "registrant_street",
        "registrant s address": "registrant_street",
        "registrant city": "registrant_city",
        "registrant contact city": "registrant_city",
        "registrant state province": "registrant_state_province",
        "registrant contact state province": "registrant_state_province",
        "registrant postal code": "registrant_postal_code",
        "registrant contact postal code": "registrant_postal_code",
        "registrant country": "registrant_country",
        "registrant country economy": "registrant_country",
        "registrant contact country": "registrant_country",
        "registrant phone": "registrant_phone",
        "registrant phone number": "registrant_phone",
        "registrant contact phone": "registrant_phone",
        "registrant contact phone number": "registrant_phone",
        "registrant phone ext": "registrant_phone_ext",
        "registrant contact phone ext": "registrant_phone_ext",
        "registrant fax": "registrant_fax",
        "registrant fax number": "registrant_fax",
        "registrant facsimile": "registrant_fax",
        "registrant facsimile number": "registrant_fax",
        "registrant contact fax": "registrant_fax",
        "registrant contact fax number": "registrant_fax",
        "registrant contact facsimile": "registrant_fax",
        "registrant contact facsimile number": "registrant_fax",
        "registrant fax ext": "registrant_fax_ext",
        "registrant contact fax ext": "registrant_fax_ext",
        "registrant mail": "registrant_email",
        "registrant email": "registrant_email",
        "registrant e mail": "registrant_email",
        "registrant contact mail": "registrant_email",
        "registrant contact email": "registrant_email",
        "registrant contact e mail": "registrant_email",
    }
)