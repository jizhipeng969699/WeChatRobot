//go:generate stringer -type=Sender

package gogenerate

type Sender int

const (
	Kafka Sender = iota
	Etcd
	File
	Bytes
)
