package main

type Directory struct {
	name        string
	files       []File
	directories map[string]*Directory
	previous    *Directory
	size        int
}

type Command struct {
	name string
	arg  string
}

type File struct {
	name string
	size int
}
