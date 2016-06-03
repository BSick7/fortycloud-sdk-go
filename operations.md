# API Operations

## Entities

- [Applications](#applications)
- [Audit Logs](#audit-logs)
- [DNAT Actions](#dnat-actions)
- [DNAT Rules](#dnat-rules)
- [Gateways](#gateways)
- [IP Address Sets](#ip-address-sets)
- [IP Pools](#ip-pools)
- [IPSec Connections](#ipsec-connections)
- [IPSec Peers](#ipsec-peers)
- [LDAP Servers](#ldap-servers)
- [LDAP User Groups](#ldap-user-groups)
- [Radius Servers](#radius-servers)
- [Resource Groups](#resource-groups)
- [Security Groups](#security-groups)
- [Servers](#servers)
- [Simple Server Policies](#simple-server-policies)
- [Simple User Policies](#simple-user-policies)
- [SNAT Actions](#snat-actions)
- [SNAT Rules](#snat-rules)
- [Subnets](#subnets)
- [System Alerts](#system-alerts)
- [System Events](#system-events)
- [Users](#users)
- [User Roles](#user-roles)

### Applications

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/applications/{id}                         | Get Application by id                     |
| GET    | /v0.4/applications                              | Get list of all Applications              |
| POST   | /v0.4/applications                              | Create Application                        |
| PUT    | /v0.4/applications/{id}                         | Update Application                        |
| DELETE | /v0.4/applications/{id}                         | Delete Application                        |

### Audit Logs

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/audit-logs/{id}                           | Get Audit Log by id                       |
| GET    | /v0.4/audit-logs                                | Get list of all Audit Logs                |

### DNAT Actions

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/dnat-actions/{id}                         | Get DNAT Action by id                     |
| GET    | /v0.4/dnat-actions                              | Get list of all DNAT Actions              |
| POST   | /v0.4/dnat-actions                              | Create DNAT Action                        |
| PUT    | /v0.4/dnat-actions/{id}                         | Update DNAT Action                        |
| DELETE | /v0.4/dnat-actions/{id}                         | Delete DNAT Action                        |

### DNAT Rules

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/dnat-rules/{id}                           | Get DNAT Rule by id                       |
| GET    | /v0.4/dnat-rules                                | Get list of all DNAT Rules                |
| POST   | /v0.4/dnat-rules                                | Create DNAT Rule                          |
| PUT    | /v0.4/dnat-rules/{id}                           | Update DNAT Rule                          |
| PUT    | /v0.4/dnat-rules/{id}/dnat-actions/{id}         | Assign DNAT Action to DNAT Rule           |
| DELETE | /v0.4/dnat-rules/{id}                           | Delete DNAT Action                        |

### Gateways

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/gateways/{id}                             | Get Gateway by id                         |
| GET    | /v0.4/gateways                                  | Get list of all Gateways                  |
| PUT    | /v0.4/gateways/{id}                             | Update Gateway                            |
| DELETE | /v0.4/gateways/{id}                             | Delete Gateway                            |
| POST   | /v0.4/installation-scripts                      | Generate gateway install script           |
| POST   | /v0.4/registration-tokens                       | Generate gateway registration token       |

### IP Address Sets

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/ip-address-sets/{id}                      | Get IP Address Set by id                  |
| GET    | /v0.4/ip-address-sets                           | Get list of all IP Address Sets           |
| POST   | /v0.4/ip-address-sets                           | Create IP Address Set                     |
| PUT    | /v0.4/ip-address-sets/{id}                      | Update IP Address Set                     |
| DELETE | /v0.4/ip-address-sets/{id}                      | Delete IP Address Set                     |

### IP Pools

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/ip-pools/{id}                             | Get IP Pool by id                         |
| GET    | /v0.4/ip-pools                                  | Get list of all IP Pool                   |
| POST   | /v0.4/ip-pools                                  | Create IP Pool                            |
| PUT    | /v0.4/ip-pools/{id}                             | Update IP Pool                            |
| DELETE | /v0.4/ip-pools/{id}                             | Delete IP Pool                            |

### IPSec Connections

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/ip-sec-connections/{id}                   | Get IPSec Connection by id                |
| GET    | /v0.4/ip-sec-connections                        | Get list of all IPSec Connections         |
| POST   | /v0.4/ip-sec-connections                        | Create IPSec Connection                   |
| PUT    | /v0.4/ip-sec-connections/{id}                   | Update IPSec Connection                   |
| DELETE | /v0.4/ip-sec-connections/{id}                   | Delete IPSec Connection                   |

### IPSec Peer Connections

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/ip-sec-peer-connections/{id}              | Get IPSec Connection to Peer by id        |
| GET    | /v0.4/ip-sec-peer-connections                   | Get list of all IPSec Connection to Peer  |
| POST   | /v0.4/ip-sec-peer-connections                   | Create IPSec Connection to Peer           |
| PUT    | /v0.4/ip-sec-peer-connections/{id}              | Update IPSec Connection to Peer           |
| DELETE | /v0.4/ip-sec-peer-connections/{id}              | Delete IPSec Connection to Peer           |

### IPSec Peers

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/ip-sec-peers/{id}                         | Get IPSec Peer by id                      |
| GET    | /v0.4/ip-sec-peers                              | Get list of all IPSec Peers               |
| POST   | /v0.4/ip-sec-peers                              | Create IPSec Peer                         |
| PUT    | /v0.4/ip-sec-peers/{id}                         | Update IPSec Peer                         |
| DELETE | /v0.4/ip-sec-peers/{id}                         | Delete IPSec Peer                         |

### LDAP Servers

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/ldap-servers/{id}                         | Get LDAP Server by id                     |
| GET    | /v0.4/ldap-servers                              | Get list of all LDAP Servers              |
| POST   | /v0.4/ldap-servers                              | Create LDAP Server                        |
| PUT    | /v0.4/ldap-servers/{id}                         | Update LDAP Server                        |
| DELETE | /v0.4/ldap-servers/{id}                         | Delete LDAP Server                        |

### LDAP User Groups

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/ldap-user-groups/{id}                     | Get LDAP User Group by id                 |
| GET    | /v0.4/ldap-user-groups                          | Get list of all LDAP User Groups          |
| PUT    | /v0.4/ldap-user-groups/{id}/user-roles/{id}     | Assign User Role to an LDAP User Group    |

### Radius Servers

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/radius-servers/{id}                       | Get Radius Server by id                   |
| GET    | /v0.4/radius-servers                            | Get list of all Radius Servers            |
| POST   | /v0.4/radius-servers                            | Create Radius Server                      |
| PUT    | /v0.4/radius-servers/{id}                       | Update Radius Server                      |
| DELETE | /v0.4/radius-servers/{id}                       | Delete Radius Server                      |

### Resource Groups

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/resource-groups/{id}                      | Get Resource Group by id                  |
| GET    | /v0.4/resource-groups                           | Get list of all Resource Groups           |
| POST   | /v0.4/resource-groups                           | Create Resource Group                     |
| PUT    | /v0.4/resource-groups/{id}                      | Update Resource Group                     |
| DELETE | /v0.4/resource-groups/{id}                      | Delete Resource Group                     |

### Security Groups

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/security-groups/{id}                      | Get Security Group by id                  |
| GET    | /v0.4/security-groups                           | Get list of all Security Groups           |
| PUT    | /v0.4/security-groups/{id}/resource-groups/{id} | Assign Resource Group to a Security Group |

### Servers

*NOTE: Deprecated

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/servers                                   | Get list of all servers                   |
| POST   | /v0.4/enrollment-scripts                        | Generate gateway install script           |

### Simple Server Policies

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/simple-server-policies/{id}               | Get Simple Server Policy by id            |
| GET    | /v0.4/simple-server-policies                    | Get list of all Simple Server Policies    |
| POST   | /v0.4/simple-server-policies                    | Create Simple Server Policy               |
| PUT    | /v0.4/simple-server-policies/{id}               | Update Simple Server Policy               |
| DELETE | /v0.4/simple-server-policies/{id}               | Delete Simple Server Policy               |

### Simple User Policies

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/simple-user-policies/{id}                 | Get Simple User Policy by id              |
| GET    | /v0.4/simple-user-policies                      | Get list of all Simple User Policies      |
| POST   | /v0.4/simple-user-policies                      | Create Simple User Policy                 |
| PUT    | /v0.4/simple-user-policies/{id}                 | Update Simple User Policy                 |
| DELETE | /v0.4/simple-user-policies/{id}                 | Delete Simple User Policy                 |

### SNAT Actions

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/snat-actions/{id}                         | Get SNAT Action by id                     |
| GET    | /v0.4/snat-actions                              | Get list of all SNAT Actions              |
| POST   | /v0.4/snat-actions                              | Create SNAT Action                        |
| PUT    | /v0.4/snat-actions/{id}                         | Update SNAT Action                        |
| DELETE | /v0.4/snat-actions/{id}                         | Delete SNAT Action                        |

### SNAT Rules

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/snat-rules/{id}                           | Get SNAT Rule by id                       |
| GET    | /v0.4/snat-rules                                | Get list of all SNAT Rules                |
| POST   | /v0.4/snat-rules                                | Create SNAT Rule                          |
| PUT    | /v0.4/snat-rules/{id}                           | Update SNAT Rule                          |
| PUT    | /v0.4/snat-rules/{id}/snat-actions/{id}         | Assign SNAT Action to SNAT Rule           |
| DELETE | /v0.4/snat-rules/{id}                           | Delete SNAT Rule                          |

### Subnets

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/subnets/{id}                              | Get Subnet by id                          |
| GET    | /v0.4/subnets                                   | Get list of all Subnets                   |
| POST   | /v0.4/subnets                                   | Create Subnet                             |
| PUT    | /v0.4/subnets/{id}                              | Update Subnet                             |
| PUT    | /v0.4/subnets/{id}/gateways/{id}                | Assign Gateway to Subnet                  |
| PUT    | /v0.4/subnets/{id}/resource-groups/{id}         | Assign Resource Group to Subnet           |
| DELETE | /v0.4/subnets/{id}                              | Delete Subnet                             |

### System Alerts

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/system-alerts/{id}                        | Get System Alert by id                    |
| GET    | /v0.4/system-alerts                             | Get list of all System Alerts             |

### System Events

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/system-events/{id}                        | Get System Event by id                    |
| GET    | /v0.4/system-events                             | Get list of all System Events             |

### Users

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/users/{id}                                | Get User by id                            |
| GET    | /v0.4/users                                     | Get list of all Users                     |
| POST   | /v0.4/users                                     | Create User                               |
| PUT    | /v0.4/users/{id}                                | Update User                               |
| DELETE | /v0.4/users/{id}                                | Delete User                               |

### User Roles

| Method | URI                                             | Description                               |
|--------|-------------------------------------------------|-------------------------------------------|
| GET    | /v0.4/user-roles/{id}                           | Get User Role by id                       |
| GET    | /v0.4/user-roles                                | Get list of all User Roles                |
| POST   | /v0.4/user-roles                                | Create User Role                          |
| PUT    | /v0.4/user-roles/{id}                           | Update User Role                          |
| DELETE | /v0.4/user-roles/{id}                           | Delete User Role                          |
