package constants

import "time"

// General kubernetes-related constants
const (
	PostgresContainerName    = "postgres"
	PostgresContainerIdx     = 0
	K8sAPIPath               = "/apis"
	ResourceDeletionInterval = 1 * time.Second
	ResourceDeletionTimeout  = 30 * time.Second

	QueueResyncPeriodPod  = 5 * time.Minute
	QueueResyncPeriodTPR  = 5 * time.Minute
	QueueResyncPeriodNode = 5 * time.Minute
)
