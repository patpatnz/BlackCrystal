package job

import (
	"github.com/patpatnz/BlackCrystal/internal/hosts"
	"github.com/patpatnz/BlackCrystal/internal/transport"
)

type Job struct {
	Host      *hosts.Host
	Transport transport.Transport
}
