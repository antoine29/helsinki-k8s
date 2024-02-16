package controllers

import (
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

func fillMemo() {
  fmt.Println("starting memory usage func.")
  dummy_array := make([]int, 32000000)
  for i := 0; i < len(dummy_array); i++ {
   dummy_array[i] = math.MaxUint32
  }
  
  time.Sleep(5*time.Minute)
  fmt.Printf("exiting memory usage func %d \n", len(dummy_array))
}

func MemoUsage(c *gin.Context) {
  go fillMemo()
}

