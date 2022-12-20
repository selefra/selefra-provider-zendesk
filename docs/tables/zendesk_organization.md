# Table: zendesk_organization

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | A unique name for the organization | 
| shared_comments | bool | X | √ | End users in this organization are able to see each other's comments on tickets | 
| shared_tickets | bool | X | √ | End users in this organization are able to see each other's tickets | 
| organization_fields | json | X | √ | Custom fields for this organization | 
| tags | json | X | √ | The tags of the organization | 
| updated_at | timestamp | X | √ | The time of the last update of the organization | 
| url | string | X | √ | The API url of this organization | 
| created_at | timestamp | X | √ | The time the organization was created | 
| domain_names | json | X | √ | An array of domain names associated with this organization | 
| group_id | int | X | √ | New tickets from users in this organization are automatically put in this group | 
| id | int | X | √ | Automatically assigned when the organization is created | 


