terraform {
    required_providers {
        coralogix = {
            version = "~> 2.0"
            source  = "coralogix/coralogix"
        }
    }
}

provider "coralogix" {
    api_key = "cxtp_obZw8PugYD7R64OIp9d3u98Juvj7qb"
    env = "EU2"
}

resource "coralogix_alert" "test" {
    name        = "alert threshold with destinations should trigger for sure fsdjkfskjf"
    description = ""
    enabled     = true
    phantom_mode = false
    notification_group = {
    }
    incidents_settings = {
        notify_on = "Triggered Only"
        retriggering_period = {
            minutes = 10
        }
    }
    type_definition = {
        metric_threshold = {
            metric_filter = {
                promql = "http_request"
            }
            rules = [
                {
                    condition = {
                        threshold      = 4000
                        for_over_pct   = 30
                        of_the_last    = "10_MINUTES"
                        condition_type = "MORE_THAN"
                    }
                    override = {
                        priority = "P3"
                    }
                },
            ]
            missing_values = {
                min_non_null_values_pct = 0
            }
        }
    }
}

