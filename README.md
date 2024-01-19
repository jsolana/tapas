# tapas

<p align="center"><img src="logo.png" alt="atelier logo"/></p>
<p align="center">
    <a href="https://github.com/jsolana/tapas/releases"><img src="https://img.shields.io/github/release/jsolana/tapas.svg" alt="release"></a>
    <a href="https://github.com/jsolana/tapas/actions/workflows/build.yml"><img src="https://github.com/jsolana/tapas/actions/workflows/build.yml/badge.svg" alt="Build status"></a>
    <a href="https://github.com/jsolana/tapas/issues"><img src="https://img.shields.io/github/issues/jsolana/tapas.svg?style=flat-square" alt="issues"></a>
    <a href="https://github.com/jsolana/tapas/blob/master/LICENSE"><img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="license"></a>
</p>

`tapas` is a Go library providing textbook implementations of generic data structures & algorithms. Its design focuses on simplicity, performance, and concurrent usage.

## Usage

Using this package requires a working Go environment. [See the install instructions for Go](http://golang.org/doc/install.html).

```bash
go get github.com/jsolana/tapas
```

## Data structures

### List

A *doubly linked list* is a type of linked list in which each node contains a data element and two pointers, one pointing to the next node in the sequence (forward direction) and another pointing to the previous node (backward direction). Unlike a singly linked list, where nodes only have a reference to the next node, a doubly linked list allows traversal in both forward and backward directions.

Each node in a doubly linked list consists of three fields:

- *Value*: The actual data or value stored in the node.
- *Next Pointer*: A pointer that points to the next node in the sequence.
- *Previous Pointer*: A pointer that points to the previous node in the sequence.

The doubly linked list structure provides advantages such as easy backward traversal, efficient insertion and deletion of nodes at both the beginning and end of the list, and the ability to delete a node without needing a reference to its predecessor. However, it requires more memory compared to a singly linked list due to the additional storage needed for the backward pointers.

## Credits

`tapas` logo was generated using [replicate](https://replicate.com/)