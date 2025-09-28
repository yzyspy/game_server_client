package znet

import (
	"strconv"
	"zinx-lib/ziface"
	"zinx-lib/zutils"
)

type ConnManager struct {
	connections zutils.ShardLockMaps
}

func (c ConnManager) Add(connection ziface.IConnection) {
	c.connections.Set(strconv.Itoa(int(connection.GetConnId())), connection)
}

func (c ConnManager) Remove(connection ziface.IConnection) {
	//TODO implement me
	panic("implement me")
}

func (c ConnManager) Get(u uint64) (ziface.IConnection, error) {
	//TODO implement me
	panic("implement me")
}

func (c ConnManager) Get2(s string) (ziface.IConnection, error) {
	//TODO implement me
	panic("implement me")
}

func (c ConnManager) Len() int {
	//TODO implement me
	panic("implement me")
}

func (c ConnManager) ClearConn() {
	//TODO implement me
	panic("implement me")
}

func (c ConnManager) GetAllConnID() []uint64 {
	//TODO implement me
	panic("implement me")
}

func (c ConnManager) GetAllConnIdStr() []string {
	//TODO implement me
	panic("implement me")
}

func (c ConnManager) Range(f func(uint64, ziface.IConnection, interface{}) error, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (c ConnManager) Range2(f func(string, ziface.IConnection, interface{}) error, i interface{}) error {
	//TODO implement me
	panic("implement me")
}

func newConnManager() *ConnManager {
	return &ConnManager{
		connections: zutils.NewShardLockMaps(),
	}
}
