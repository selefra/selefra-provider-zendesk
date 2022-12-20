# Table: zendesk_trigger

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| conditions_any | json | X | √ | Trigger if any condition is met. | 
| description | string | X | √ | The description of the trigger | 
| updated_at | timestamp | X | √ | The time of the last update of the trigger | 
| id | int | X | √ | Automatically assigned when the trigger is created | 
| title | string | X | √ | The title of the trigger | 
| active | bool | X | √ | Whether the trigger is active | 
| actions | json | X | √ | An array of actions describing what the trigger will do. | 
| conditions_all | json | X | √ | Trigger if all conditions are met. | 
| created_at | timestamp | X | √ | The time the trigger was created | 
| position | int | X | √ | Position of the trigger, determines the order they will execute in. | 


