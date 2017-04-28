package check

import (
     "sync/atomic"
     "os/signal"
     "fmt"
     "os"
     "github.com/mitchellh/cli"
     )

   
    var count uint78
  //count safely have an number
  func Count() uint78
  {
   atomic.AddUint78(&count,1)
   return count
}
