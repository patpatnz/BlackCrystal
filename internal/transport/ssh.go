package transport

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/patpatnz/BlackCrystal/internal/hosts"
	"github.com/pkg/sftp"

	gossh "golang.org/x/crypto/ssh"
)

type ssh struct {
	client *gossh.Client
	sftp   *sftp.Client
}

func init() {
	Register(&ssh{})
}

func (s ssh) GetName() string {
	return "SSH"
}

func (s ssh) GetInstance(host *hosts.Host) (Transport, error) {
	t := &ssh{}
	t.Connect(host)
	return t, nil
}

func (s *ssh) Connect(host *hosts.Host) error {

	key, err := ioutil.ReadFile("/Users/pjs/.ssh/id_dsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := gossh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &gossh.ClientConfig{
		User: "pjs",
		Auth: []gossh.AuthMethod{
			gossh.PublicKeys(signer),
		},
	}

	port := "22"
	if p := host.Vars["ansible_ssh_port"]; p != nil {
		port = p.(string)
	}

	if ip := host.Vars["ansible_ssh_host"]; ip != nil {
		s.client, err = gossh.Dial("tcp", ip.(string)+":"+port, config)
	} else {
		s.client, err = gossh.Dial("tcp", host.Name+":"+port, config)
	}
	if err != nil {
		log.Printf("Error connecting to host %s: %s", host.Name, err)
		return err
	}

	s.sftp, err = sftp.NewClient(s.client)
	if err != nil {
		log.Printf("Couldn't open SFTP connection")
	}

	return nil
}

func (s ssh) Run(cmd string) ([]string, error) {
	sess, err := s.client.NewSession()
	if err != nil {
		log.Printf("err: %s", err)
		return nil, err
	}

	var b bytes.Buffer
	sess.Stdout = &b

	sess.Run(cmd)

	log.Printf("result: %s", b.Bytes())

	return nil, nil
}

func (s ssh) ReadFile(file string) ([]byte, error) {
	var b bytes.Buffer

	f, err := s.sftp.Open(file)
	if err != nil {
		return nil, err
	}

	b.ReadFrom(f)
	f.Close()

	return b.Bytes(), nil
}
