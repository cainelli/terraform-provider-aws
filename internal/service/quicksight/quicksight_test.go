package quicksight_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccQuickSight_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"AccountSubscription": {
			"basic":      testAccAccountSubscription_basic,
			"disappears": testAccAccountSubscription_disappears,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
