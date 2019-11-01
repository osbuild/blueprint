package blueprint

import "github.com/osbuild/osbuild-composer/internal/pipeline"

type tarOutput struct{}

func (t *tarOutput) translate(b *Blueprint) (*pipeline.Pipeline, error) {
	packages := [...]string{
		"policycoreutils",
		"selinux-policy-targeted",
		"kernel",
		"firewalld",
		"chrony",
		"langpacks-en",
	}
	excludedPackages := [...]string{
		"dracut-config-rescue",
	}
	p := getCustomF30PackageSet(packages[:], excludedPackages[:], b)
	addF30LocaleStage(p)
	addF30GRUB2Stage(p, b.getKernelCustomization())
	addF30FixBlsStage(p)
	addF30SELinuxStage(p)
	addF30TarAssembler(p, t.getName(), "xz")

	if b.Customizations != nil {
		err := b.Customizations.customizeAll(p)
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}

func (t *tarOutput) getName() string {
	return "root.tar.xz"
}

func (t *tarOutput) getMime() string {
	return "application/x-tar"
}
