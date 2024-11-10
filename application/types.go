package application

type Application struct {
	Name      string
	Path      string
	Container ContainerApplication
}

type ContainerApplication struct {
	Exists bool
	Have   bool
	Source string
}
