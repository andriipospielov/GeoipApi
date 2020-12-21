package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net"
	"testing"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) CloseConnection() {

}

func (r *RepoMock) FindByIp(ip net.IP) (interface{}, error) {
	return 1, nil
}

func TestIpReposRegistry_Instances(t *testing.T) {
	var repoMock GeoipRepository = new(RepoMock)
	sut := IpReposRegistry{"1": repoMock}
	assert.Same(t, repoMock, sut.Instances()["1"])
}
