// Package boot contains low level code.
package boot

import (
	"github.com/krzysztofzaucha/plugin-playground/internal"
	"github.com/sarulabs/di"
)

// Container service container.
type Container struct {
	app di.Container
}

// New container factory.
func New() *Container {
	builder, err := di.NewBuilder()
	if err != nil {
		panic(err)
	}

	err = builder.Add([]di.Def{
		{
			Name: "manager",
			Build: func(ctn di.Container) (interface{}, error) {
				return internal.NewManager(), nil
			},
		},
	}...)
	if err != nil {
		panic(err)
	}

	return &Container{
		app: builder.Build(),
	}
}

// Manager is a loader and executor of a plugin.
func (c *Container) Manager() *internal.Manager {
	p, ok := c.app.Get("manager").(*internal.Manager)
	if !ok {
		panic("manager service type is incorrect")
	}

	return p
}
