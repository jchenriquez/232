package main

import "fmt"

type MyQueue struct {
   tail []int
   head []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{make([]int, 0), make([]int, 0)}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.tail = append(this.tail, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
   if len(this.head) == 0 {
     for i := len(this.tail)-1; i >= 0; i-- {
       val := this.tail[i]
       this.head = append(this.head, val)
     }

    this.tail = make([]int, 0)
   } 

   retVal := this.head[len(this.head)-1]
   this.head = this.head[:len(this.head)-1]

   return retVal
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
  if len(this.head) == 0 {
     for i := len(this.tail)-1; i >= 0; i-- {
       val := this.tail[i]
       this.head = append(this.head, val)
     }
     this.tail = make([]int, 0)
   } 

  return this.head[len(this.head)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.head) == 0 && len(this.tail) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
  fmt.Println("Hello World")
}