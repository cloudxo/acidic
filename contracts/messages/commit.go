package messages

import "time"

type CommitTransactionCommand struct {
	TransactionID string
}
type TransactionCommittingEvent struct {
	TransactionID string
	Contents      map[string]*CommitItem
}
type TransactionCommittedEvent struct {
	Timestamp     time.Time
	Commit        uint64
	TransactionID string
}
type TransactionCommitFailedEvent struct {
	Timestamp     time.Time
	TransactionID string
}

type CommitItem struct {
	Key      string
	Revision string
	ETag     string
}

type CommitWrittenEvent struct {
	Timestamp      time.Time
	Commit         uint64
	TransactionIDs []string
}

func (this CommitTransactionCommand) CorrelationID() string {
	return "" // TODO
}

func (this TransactionCommittedEvent) CorrelationID() string {
	return "" // TODO
}
