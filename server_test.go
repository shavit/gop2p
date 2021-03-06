package gop2p

import (
  "testing"
  "time"
)

func TestServerStart(t *testing.T){
  var err error
  var srv *server = NewServer()

  if srv == nil {
    t.Error("Error creating a server")
  }

  // Close after n time
  go func() {
    <-time.After(20 * time.Millisecond)
    srv.Close(nil)
  }()

  err = srv.Start()
  if err != nil {
    t.Error(err)
  }
}
