# Table: zendesk_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| custom_role_id | int | X | √ | A custom role if the user is an agent on the Enterprise plan | 
| name | string | X | √ | The user's name | 
| only_private_comments | bool | X | √ | true if the user can only create private comments | 
| report_csv | bool | X | √ | Whether or not the user can access the CSV report on the Search tab of the Reporting page in the Support admin interface. | 
| active | bool | X | √ | False if the user has been deleted | 
| role | string | X | √ | The user's role. Possible values are end-user, agent, or admin | 
| shared | bool | X | √ | If the user is shared from a different Zendesk Support instance. Ticket sharing accounts only | 
| updated_at | timestamp | X | √ | The time the user was last updated | 
| user_fields | json | X | √ | Values of custom fields in the user's profile. | 
| id | int | X | √ | Automatically assigned when the user is created | 
| photo_inline | bool | X | √ | If true, the attachment is excluded from the attachment list and the attachment's URL can be referenced within the comment of a ticket. Default is false | 
| restricted_agent | bool | X | √ | If the agent has any restrictions; false for admins and unrestricted agents, true for other agents | 
| suspended | bool | X | √ | If the agent is suspended. Tickets from suspended users are also suspended, and these users cannot sign in to the end user portal | 
| photo_thumbnails | json | X | √ | An array of attachment objects. Note that photo thumbnails do not have thumbnails | 
| ticket_restriction | string | X | √ | Specifies which tickets the user has access to. Possible values are: "organization", "groups", "assigned", "requested", null | 
| alias | string | X | √ | An alias displayed to end users | 
| default_group_id | int | X | √ | The id of the user's default group | 
| locale | string | X | √ | The user's locale. A BCP-47 compliant tag for the locale. If both "locale" and "locale_id" are present on create or update, "locale_id" is ignored and only "locale" is used. | 
| locale_id | int | X | √ | The user's language identifier | 
| photo_file_name | string | X | √ | The name of the image file | 
| photo_size | int | X | √ | The size of the image file in bytes | 
| external_id | string | X | √ | A unique identifier from another system. The API treats the id as case insensitive. Example: "ian1" and "Ian1" are the same user | 
| last_login_at | timestamp | X | √ | The last time the user signed in to Zendesk Support | 
| photo_content_url | string | X | √ | A full URL where the attachment image file can be downloaded | 
| verified | bool | X | √ | Any of the user's identities is verified. | 
| details | string | X | √ | Any details you want to store about the user, such as an address | 
| moderator | bool | X | √ | Designates whether the user has forum moderation capabilities | 
| notes | string | X | √ | Any notes you want to store about the user | 
| signature | string | X | √ | The user's signature. Only agents and admins can have signatures | 
| tags | json | X | √ | The user's tags. Only present if your account has user tagging enabled | 
| two_factor_auth_enabled | bool | X | √ | If two factor authentication is enabled | 
| photo_deleted | string | X | √ | If true, the attachment has been deleted | 
| shared_agent | bool | X | √ | If the user is a shared agent from a different Zendesk Support instance. Ticket sharing accounts only | 
| shared_phone_number | bool | X | √ | Whether the phone number is shared or not. | 
| photo_id | int | X | √ | Automatically assigned when created | 
| role_type | int | X | √ | The user's role id. 0 for custom agents, 1 for light agent, 2 for chat agent, and 3 for chat agent added to the Support account as a contributor (Chat Phase 4) | 
| chat_only | bool | X | √ | Whether or not the user is a chat-only agent | 
| created_at | timestamp | X | √ | The time the user was created | 
| email | string | X | √ | The user's primary email address. *Writeable on create only. On update, a secondary email is added. | 
| organization_id | int | X | √ | The id of the user's organization. If the user has more than one organization memberships, the id of the user's default organization | 
| phone | string | X | √ | The user's primary phone number. | 
| photo_content_type | string | X | √ | The content type of the image. Example value: "image/png" | 
| timezone | string | X | √ | The user's time zone. | 
| url | string | X | √ | The user's API url | 


