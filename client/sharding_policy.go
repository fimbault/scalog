package client

import (
	"math/rand"
	"time"

	disc "github.com/scalog/scalog/discovery"
)

type DefaultShardingPolicy struct {
	shardID    int32
	replicaID  int32
	numReplica int32
	seed       rand.Source
}

func NewDefaultShardingPolicy(numReplica int32) *DefaultShardingPolicy {
	s := &DefaultShardingPolicy{
		shardID:    -1,
		replicaID:  -1,
		numReplica: numReplica,
		seed:       rand.NewSource(time.Now().UnixNano()),
	}
	return s
}

func (p DefaultShardingPolicy) Shard(view *disc.View, record string) (int32, int32) {
	if view == nil {
		return -1, -1
	}
	if s, ok := view.Shards[p.shardID]; ok && s {
		return p.shardID, p.replicaID
	}
	rs := rand.New(p.seed).Intn(len(view.LiveShards))
	rr := int32(rand.New(p.seed).Intn(int(p.numReplica)))
	return view.LiveShards[rs], rr
}