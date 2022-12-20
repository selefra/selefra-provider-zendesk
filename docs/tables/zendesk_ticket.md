# Table: zendesk_ticket

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| allow_attachments | bool | X | √ | Permission for agents to add add attachments to a comment. Defaults to true | 
| satisfaction_rating_score | string | X | √ | The rating "offered", "unoffered", "good" or "bad" | 
| via_source_from | json | X | √ | Source the ticket was sent to | 
| via_source_ref | string | X | √ | Medium used to raise the ticket | 
| tags | json | X | √ | The array of tags applied to this ticket | 
| brand_id | int | X | √ | Enterprise only. The id of the brand this ticket is associated with | 
| created_at | timestamp | X | √ | When this record was created | 
| email_cc_ids | json | X | √ | The ids of agents or end users currently CC'ed on the ticket. | 
| group_id | int | X | √ | The group this ticket is assigned to | 
| is_public | bool | X | √ | Is true if any comments are public, false otherwise | 
| requester_id | int | X | √ | The user who requested this ticket | 
| satisfaction_rating_id | int | X | √ | Unique identifier for the satisfaction rating on this ticket | 
| url | string | X | √ | The API url of this ticket | 
| via_followup_source_id | string | X | √ | The id of a closed ticket when creating a follow-up ticket. | 
| allow_channelback | bool | X | √ | Is false if channelback is disabled, true otherwise. Only applicable for channels framework ticket | 
| custom_fields | json | X | √ | Custom fields for the ticket. | 
| assignee_id | int | X | √ | The agent currently assigned to the ticket | 
| follower_ids | json | X | √ | The ids of agents currently following the ticket. | 
| has_incidents | bool | X | √ | Is true if a ticket is a problem type and has one or more incidents linked to it. Otherwise, the value is false. | 
| id | int | X | √ | Automatically assigned when the ticket is created | 
| macro_ids | json | X | √ | POST requests only. List of macro IDs to be recorded in the ticket audit | 
| updated_at | timestamp | X | √ | When this record last got updated | 
| due_at | timestamp | X | √ | If this is a ticket of type "task" it has a due date. Due date format uses ISO 8601 format. | 
| organization_id | int | X | √ | The organization of the requester. You can only specify the ID of an organization associated with the requester. | 
| subject | string | X | √ | The value of the subject field for this ticket | 
| via_channel | string | X | √ | How the ticket or event was created. Examples: "web", "mobile", "rule", "system" | 
| collaborator_ids | json | X | √ | The ids of users currently CC'ed on the ticket | 
| raw_subject | string | X | √ | The dynamic content placeholder, if present, or the "subject" value, if not. | 
| recipient | string | X | √ | The original recipient e-mail address of the ticket | 
| sharing_agreement_ids | json | X | √ | The ids of the sharing agreements used for this ticket | 
| ticket_form_id | int | X | √ | Enterprise only. The id of the ticket form to render for the ticket | 
| type | string | X | √ | The type of this ticket. Allowed values are "problem", "incident", "question", or "task". | 
| description | string | X | √ | Read-only first comment on the ticket. When creating a ticket, use comment to set the description. | 
| external_id | int | X | √ | An id you can use to link Zendesk Support tickets to local records | 
| followup_ids | json | X | √ | The ids of the followups created from this ticket. Ids are only visible once the ticket is closed | 
| priority | string | X | √ | The urgency with which the ticket should be addressed. Allowed values are "urgent", "high", "normal", or "low". | 
| satisfaction_rating_comment | string | X | √ | The comment received with this rating, if available | 
| status | string | X | √ | The state of the ticket. Allowed values are "new", "open", "pending", "hold", "solved", or "closed". | 
| submitter_id | int | X | √ | The user who submitted the ticket. The submitter always becomes the author of the first comment on the ticket | 
| forum_topic_id | int | X | √ | The topic in the Zendesk Web portal this ticket originated from, if any. The Web portal is deprecated | 
| problem_id | int | X | √ | For tickets of type "incident", the ID of the problem the incident is linked to | 
| via_source_to | json | X | √ | Target that received the ticket | 


