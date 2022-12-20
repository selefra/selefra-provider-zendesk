package provider

import (
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-zendesk/table_schema_generator"
	"github.com/selefra/selefra-provider-zendesk/tables"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableZendeskGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZendeskOrganizationGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZendeskTicketGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZendeskTicketAuditGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZendeskTriggerGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableZendeskUserGenerator{}),
	}
}
