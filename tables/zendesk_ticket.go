package tables

import (
	"context"
	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-zendesk/table_schema_generator"
	"github.com/selefra/selefra-provider-zendesk/zendesk_client"
)

type TableZendeskTicketGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZendeskTicketGenerator{}

func (x *TableZendeskTicketGenerator) GetTableName() string {
	return "zendesk_ticket"
}

func (x *TableZendeskTicketGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZendeskTicketGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZendeskTicketGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZendeskTicketGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := zendesk_client.Connect(ctx, taskClient.(*zendesk_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := &zendesk.TicketListOptions{
				SortBy:    "created_at",
				SortOrder: "desc",
				PageOptions: zendesk.PageOptions{
					Page:    1,
					PerPage: 100,
				},
			}
			for {
				tickets, page, err := conn.GetTickets(ctx, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, t := range tickets {
					resultChannel <- t
				}
				if !page.HasNext() {
					break
				}
				opts.Page++
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableZendeskTicketGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZendeskTicketGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("allow_attachments").ColumnType(schema.ColumnTypeBool).Description("Permission for agents to add add attachments to a comment. Defaults to true").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfaction_rating_score").ColumnType(schema.ColumnTypeString).Description("The rating \"offered\", \"unoffered\", \"good\" or \"bad\"").
			Extractor(column_value_extractor.StructSelector("SatisfactionRating.Score")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_source_from").ColumnType(schema.ColumnTypeJSON).Description("Source the ticket was sent to").
			Extractor(column_value_extractor.StructSelector("Via.Source.From")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_source_ref").ColumnType(schema.ColumnTypeString).Description("Medium used to raise the ticket").
			Extractor(column_value_extractor.StructSelector("Via.Source.Ref")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("The array of tags applied to this ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("brand_id").ColumnType(schema.ColumnTypeInt).Description("Enterprise only. The id of the brand this ticket is associated with").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When this record was created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_cc_ids").ColumnType(schema.ColumnTypeJSON).Description("The ids of agents or end users currently CC'ed on the ticket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_id").ColumnType(schema.ColumnTypeInt).Description("The group this ticket is assigned to").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_public").ColumnType(schema.ColumnTypeBool).Description("Is true if any comments are public, false otherwise").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requester_id").ColumnType(schema.ColumnTypeInt).Description("The user who requested this ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfaction_rating_id").ColumnType(schema.ColumnTypeInt).Description("Unique identifier for the satisfaction rating on this ticket").
			Extractor(column_value_extractor.StructSelector("SatisfactionRating.ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("The API url of this ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_followup_source_id").ColumnType(schema.ColumnTypeString).Description("The id of a closed ticket when creating a follow-up ticket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_channelback").ColumnType(schema.ColumnTypeBool).Description("Is false if channelback is disabled, true otherwise. Only applicable for channels framework ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_fields").ColumnType(schema.ColumnTypeJSON).Description("Custom fields for the ticket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assignee_id").ColumnType(schema.ColumnTypeInt).Description("The agent currently assigned to the ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("follower_ids").ColumnType(schema.ColumnTypeJSON).Description("The ids of agents currently following the ticket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_incidents").ColumnType(schema.ColumnTypeBool).Description("Is true if a ticket is a problem type and has one or more incidents linked to it. Otherwise, the value is false.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("Automatically assigned when the ticket is created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("macro_ids").ColumnType(schema.ColumnTypeJSON).Description("POST requests only. List of macro IDs to be recorded in the ticket audit").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When this record last got updated").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("due_at").ColumnType(schema.ColumnTypeTimestamp).Description("If this is a ticket of type \"task\" it has a due date. Due date format uses ISO 8601 format.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_id").ColumnType(schema.ColumnTypeInt).Description("The organization of the requester. You can only specify the ID of an organization associated with the requester.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subject").ColumnType(schema.ColumnTypeString).Description("The value of the subject field for this ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_channel").ColumnType(schema.ColumnTypeString).Description("How the ticket or event was created. Examples: \"web\", \"mobile\", \"rule\", \"system\"").
			Extractor(column_value_extractor.StructSelector("Via.Channel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collaborator_ids").ColumnType(schema.ColumnTypeJSON).Description("The ids of users currently CC'ed on the ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("raw_subject").ColumnType(schema.ColumnTypeString).Description("The dynamic content placeholder, if present, or the \"subject\" value, if not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recipient").ColumnType(schema.ColumnTypeString).Description("The original recipient e-mail address of the ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sharing_agreement_ids").ColumnType(schema.ColumnTypeJSON).Description("The ids of the sharing agreements used for this ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ticket_form_id").ColumnType(schema.ColumnTypeInt).Description("Enterprise only. The id of the ticket form to render for the ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of this ticket. Allowed values are \"problem\", \"incident\", \"question\", or \"task\".").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Read-only first comment on the ticket. When creating a ticket, use comment to set the description.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_id").ColumnType(schema.ColumnTypeInt).Description("An id you can use to link Zendesk Support tickets to local records").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("followup_ids").ColumnType(schema.ColumnTypeJSON).Description("The ids of the followups created from this ticket. Ids are only visible once the ticket is closed").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeString).Description("The urgency with which the ticket should be addressed. Allowed values are \"urgent\", \"high\", \"normal\", or \"low\".").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfaction_rating_comment").ColumnType(schema.ColumnTypeString).Description("The comment received with this rating, if available").
			Extractor(column_value_extractor.StructSelector("SatisfactionRating.Comment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The state of the ticket. Allowed values are \"new\", \"open\", \"pending\", \"hold\", \"solved\", or \"closed\".").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("submitter_id").ColumnType(schema.ColumnTypeInt).Description("The user who submitted the ticket. The submitter always becomes the author of the first comment on the ticket").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("forum_topic_id").ColumnType(schema.ColumnTypeInt).Description("The topic in the Zendesk Web portal this ticket originated from, if any. The Web portal is deprecated").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("problem_id").ColumnType(schema.ColumnTypeInt).Description("For tickets of type \"incident\", the ID of the problem the incident is linked to").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_source_to").ColumnType(schema.ColumnTypeJSON).Description("Target that received the ticket").
			Extractor(column_value_extractor.StructSelector("Via.Source.From")).Build(),
	}
}

func (x *TableZendeskTicketGenerator) GetSubTables() []*schema.Table {
	return nil
}
