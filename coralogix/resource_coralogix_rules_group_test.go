package coralogix

import (
	"context"
	"fmt"
	"testing"

	"terraform-provider-coralogix/coralogix/clientset"
	rulesv1 "terraform-provider-coralogix/coralogix/clientset/grpc/com/coralogix/rules/v1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

/*
func TestAccCoralogixResourceRuleGroup_minimal(t *testing.T) {
	name := acctest.RandomWithPrefix("tf-acc-test")
	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupMinimal(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				Config:   testAccCoralogixResourceRuleGroupMinimal(name),
				PlanOnly: true,
			},
		},
	})
}*/

func TestAccCoralogixResourceRuleGroup_block(t *testing.T) {
	r := getRandomRuleGroup()

	keepBlockedLogs := "true"
	regEx := `sql_error_code\\s*=\\s*28000`

	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupBlock(r, regEx, keepBlockedLogs),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttrSet(resourceName, "rule_subgroups.0.rules.0.block.0.id"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.active", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.name", r.ruleParams.name),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.description", r.ruleParams.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.source_field", "text"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.keep_blocked_logs", keepBlockedLogs),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.blocking_all_matching_blocks", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.regular_expression", "sql_error_code\\s*=\\s*28000"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_allow(t *testing.T) {
	r := getRandomRuleGroup()

	keepBlockedLogs := "true"
	regEx := `sql_error_code\\s*=\\s*28000`

	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupAllow(r, regEx, keepBlockedLogs),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttrSet(resourceName, "rule_subgroups.0.rules.0.block.0.id"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.active", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.name", r.ruleParams.name),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.description", r.ruleParams.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.keep_blocked_logs", keepBlockedLogs),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.blocking_all_matching_blocks", "false"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.block.0.regular_expression", "sql_error_code\\s*=\\s*28000"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_jsonExtract(t *testing.T) {
	r := getRandomRuleGroup()

	jsonKey := "worker"
	destinationField := "Category"

	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupJsonExtract(r, jsonKey, destinationField),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttrSet(resourceName, "rule_subgroups.0.rules.0.json_extract.0.id"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_extract.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_extract.0.active", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_extract.0.name", r.ruleParams.name),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_extract.0.description", r.ruleParams.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_extract.0.destination_field", destinationField),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_extract.0.json_key", jsonKey),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_replace(t *testing.T) {
	r := getRandomRuleGroup()

	regEx := ".*{"
	replacementString := "{"

	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupReplace(r, regEx, replacementString),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttrSet(resourceName, "rule_subgroups.0.rules.0.replace.0.id"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.replace.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.replace.0.active", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.replace.0.name", r.ruleParams.name),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.replace.0.description", r.ruleParams.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.replace.0.source_field", "text"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.replace.0.destination_field", "text"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.replace.0.regular_expression", regEx),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.replace.0.replacement_string", replacementString),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_extractTimestamp(t *testing.T) {
	r := getRandomRuleGroup()

	timeFormat := ""

	fieldFormatStandard := "NanoTS"

	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupExtractTimestamp(r, timeFormat, fieldFormatStandard),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttrSet(resourceName, "rule_subgroups.0.rules.0.extract_timestamp.0.id"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract_timestamp.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract_timestamp.0.active", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract_timestamp.0.name", r.ruleParams.name),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract_timestamp.0.description", r.ruleParams.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract_timestamp.0.time_format", timeFormat),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract_timestamp.0.field_format_standard", fieldFormatStandard),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_removeFields(t *testing.T) {
	r := getRandomRuleGroup()

	excludedFields := `["coralogix.metadata.applicationName", "coralogix.metadata.className"]`

	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupRemoveFields(r, excludedFields),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttrSet(resourceName, "rule_subgroups.0.rules.0.remove_fields.0.id"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.remove_fields.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.remove_fields.0.active", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.remove_fields.0.name", r.ruleParams.name),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.remove_fields.0.description", r.ruleParams.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.remove_fields.0.excluded_fields.0", "coralogix.metadata.applicationName"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.remove_fields.0.excluded_fields.1", "coralogix.metadata.className"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_jsonStringify(t *testing.T) {
	r := getRandomRuleGroup()

	keepSourceField := "true"

	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupJsonStringify(r, keepSourceField),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttrSet(resourceName, "rule_subgroups.0.rules.0.json_stringify.0.id"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_stringify.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_stringify.0.active", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_stringify.0.name", r.ruleParams.name),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_stringify.0.description", r.ruleParams.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_stringify.0.source_field", "text"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_stringify.0.destination_field", "text"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.json_stringify.0.keep_source_field", keepSourceField),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_extract(t *testing.T) {
	r := getRandomRuleGroup()

	regEx := `\\b(?P<severity>DEBUG|TRACE|INFO|WARN|ERROR|FATAL|EXCEPTION|[Dd]ebug|[Tt]race|[Ii]nfo|[Ww]arn|[Ee]rror|[Ff]atal|[Ee]xception)\\b`

	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupExtract(r, regEx),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttrSet(resourceName, "rule_subgroups.0.rules.0.extract.0.id"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract.0.active", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract.0.name", r.ruleParams.name),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract.0.description", r.ruleParams.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract.0.source_field", "text"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract.0.regular_expression",
						"\\b(?P<severity>DEBUG|TRACE|INFO|WARN|ERROR|FATAL|EXCEPTION|[Dd]ebug|[Tt]race|[Ii]nfo|[Ww]arn|[Ee]rror|[Ff]atal|[Ee]xception)\\b"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_parse(t *testing.T) {
	r := getRandomRuleGroup()

	regEx := `(?P<remote_addr>\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3})\\s*-\\s*(?P<user>[^ ]+)\\s*\\[(?P<timestemp>\\d{4}-\\d{2}\\-\\d{2}T\\d{2}\\:\\d{2}\\:\\d{2}\\.\\d{1,6}Z)\\]\\s*\\\\\\\"(?P<method>[A-z]+)\\s[\\/\\\\]+(?P<request>[^\\s]+)\\s*(?P<protocol>[A-z0-9\\/\\.]+)\\\\\\\"\\s*(?P<status>\\d+)\\s*(?P<body_bytes_sent>\\d+)?\\s*?\\\\\\\"(?P<http_referer>[^\"]+)\\\"\\s*\\\\\\\"(?P<http_user_agent>[^\"]+)\\\"\\s(?P<request_time>\\d{1,6})\\s*(?P<response_time>\\d{1,6})`

	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleGroupParse(r, regEx),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttrSet(resourceName, "rule_subgroups.0.rules.0.parse.0.id"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.active", "true"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.name", r.ruleParams.name),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.description", r.ruleParams.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.source_field", "text"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.destination_field", "text"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.regular_expression",
						"(?P<remote_addr>\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3})\\s*-\\s*(?P<user>[^ ]+)\\s*\\[(?P<timestemp>\\d{4}-\\d{2}\\-\\d{2}T\\d{2}\\:\\d{2}\\:\\d{2}\\.\\d{1,6}Z)\\]\\s*\\\\\\\"(?P<method>[A-z]+)\\s[\\/\\\\]+(?P<request>[^\\s]+)\\s*(?P<protocol>[A-z0-9\\/\\.]+)\\\\\\\"\\s*(?P<status>\\d+)\\s*(?P<body_bytes_sent>\\d+)?\\s*?\\\\\\\"(?P<http_referer>[^\"]+)\\\"\\s*\\\\\\\"(?P<http_user_agent>[^\"]+)\\\"\\s(?P<request_time>\\d{1,6})\\s*(?P<response_time>\\d{1,6})"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_rules_combination(t *testing.T) {
	r := getRandomRuleGroup()
	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleRulesCombination(r),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.name", "rule1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.0.name", "rule2"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.0.order", "2"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.0.name", "rule3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.0.order", "3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.0.name", "rule1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_update(t *testing.T) {
	r1 := getRandomRuleGroup()
	r2 := getRandomRuleGroup()
	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleRulesCombination(r1),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r1.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r1.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r1.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.name", "rule1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.0.name", "rule2"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.0.name", "rule3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.0.name", "rule1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCoralogixResourceRuleRulesCombination(r2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r2.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r2.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r2.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.name", "rule1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.0.name", "rule2"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.0.name", "rule3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.0.name", "rule1"),
				),
			},
		},
	})
}

func TestAccCoralogixResourceRuleGroup_update_order_inside_rule_group(t *testing.T) {
	r := getRandomRuleGroup()
	resourceName := "coralogix_rules_group.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckRuleGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixResourceRuleRulesCombination(r),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.name", "rule1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.parse.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.0.name", "rule2"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.extract.0.order", "2"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.0.name", "rule3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.0.order", "3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.0.name", "rule1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCoralogixResourceRuleRulesCombinationDifferentOrders(r),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "order"),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
					resource.TestCheckResourceAttr(resourceName, "hidden", "false"),
					resource.TestCheckResourceAttr(resourceName, "name", r.name),
					resource.TestCheckResourceAttr(resourceName, "creator", r.creator),
					resource.TestCheckResourceAttr(resourceName, "description", r.description),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract.0.name", "rule2"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.0.extract.0.order", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.parse.0.name", "rule3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.1.parse.0.order", "2"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.0.name", "rule1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.0.rules.2.parse.0.order", "3"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rule_subgroups.1.rules.0.extract_timestamp.0.name", "rule1"),
				),
			},
		},
	})
}

func getRandomRuleGroup() *ruleGroupParams {
	return &ruleGroupParams{
		name:        acctest.RandomWithPrefix("tf-acc-test"),
		description: acctest.RandomWithPrefix("tf-acc-test"),
		creator:     acctest.RandomWithPrefix("tf-acc-test"),
		ruleParams: ruleParams{
			name:        acctest.RandomWithPrefix("tf-acc-test"),
			description: acctest.RandomWithPrefix("tf-acc-test"),
		},
	}
}

func testAccCheckRuleGroupDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*clientset.ClientSet).RuleGroups()

	ctx := context.TODO()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "coralogix_rules_group" {
			continue
		}

		req := &rulesv1.GetRuleGroupRequest{
			GroupId: rs.Primary.ID,
		}

		resp, err := client.GetRuleGroup(ctx, req)
		if err == nil {
			if resp.RuleGroup.Id.Value == rs.Primary.ID {
				return fmt.Errorf("RuleGroup still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

/*func testAccCoralogixResourceRuleGroupMinimal(name string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "rule group from terraform provider"
 }
`, name)
}*/

func testAccCoralogixResourceRuleGroupBlock(r *ruleGroupParams, regEx, keepBlockedLogs string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
    rules{
	 block {
     	name               	= "%s"
	    description        	= "%s"
        source_field 		= "text"
        regular_expression	= "%s"
        keep_blocked_logs  	= "%s"
    	}
	}
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description, regEx, keepBlockedLogs)
}

func testAccCoralogixResourceRuleGroupJsonExtract(r *ruleGroupParams, jsonKey, destinationField string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
	rules{
     json_extract {
       name               	= "%s"
       description        	= "%s"
       json_key     		= "%s"
       destination_field  	= "%s"
     }
	}
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description, jsonKey, destinationField)
}

func testAccCoralogixResourceRuleGroupReplace(r *ruleGroupParams, regEx, replacementString string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
    rules{
      replace {
      name               	= "%s"
      description        	= "%s"
	  source_field       	= "text"
      destination_field  	= "text"
      regular_expression	= "%s"
      replacement_string     = "%s"
     }
	}
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description, regEx, replacementString)
}

func testAccCoralogixResourceRuleGroupAllow(r *ruleGroupParams, regEx, keepBlockedLogs string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
    rules{
       block {
      name               	= "%s"
      description        	= "%s"
      source_field 			= "text"
      regular_expression	= "%s"
      keep_blocked_logs     = "%s"
      blocking_all_matching_blocks = false
    }
   }
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description, regEx, keepBlockedLogs)
}

func testAccCoralogixResourceRuleGroupExtractTimestamp(r *ruleGroupParams, timeFormat, fieldFormatStandard string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
    rules{
       extract_timestamp {
      name               	= "%s"
      description        	= "%s"
      source_field 			= "text"
      time_format        	= "%s"
      field_format_standard = "%s"
    }
   }
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description, timeFormat, fieldFormatStandard)
}

func testAccCoralogixResourceRuleGroupRemoveFields(r *ruleGroupParams, excludedFields string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
	rules{
     remove_fields {
       name               = "%s"
       description        = "%s"
       excluded_fields    = %s
     }
   }
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description, excludedFields)
}

func testAccCoralogixResourceRuleGroupJsonStringify(r *ruleGroupParams, keepSourceField string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
    rules{
      json_stringify {
      name               = "%s"
      description        = "%s"
      source_field       = "text"
      destination_field  = "text"
      keep_source_field  = "%s"
    }
   }
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description, keepSourceField)
}

func testAccCoralogixResourceRuleGroupExtract(r *ruleGroupParams, regEx string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
    rules{
      extract {
       name               = "%s"
       description        = "%s"
       source_field       = "text"
       regular_expression = "%s"
     }
    }
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description, regEx)
}

func testAccCoralogixResourceRuleGroupParse(r *ruleGroupParams, regEx string) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
    rules{
      parse {
		name               = "%s"
		description        = "%s"
		source_field       = "text"
		destination_field  = "text"
		regular_expression  = "%s"
	  }
    }
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description, regEx)
}

func testAccCoralogixResourceRuleRulesCombination(r *ruleGroupParams) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
	rules{
    parse {
      name               = "rule1"
      description        = "description"
      source_field       = "text"
 	  destination_field  = "text"
      regular_expression  = "(?P<remote_addr>\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3})\\s*-\\s*(?P<user>[^ ]+)\\s*\\[(?P<timestemp>\\d{4}-\\d{2}\\-\\d{2}T\\d{2}\\:\\d{2}\\:\\d{2}\\.\\d{1,6}Z)\\]\\s*\\\\\\\"(?P<method>[A-z]+)\\s[\\/\\\\]+(?P<request>[^\\s]+)\\s*(?P<protocol>[A-z0-9\\/\\.]+)\\\\\\\"\\s*(?P<status>\\d+)\\s*(?P<body_bytes_sent>\\d+)?\\s*?\\\\\\\"(?P<http_referer>[^\"]+)\\\"\\s*\\\\\\\"(?P<http_user_agent>[^\"]+)\\\"\\s(?P<request_time>\\d{1,6})\\s*(?P<response_time>\\d{1,6})"
     }
    }
	rules{
     extract {
       name               = "rule2"
       description        = "description"
       source_field       = "text"
       regular_expression  = "(?P<remote_addr>\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3})\\s*-\\s*(?P<user>[^ ]+)\\s*\\[(?P<timestemp>\\d{4}-\\d{2}\\-\\d{2}T\\d{2}\\:\\d{2}\\:\\d{2}\\.\\d{1,6}Z)\\]\\s*\\\\\\\"(?P<method>[A-z]+)\\s[\\/\\\\]+(?P<request>[^\\s]+)\\s*(?P<protocol>[A-z0-9\\/\\.]+)\\\\\\\"\\s*(?P<status>\\d+)\\s*(?P<body_bytes_sent>\\d+)?\\s*?\\\\\\\"(?P<http_referer>[^\"]+)\\\"\\s*\\\\\\\"(?P<http_user_agent>[^\"]+)\\\"\\s(?P<request_time>\\d{1,6})\\s*(?P<response_time>\\d{1,6})"
     }
    }

	rules{
	 parse {
       name               = "rule3"
       description        = "description"
       source_field       = "text"
 	   destination_field  = "text"
       regular_expression  = "(?P<remote_addr>\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3})\\s*-\\s*(?P<user>[^ ]+)\\s*\\[(?P<timestemp>\\d{4}-\\d{2}\\-\\d{2}T\\d{2}\\:\\d{2}\\:\\d{2}\\.\\d{1,6}Z)\\]\\s*\\\\\\\"(?P<method>[A-z]+)\\s[\\/\\\\]+(?P<request>[^\\s]+)\\s*(?P<protocol>[A-z0-9\\/\\.]+)\\\\\\\"\\s*(?P<status>\\d+)\\s*(?P<body_bytes_sent>\\d+)?\\s*?\\\\\\\"(?P<http_referer>[^\"]+)\\\"\\s*\\\\\\\"(?P<http_user_agent>[^\"]+)\\\"\\s(?P<request_time>\\d{1,6})\\s*(?P<response_time>\\d{1,6})"
     }
	} 
  }

  rule_subgroups {
   rules{
	extract_timestamp {
      name               	= "rule1"
      description        	= "description"
      source_field 			= "text"
      time_format        	= "2006-01-02T15:04:05.999999999Z07:00"
      field_format_standard = "Golang"
    }
   }
  }
 }
`, r.name, r.description, r.creator)
}

func testAccCoralogixResourceRuleRulesCombinationDifferentOrders(r *ruleGroupParams) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
	rules{
     extract {
       name               = "rule2"
       description        = "description"
       source_field       = "text"
       regular_expression  = "(?P<remote_addr>\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3})\\s*-\\s*(?P<user>[^ ]+)\\s*\\[(?P<timestemp>\\d{4}-\\d{2}\\-\\d{2}T\\d{2}\\:\\d{2}\\:\\d{2}\\.\\d{1,6}Z)\\]\\s*\\\\\\\"(?P<method>[A-z]+)\\s[\\/\\\\]+(?P<request>[^\\s]+)\\s*(?P<protocol>[A-z0-9\\/\\.]+)\\\\\\\"\\s*(?P<status>\\d+)\\s*(?P<body_bytes_sent>\\d+)?\\s*?\\\\\\\"(?P<http_referer>[^\"]+)\\\"\\s*\\\\\\\"(?P<http_user_agent>[^\"]+)\\\"\\s(?P<request_time>\\d{1,6})\\s*(?P<response_time>\\d{1,6})"
     }
    }

	rules{
	 parse {
       name               = "rule3"
       description        = "description"
       source_field       = "text"
 	   destination_field  = "text"
       regular_expression  = "(?P<remote_addr>\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3})\\s*-\\s*(?P<user>[^ ]+)\\s*\\[(?P<timestemp>\\d{4}-\\d{2}\\-\\d{2}T\\d{2}\\:\\d{2}\\:\\d{2}\\.\\d{1,6}Z)\\]\\s*\\\\\\\"(?P<method>[A-z]+)\\s[\\/\\\\]+(?P<request>[^\\s]+)\\s*(?P<protocol>[A-z0-9\\/\\.]+)\\\\\\\"\\s*(?P<status>\\d+)\\s*(?P<body_bytes_sent>\\d+)?\\s*?\\\\\\\"(?P<http_referer>[^\"]+)\\\"\\s*\\\\\\\"(?P<http_user_agent>[^\"]+)\\\"\\s(?P<request_time>\\d{1,6})\\s*(?P<response_time>\\d{1,6})"
      }
    }

	rules{
     parse {
      name               = "rule1"
      description        = "description"
      source_field       = "text"
 	  destination_field  = "text"
      regular_expression  = "(?P<remote_addr>\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3})\\s*-\\s*(?P<user>[^ ]+)\\s*\\[(?P<timestemp>\\d{4}-\\d{2}\\-\\d{2}T\\d{2}\\:\\d{2}\\:\\d{2}\\.\\d{1,6}Z)\\]\\s*\\\\\\\"(?P<method>[A-z]+)\\s[\\/\\\\]+(?P<request>[^\\s]+)\\s*(?P<protocol>[A-z0-9\\/\\.]+)\\\\\\\"\\s*(?P<status>\\d+)\\s*(?P<body_bytes_sent>\\d+)?\\s*?\\\\\\\"(?P<http_referer>[^\"]+)\\\"\\s*\\\\\\\"(?P<http_user_agent>[^\"]+)\\\"\\s(?P<request_time>\\d{1,6})\\s*(?P<response_time>\\d{1,6})"
     }
    }
  }

  rule_subgroups {
   rules{
	extract_timestamp {
      name               	= "rule1"
      description        	= "description"
      source_field 			= "text"
      time_format        	= "2006-01-02T15:04:05.999999999Z07:00"
      field_format_standard = "Golang"
    }
   }
  }
 }
`, r.name, r.description, r.creator)
}

type ruleParams struct {
	name, description string
}

type ruleGroupParams struct {
	ruleParams
	name, description, creator string
}
