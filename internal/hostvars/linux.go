package hostvars

import (
	"log"

	"github.com/patpatnz/BlackCrystal/internal/job"
)

func init() {

}

func collectLinuxVars(job *job.Job) error {

	v, err := job.Transport.ReadFile("/proc/cpuinfo")
	if err != nil {
		return nil
	}
	log.Printf("v: %s", v)
	return nil
}

func Run(job *job.Job) {
	collectLinuxVars(job)
}
