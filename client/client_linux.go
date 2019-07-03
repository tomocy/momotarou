package client

import (
	"fmt"

	kibidangoPkg "github.com/tomocy/kibidango"
	factoryPkg "github.com/tomocy/kibidango/factory"
	"github.com/tomocy/momotaro/spec"
)

type linux struct{}

func (l *linux) list() ([]kibidango, error) {
	factory := l.factory()
	kibis, err := factory.List()
	if err != nil {
		return nil, err
	}

	return l.adaptAll(kibis), nil
}

func (l *linux) adaptAll(kibis []*kibidangoPkg.Linux) []kibidango {
	adapteds := make([]kibidango, len(kibis))
	for i, kibi := range kibis {
		adapteds[i] = kibi
	}

	return adapteds
}

func (l *linux) create(spec *spec.Spec) (kibidango, error) {
	factory := l.factory()
	kibi, err := factory.Manufacture(spec.ID)
	if err != nil {
		return nil, err
	}
	if err := l.letKibidangoMeetSpec(kibi, spec); err != nil {
		return nil, err
	}

	if err := factory.Save(kibi); err != nil {
		return nil, err
	}

	return kibi, nil
}

func (l *linux) letKibidangoMeetSpec(kibi *kibidangoPkg.Linux, spec *spec.Spec) error {
	return kibi.UpdateProcess(spec.Process)
}

func (l *linux) load(id string) (kibidango, error) {
	factory := l.factory()
	return factory.Load(id)
}

func (l *linux) delete(id string) error {
	factory := l.factory()
	return factory.Delete(id)
}

func (l *linux) factory() *factoryPkg.Linux {
	return factoryPkg.ForLinux()
}

func (l *linux) printAll(kibis []kibidango) {
	printHeader()
	for _, kibi := range kibis {
		l.print(kibi)
	}
}

func (l *linux) print(kibi kibidango) {
	linux := kibi.(*kibidangoPkg.Linux)
	printable := printableLinux(*linux)
	fmt.Println(printable)
}

type printableLinux kibidangoPkg.Linux

func (p printableLinux) String() string {
	return fmt.Sprintf("%s", p.ID())
}
