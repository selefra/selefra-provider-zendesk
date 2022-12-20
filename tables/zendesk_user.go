package tables

import (
	"context"
	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-zendesk/table_schema_generator"
	"github.com/selefra/selefra-provider-zendesk/zendesk_client"
)

type TableZendeskUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableZendeskUserGenerator{}

func (x *TableZendeskUserGenerator) GetTableName() string {
	return "zendesk_user"
}

func (x *TableZendeskUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableZendeskUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableZendeskUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableZendeskUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := zendesk_client.Connect(ctx, taskClient.(*zendesk_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := &zendesk.UserListOptions{
				PageOptions: zendesk.PageOptions{
					Page:    1,
					PerPage: 100,
				},
			}
			for {
				users, page, err := conn.GetUsers(ctx, opts)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, t := range users {
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

func (x *TableZendeskUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableZendeskUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("custom_role_id").ColumnType(schema.ColumnTypeInt).Description("A custom role if the user is an agent on the Enterprise plan").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The user's name").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("only_private_comments").ColumnType(schema.ColumnTypeBool).Description("true if the user can only create private comments").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("report_csv").ColumnType(schema.ColumnTypeBool).Description("Whether or not the user can access the CSV report on the Search tab of the Reporting page in the Support admin interface.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active").ColumnType(schema.ColumnTypeBool).Description("False if the user has been deleted").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).Description("The user's role. Possible values are end-user, agent, or admin").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shared").ColumnType(schema.ColumnTypeBool).Description("If the user is shared from a different Zendesk Support instance. Ticket sharing accounts only").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time the user was last updated").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_fields").ColumnType(schema.ColumnTypeJSON).Description("Values of custom fields in the user's profile.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).Description("Automatically assigned when the user is created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photo_inline").ColumnType(schema.ColumnTypeBool).Description("If true, the attachment is excluded from the attachment list and the attachment's URL can be referenced within the comment of a ticket. Default is false").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restricted_agent").ColumnType(schema.ColumnTypeBool).Description("If the agent has any restrictions; false for admins and unrestricted agents, true for other agents").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("suspended").ColumnType(schema.ColumnTypeBool).Description("If the agent is suspended. Tickets from suspended users are also suspended, and these users cannot sign in to the end user portal").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photo_thumbnails").ColumnType(schema.ColumnTypeJSON).Description("An array of attachment objects. Note that photo thumbnails do not have thumbnails").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ticket_restriction").ColumnType(schema.ColumnTypeString).Description("Specifies which tickets the user has access to. Possible values are: \"organization\", \"groups\", \"assigned\", \"requested\", null").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alias").ColumnType(schema.ColumnTypeString).Description("An alias displayed to end users").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_group_id").ColumnType(schema.ColumnTypeInt).Description("The id of the user's default group").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locale").ColumnType(schema.ColumnTypeString).Description("The user's locale. A BCP-47 compliant tag for the locale. If both \"locale\" and \"locale_id\" are present on create or update, \"locale_id\" is ignored and only \"locale\" is used.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locale_id").ColumnType(schema.ColumnTypeInt).Description("The user's language identifier").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photo_file_name").ColumnType(schema.ColumnTypeString).Description("The name of the image file").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photo_size").ColumnType(schema.ColumnTypeInt).Description("The size of the image file in bytes").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_id").ColumnType(schema.ColumnTypeString).Description("A unique identifier from another system. The API treats the id as case insensitive. Example: \"ian1\" and \"Ian1\" are the same user").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_login_at").ColumnType(schema.ColumnTypeTimestamp).Description("The last time the user signed in to Zendesk Support").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photo_content_url").ColumnType(schema.ColumnTypeString).Description("A full URL where the attachment image file can be downloaded").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("verified").ColumnType(schema.ColumnTypeBool).Description("Any of the user's identities is verified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("details").ColumnType(schema.ColumnTypeString).Description("Any details you want to store about the user, such as an address").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("moderator").ColumnType(schema.ColumnTypeBool).Description("Designates whether the user has forum moderation capabilities").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notes").ColumnType(schema.ColumnTypeString).Description("Any notes you want to store about the user").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("signature").ColumnType(schema.ColumnTypeString).Description("The user's signature. Only agents and admins can have signatures").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("The user's tags. Only present if your account has user tagging enabled").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("two_factor_auth_enabled").ColumnType(schema.ColumnTypeBool).Description("If two factor authentication is enabled").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photo_deleted").ColumnType(schema.ColumnTypeString).Description("If true, the attachment has been deleted").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shared_agent").ColumnType(schema.ColumnTypeBool).Description("If the user is a shared agent from a different Zendesk Support instance. Ticket sharing accounts only").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shared_phone_number").ColumnType(schema.ColumnTypeBool).Description("Whether the phone number is shared or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photo_id").ColumnType(schema.ColumnTypeInt).Description("Automatically assigned when created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_type").ColumnType(schema.ColumnTypeInt).Description("The user's role id. 0 for custom agents, 1 for light agent, 2 for chat agent, and 3 for chat agent added to the Support account as a contributor (Chat Phase 4)").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("chat_only").ColumnType(schema.ColumnTypeBool).Description("Whether or not the user is a chat-only agent").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The time the user was created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("The user's primary email address. *Writeable on create only. On update, a secondary email is added.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_id").ColumnType(schema.ColumnTypeInt).Description("The id of the user's organization. If the user has more than one organization memberships, the id of the user's default organization").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phone").ColumnType(schema.ColumnTypeString).Description("The user's primary phone number.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("photo_content_type").ColumnType(schema.ColumnTypeString).Description("The content type of the image. Example value: \"image/png\"").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timezone").ColumnType(schema.ColumnTypeString).Description("The user's time zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Description("The user's API url").Build(),
	}
}

func (x *TableZendeskUserGenerator) GetSubTables() []*schema.Table {
	return nil
}
