package tables

import (
	"context"
	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-zendesk/table_schema_generator"
	"github.com/selefra/selefra-provider-zendesk/zendesk_client"
)

type TableZendeskTicketAuditGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZendeskTicketAuditGenerator{}

func (x *TableZendeskTicketAuditGenerator) GetTableName() string {
	return "zendesk_ticket_audit"
}

func (x *TableZendeskTicketAuditGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZendeskTicketAuditGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZendeskTicketAuditGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZendeskTicketAuditGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := zendesk_client.Connect(ctx, taskClient.(*zendesk_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := zendesk.CursorOption{}
			for {
				ticketAudits, cursor, err := conn.GetAllTicketAudits(ctx, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, t := range ticketAudits {
					resultChannel <- t
				}
				opts.Cursor = cursor.AfterCursor
				if cursor.AfterCursor == "" {
					break
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableZendeskTicketAuditGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZendeskTicketAuditGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("ticket_id").ColumnType(schema.ColumnTypeInt).Description("The ID of the associated ticket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("author_id").ColumnType(schema.ColumnTypeInt).Description("The user who created the audit.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_channel").ColumnType(schema.ColumnTypeString).Description("How the ticket or event was created. Examples: \"web\", \"mobile\", \"rule\", \"system\".").
			Extractor(column_value_extractor.StructSelector("Via.Channel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_source_from").ColumnType(schema.ColumnTypeJSON).Description("Source the ticket was sent to.").
			Extractor(column_value_extractor.StructSelector("Via.Source.From")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("Unique identifier for the ticket update.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time the audit was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("events").ColumnType(schema.ColumnTypeJSON).Description("An array of the events that happened in this audit.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Metadata for the audit, custom and system data.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_followup_source_id").ColumnType(schema.ColumnTypeString).Description("The id of a closed ticket when creating a follow-up ticket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_source_ref").ColumnType(schema.ColumnTypeString).Description("Medium used to raise the ticket.").
			Extractor(column_value_extractor.StructSelector("Via.Source.Ref")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("via_source_to").ColumnType(schema.ColumnTypeJSON).Description("Target that received the ticket.").
			Extractor(column_value_extractor.StructSelector("Via.Source.From")).Build(),
	}
}

func (x *TableZendeskTicketAuditGenerator) GetSubTables() []*schema.Table {
	return nil
}
