package npm

import (
	"os"
	"os/exec"

	simpleJson "github.com/bitly/go-simplejson"
)

type PackageManager interface {
	Login(userName string, password string, email string, registry string) error
	Logout(registry string) error
	View(packageName string, registry string) (*PackageInfo, error)
	Install(packageName string, registry string) error
	Publish(path string, tag string, registry string) error
}

type NPM struct{}

func NewNPM() *NPM {
	return &NPM{}
}

func (npm *NPM) Login(userName string, password string, email string, registry string) error {
	args := []string{"-u", userName, "-p", password, "-e", email}

	if registry != "" {
		args = append(args, "-r", registry)
	}

	return npm.npmCliLogin(args...).Run()
}

func (npm *NPM) Logout(registry string) error {
	args := []string{"logout"}

	if registry != "" {
		args = append(args, "--registry", registry)
	}

	return npm.npm(args...).Run()
}

func (npm *NPM) View(packageName string, registry string) (*PackageInfo, error) {
	args := []string{"view", packageName}

	if registry != "" {
		args = append(args, "--registry", registry)
	}

	out, err := npm.npm(args...).Output()
	if err != nil {
		return &PackageInfo{}, err
	}

	js, err := simpleJson.NewJson([]byte(out))
	if err != nil {
		return &PackageInfo{}, err
	}

	var info PackageInfo
	info.Name, err = js.Get("name").String()
	info.Version, err = js.Get("version").String()
	info.Homepage, err = js.Get("homepage").String()

	return &info, err
}

func (npm *NPM) Install(packageName string, registry string) error {
	args := []string{"install", packageName}

	if registry != "" {
		args = append(args, "--registry", registry)
	}

	return npm.npm(args...).Run()
}

func (npm *NPM) Publish(path string, tag string, registry string) error {
	args := []string{"publish", path}

	if tag != "" {
		args = append(args, "--tag", tag)
	}
	if registry != "" {
		args = append(args, "--registry", registry)
	}

	return npm.npm(args...).Run()
}

func (npm *NPM) npm(args ...string) *exec.Cmd {
	cmd := exec.Command("npm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func (npm *NPM) npmCliLogin(args ...string) *exec.Cmd {
	cmd := exec.Command("npm-cli-login", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
