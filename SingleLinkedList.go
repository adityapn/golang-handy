package main
import ("fmt")

type Node struct{
 link *Node
 value int32
}

type LinkedList struct {
 root *Node
 size int32
}

// Create related functions
func(ll *LinkedList) append(value int32){
  ll.increment()
  if(ll.root == nil){
    ll.root = &Node{link:nil,value:value}
    return
  }
  temp := ll.root
  for ;temp.link != nil;temp = temp.link{}

  newNode := &Node{link:nil,value:value}
  temp.link = newNode
}

func (ll *LinkedList) addFirst(value int32) {
  if ll.root == nil{
    ll.append(value)
    return
  }
  
  newNode := &Node{value:value,link:ll.root}
  ll.root = newNode
  ll.increment()
}

func(ll *LinkedList) addAfter(value int32,position int32) {
  if position < 0 || position > ll.size{
    fmt.Println("Overflow")
    return
  }    
}

// Update the linked list functions
func(ll *LinkedList) reverse(){
  current := ll.root
  var next *Node
  var prev *Node
  for ;current != nil;{
    next = current.link
    current.link = prev
    prev = current
    current = next
  }
  ll.root = prev
}

// Delete related functions

func(ll *LinkedList) delete(value int32) {
  if ll.root == nil{
    fmt.Println("Overflow")
    return
  }

  iterator := ll.root
  // incase of deleting root node
  if ll.root.value == value{
    ll.root = ll.root.link
    return
  }

  for ;iterator.link != nil;iterator = iterator.link{
    if iterator.link.value == value{
      iterator.link = iterator.link.link
      break
    }    
  }
}
// Retrieve related functions
func(ll *LinkedList) show(){
  temp := ll.root
  for ;temp != nil;temp = temp.link{
    fmt.Println(temp.value);
  }
}

func(ll *LinkedList) increment() {
  ll.size++
}

func(ll *LinkedList) decrement() {
  ll.size--
}


func main(){
  ll := LinkedList{}
  ll.append(2)
  ll.append(3)
  ll.append(4)

  ll.addFirst(1)
  
  ll.append(5)
  ll.append(6)
  ll.append(7)
  
  ll.delete(2)
  //ll.reverse()
  ll.show()
}
