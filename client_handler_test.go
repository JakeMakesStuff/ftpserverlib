package ftpserver

import (
	"sync"
	"testing"

	"gopkg.in/dutchcoders/goftp.v1"
)

func TestConcurrency(t *testing.T) {
	s := NewTestServer(false)
	defer mustStopServer(s)

	nbClients := 100

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(nbClients)

	for i := 0; i < nbClients; i++ {
		go func() {
			var err error
			var ftp *goftp.FTP

			if ftp, err = goftp.Connect(s.Addr()); err != nil {
				panic(err)
			}

			defer func() { panicOnError(ftp.Close()) }()

			if err = ftp.Login("test", "test"); err != nil {
				panic(err)
			}

			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
