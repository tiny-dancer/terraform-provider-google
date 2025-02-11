// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeNetwork_networkBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeNetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetwork_networkBasicExample(context),
			},
			{
				ResourceName:      "google_compute_network.vpc_network",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeNetwork_networkBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "vpc_network" {
  name = "tf-test-vpc-network%{random_suffix}"
}
`, context)
}

func TestAccComputeNetwork_networkCustomMtuExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeNetworkDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeNetwork_networkCustomMtuExample(context),
			},
			{
				ResourceName:      "google_compute_network.vpc_network",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeNetwork_networkCustomMtuExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_network" "vpc_network" {
  name = "tf-test-vpc-network%{random_suffix}"
  mtu  = 1500
}
`, context)
}

func testAccCheckComputeNetworkDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_network" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/global/networks/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ComputeNetwork still exists at %s", url)
			}
		}

		return nil
	}
}
