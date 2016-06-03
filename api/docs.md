|HTTP Error Code|Returned element|Description|
|---|---|---|
|200 - OK|Response object per the updated object type, includes the updated object|Successful update|
|201 - Created|Response object per the created object type, includes the created object|Successful create|
|204 – No Content|Empty|Successful delete|
|400 - Bad Request|BadRequestResponse|Applicative error on request|
|401 - Unauthorized|IdentityFaultResponse|Token expiration|
|404 - Not Found|itemNotFoundResponse|Usually will indicate wrong URI|
|405 - Bad Method|badMethodResponse|Method is not supported for object|
|500 – Server Error|faultResponse| |


## Fault response examples

### Expired Token response example
```json
{
  "identityFault": {
    "code": "401",
    "message": "Unauthorized",
    "details": "Token expired"
  }
}
```

### Bad Request response example
```json
{
  "badRequest": {
    "code": "400",
    "message": "Bad Request",
    "details": "Object ID on URL doesn't match object"
  }
}
```

### Bad Method response example
```json
{
  "badMethod": {
    "code": "405",
    "message": "Bad Method"
  }
}
```


### Servers
#### GET
/v0.4/servers - Get list of all servers
#### POST
/v0.4/enrollment-scripts - Generate gw/agent install script

### Resource Groups
#### GET
/v0.4/resource-groups/{id} - Get Resource Group by id
#### GET
/v0.4/resource-groups - Get list of all Resource Groups
#### POST
/v0.4/resource-groups - Create Resource Group

#### PUT
/v0.4/resource-groups/{id} - Update Resource Group

#### DELETE
/v0.4/resource-groups/{id} - Delete resource Group

### User Role
#### GET
/v0.4/user-roles/{id} - Get User Role by id
#### GET
/v0.4/user-roles - Get list of all User Roles

### Security Groups
#### GET
/v0.4/security-groups/{id} - Get Security Group by id
#### GET
/v0.4/ security-groups - Get list of all Security Group
#### PUT
/v0.4/security-groups/{id}/resource-groups/{id} - Assign Resource Group to a Security Group

### LDAP User Groups
#### GET
/v0.4/ldap-user-groups/{id} - Get LDAP User Group by id
#### GET
/v0.4/ldap-user-groups - Get list of all LDAP User Group
#### PUT
/v0.4/ldap-user-groups/{id}/user-roles/{id} - Assign User Role to an LDAP User Group

### Applications
#### GET
/v0.4/applications/{id} - Get Application by id
#### GET
/v0.4/applications - Get list of all Applications

### Simple Server Policy
#### GET
/v0.4/simple-server-policies/{id} - Get Simple Server Policy By Id
#### GET
/v0.4/simple-server-policies - Get list of all Simple Server Policies
#### POST
/v0.4/simple-server-policies - Create Simple Server Policy
#### PUT
/v0.4/simple-server-policies/{id} - Update Simple Server Policy
#### DELETE
/v0.4/simple-server-policies/{id} - Delete Simple Server Policy

### Simple User Policy
#### GET
/v0.4/simple-user-policies/{id} - Get Simple Server Policy By Id
#### GET
/v0.4/simple-user-policies - Get list of all Simple Server Policies
#### POST
/v0.4/simple-user-policies - Create Simple Server Policy
#### PUT
/v0.4/simple-user-policies/{id} - Update Simple Server Policy
#### DELETE
/v0.4/simple-user-policies/{id} - Delete Simple Server Policy

### IP Address Set
#### GET
/v0.4/ip-address-sets/{id} - Get IP Address Set By Id
#### GET
/v0.4/ip-address-sets - Get list of all IP Address Sets
#### POST
/v0.4/ip-address-sets - Create IP Address Set
#### PUT
/v0.4/ip-address-sets/{id} - Update IP Address Set
#### DELETE
/v0.4/ip-address-sets/{id} - Delete IP Address Set