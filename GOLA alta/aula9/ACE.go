package main

import (
	"fmt"
)

// Estrutura de um nó da fila
type Node struct {
	data interface{}
	next *Node
}

// Estrutura da fila
type Queue struct {
	front *Node
	rear  *Node
	size  int
}

// Função para criar uma nova fila vazia
func NewQueue() *Queue {
	return &Queue{}
}

// Função para verificar se a fila está vazia
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Função para obter o tamanho da fila
func (q *Queue) Size() int {
	return q.size
}

// Função para enfileirar um novo elemento
func (q *Queue) Enqueue(data interface{}) {
	newNode := &Node{data, nil}
	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

// Função para desenfileirar um elemento
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.front.data
	q.front = q.front.next
	q.size--
	if q.IsEmpty() {
		q.rear = nil
	}
	return data
}

// Função para percorrer e exibir os elementos da fila
func (q *Queue) Traverse() {
	if q.IsEmpty() {
		fmt.Println("Fila vazia")
		return
	}
	current := q.front
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	queue := NewQueue()

	// Teste de inserção
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Teste de remoção
	fmt.Println("Elemento removido:", queue.Dequeue())

	// Teste de percorrer a fila
	fmt.Println("Elementos na fila:")
	queue.Traverse()
}
