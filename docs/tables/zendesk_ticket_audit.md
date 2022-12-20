# Table: zendesk_ticket_audit

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ticket_id | int | X | √ | The ID of the associated ticket. | 
| author_id | int | X | √ | The user who created the audit. | 
| via_channel | string | X | √ | How the ticket or event was created. Examples: "web", "mobile", "rule", "system". | 
| via_source_from | json | X | √ | Source the ticket was sent to. | 
| id | int | X | √ | Unique identifier for the ticket update. | 
| created_at | timestamp | X | √ | The time the audit was created. | 
| events | json | X | √ | An array of the events that happened in this audit. | 
| metadata | json | X | √ | Metadata for the audit, custom and system data. | 
| via_followup_source_id | string | X | √ | The id of a closed ticket when creating a follow-up ticket. | 
| via_source_ref | string | X | √ | Medium used to raise the ticket. | 
| via_source_to | json | X | √ | Target that received the ticket. | 


