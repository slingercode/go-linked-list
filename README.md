# Linked List

## Build project
`go build -o bin/linked-list`

## Advantages over arrays

1. Dynamic size
2. Ease of insertion/deletion

**Note:**

Let's represent an array of ordered numbers:

`const numbers = [10, 20, 30, 40, 50];`

In this example, in the case of the of a new insertion of a value
like **25**, to mantain the order we need to move all of the elements
after 20 (excluding 20).
Deletion is also expensive, the exemple here is if we remove the value of **10**
we need to move all of the elements after 10.

## Drawbacks

1. Non existing random access, the reason behind this is becouse we need to access
   the elements starting from the first node
2. Extra memory, we need to store the value of a pointer
3. Not cache friendly

## Representation

A linked list is representated by a pointer to the first node of the linked list.
The first node is called _Head_. If the linked list is empty then the value of
the head points to `null`.

Each node in a list consists in at least two things:

- Data, the information that we are going to store
- Pointer, the reference to the next node

## References

- [geeksforgeeks - linked lists](https://www.geeksforgeeks.org/data-structures/linked-list)
- [TheAlgorithms](https://github.com/TheAlgorithms/Javascript)
