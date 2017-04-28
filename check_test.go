package check_test

import (
   "testing"
   "github.com/thasviya/go-unittest/check"  //here it checks whether do we have check.go file in our path are not
   )

 func TestCount(t *testing.T)
 {
  if check.Count() != 1
  {           //checks first for only the common test file
   t.Error("expected 1")
  }

  if check.Count() != 2
  {           //checks for the second time for test file
   t.Error("expected 2")
  }

  if check.Count() !=3
  {           //checks for test file
   t.Error("expected 3")
  }

  else if Count() != 1
  {       
   t.Error("expected 1")
  }
  else if count != 1
  {     //this checks for an internals of the test files
   t.Error("expected 1 for count too")
  }

 }

//here there is been merged of two different *_test file and *_internals_test file into one code. 
