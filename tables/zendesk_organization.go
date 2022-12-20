package tables

import (
	"context"
	"github.com/selefra/selefra-provider-zendesk/zendesk_client"

	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-zendesk/table_schema_generator"
)

type TableZendeskOrganizationGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZendeskOrganizationGenerator{}

func (x *TableZendeskOrganizationGenerator) GetTableName() string {
	return "zendesk_organization"
}

func (x *TableZendeskOrganizationGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZendeskOrganizationGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZendeskOrganizationGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZendeskOrganizationGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := zendesk_client.Connect(ctx, taskClient.(*zendesk_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := &zendesk.OrganizationListOptions{
				PageOptions: zendesk.PageOptions{
					Page:    1,
					PerPage: 100,
				},
			}
			for {
				organizations, page, err := conn.GetOrganizations(ctx, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, t := range organizations {
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

func (x *TableZendeskOrganizationGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZendeskOrganizationGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A unique name for the organization").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shared_comments").ColumnType(schema.ColumnTypeBool).Description("End users in this organization are able to see each other's comments on tickets").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shared_tickets").ColumnType(schema.ColumnTypeBool).Description("End users in this organization are able to see each other's tickets").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_fields").ColumnType(schema.ColumnTypeJSON).Description("Custom fields for this organization").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("The tags of the organization").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time of the last update of the organization").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("The API url of this organization").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time the organization was created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_names").ColumnType(schema.ColumnTypeJSON).Description("An array of domain names associated with this organization").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_id").ColumnType(schema.ColumnTypeInt).Description("New tickets from users in this organization are automatically put in this group").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("Automatically assigned when the organization is created").Build(),
	}
}

func (x *TableZendeskOrganizationGenerator) GetSubTables() []*schema.Table {
	return nil
}
