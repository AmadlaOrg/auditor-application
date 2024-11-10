package application

import "runtime"

type IApplication interface {
	FindApplication(name string, isContainer bool) (Application, error)
}

type SApplication struct{}

func (s *SApplication) FindApplication(name string, isContainer bool) (Application, error) {
	if runtime.GOOS == "linux" {
		s.linuxLocal(name)
	} else if runtime.GOOS == "windows" {
		s.windowsLocal(name)
	}
	return Application{}, nil
}

func (s *SApplication) linuxLocal(name string) {

}

func (s *SApplication) windowsLocal(name string) {

}

func (s *SApplication) container(name string) {

}
