package tables

import (
	"context"
	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-zendesk/table_schema_generator"
	"github.com/selefra/selefra-provider-zendesk/zendesk_client"
)

type TableZendeskTriggerGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZendeskTriggerGenerator{}

func (x *TableZendeskTriggerGenerator) GetTableName() string {
	return "zendesk_trigger"
}

func (x *TableZendeskTriggerGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZendeskTriggerGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZendeskTriggerGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZendeskTriggerGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := zendesk_client.Connect(ctx, taskClient.(*zendesk_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := &zendesk.TriggerListOptions{
				PageOptions: zendesk.PageOptions{
					Page:    1,
					PerPage: 100,
				},
			}
			for {
				triggers, page, err := conn.GetTriggers(ctx, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, t := range triggers {
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

func (x *TableZendeskTriggerGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZendeskTriggerGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("conditions_any").ColumnType(schema.ColumnTypeJSON).Description("Trigger if any condition is met.").
			Extractor(column_value_extractor.StructSelector("Conditions.Any")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the trigger").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time of the last update of the trigger").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("Automatically assigned when the trigger is created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the trigger").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active").ColumnType(schema.ColumnTypeBool).Description("Whether the trigger is active").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actions").ColumnType(schema.ColumnTypeJSON).Description("An array of actions describing what the trigger will do.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conditions_all").ColumnType(schema.ColumnTypeJSON).Description("Trigger if all conditions are met.").
			Extractor(column_value_extractor.StructSelector("Conditions.All")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time the trigger was created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("position").ColumnType(schema.ColumnTypeInt).Description("Position of the trigger, determines the order they will execute in.").Build(),
	}
}

func (x *TableZendeskTriggerGenerator) GetSubTables() []*schema.Table {
	return nil
}
