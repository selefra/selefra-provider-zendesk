package tables

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-zendesk/table_schema_generator"
	"github.com/selefra/selefra-provider-zendesk/zendesk_client"
)

type TableZendeskGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZendeskGroupGenerator{}

func (x *TableZendeskGroupGenerator) GetTableName() string {
	return "zendesk_group"
}

func (x *TableZendeskGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZendeskGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZendeskGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZendeskGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := zendesk_client.Connect(ctx, taskClient.(*zendesk_client.Client).Config)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			groups, _, err := conn.GetGroups(ctx, nil)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, t := range groups {
				resultChannel <- t
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableZendeskGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZendeskGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time the group was created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time of the last update of the group").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("Unique identifier for the group").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("API url of the group").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the group").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Description("True if the group has been deleted").Build(),
	}
}

func (x *TableZendeskGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}
