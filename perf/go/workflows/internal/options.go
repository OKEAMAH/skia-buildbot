package internal

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

var (
	// Default option for the regular activity.
	//
	// Activity usually communicates with the external services and is expected to complete
	// within a minute. RetryPolicy helps to recover from unexpected network errors or service
	// interruptions.
	// For activities that expect long running time and complex dependent services, a separate
	// option should be curated for individual activities.
	regularActivityOptions = workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 10,
		},
	}

	// Default option for the child workflow.
	//
	// This generally means time tolerance from the most top level workflow, in this case, it is
	// the bisection workflow. The actual timeout heavily depends on the swarming resources.
	// We don't want to leave this running for very long but also know there are cases where
	// the resources will not be immediately available.
	// This setting indicates that each child job should finish within 12 hours.
	childWorkflowOptions = workflow.ChildWorkflowOptions{
		// 4 hours of compile time + 8 hours of test run time
		WorkflowExecutionTimeout: 12 * time.Hour,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 4,
		},
	}
)
