package out

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/npm"
)

type Command struct {
	packageManager npm.PackageManager
}

func NewCommand(packageManager npm.PackageManager) *Command {
	return &Command{
		packageManager: packageManager,
	}
}

func (command *Command) Run(request Request) (Response, error) {
	err := command.packageManager.Login(
		request.Params.UserName,
		request.Params.Password,
		request.Params.Email,
		request.Source.Registry,
	)
	if err != nil {
		return Response{}, err
	}

	tag, err := tagFrom(request.Params)
	if err != nil {
		return Response{}, err
	}

	err = command.packageManager.Publish(
		request.Params.Path,
		tag,
		request.Source.Registry,
	)
	if err != nil {
		return Response{}, err
	}

	out, err := command.packageManager.View(
		request.Source.PackageName,
		request.Source.Registry,
	)
	if err != nil {
		return Response{}, err
	}

	err = command.packageManager.Logout(
		request.Source.Registry,
	)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Version: resource.Version{
			Version: out.Version,
		},
		Metadata: []resource.MetadataPair{
			{
				Name:  "name",
				Value: out.Name,
			},
			{
				Name:  "homepage",
				Value: out.Homepage,
			},
		},
	}, nil
}

func tagFrom(params Params) (string, error) {
	if len(params.TagFile) == 0 {
		return params.Tag, nil
	}

	t, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", os.Args[1], params.TagFile))
	if err != nil {
		return "", err
	}

	return string(t), nil
}
