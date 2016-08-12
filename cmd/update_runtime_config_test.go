package cmd_test

import (
	"errors"

	semver "github.com/cppforlife/go-semi-semantic/version"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-init/cmd"
	fakecmd "github.com/cloudfoundry/bosh-init/cmd/fakes"
	fakedir "github.com/cloudfoundry/bosh-init/director/fakes"
	boshtpl "github.com/cloudfoundry/bosh-init/director/template"
	fakeui "github.com/cloudfoundry/bosh-init/ui/fakes"
)

var _ = Describe("UpdateRuntimeConfigCmd", func() {
	var (
		ui               *fakeui.FakeUI
		director         *fakedir.FakeDirector
		uploadReleaseCmd *fakecmd.FakeReleaseUploadingCmd
		command          UpdateRuntimeConfigCmd
	)

	BeforeEach(func() {
		ui = &fakeui.FakeUI{}
		director = &fakedir.FakeDirector{}
		uploadReleaseCmd = &fakecmd.FakeReleaseUploadingCmd{}
		command = NewUpdateRuntimeConfigCmd(ui, director, uploadReleaseCmd)
	})

	Describe("Run", func() {
		var (
			opts UpdateRuntimeConfigOpts
		)

		BeforeEach(func() {
			opts = UpdateRuntimeConfigOpts{
				Args: UpdateRuntimeConfigArgs{
					RuntimeConfig: FileBytesArg{Bytes: []byte("runtime: config")},
				},
			}
		})

		act := func() error { return command.Run(opts) }

		It("updates runtime config", func() {
			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(1))

			bytes := director.UpdateRuntimeConfigArgsForCall(0)
			Expect(bytes).To(Equal([]byte("runtime: config\n")))
		})

		It("updates runtime config with evaluated vars", func() {
			opts.Args.RuntimeConfig = FileBytesArg{
				Bytes: []byte("name1: ((name1))\nname2: ((name2))"),
			}

			opts.VarKVs = []boshtpl.VarKV{
				{Name: "name1", Value: "val1-from-kv"},
			}

			opts.VarsFiles = []boshtpl.VarsFileArg{
				{Vars: boshtpl.Variables(map[string]interface{}{"name1": "val1-from-file"})},
				{Vars: boshtpl.Variables(map[string]interface{}{"name2": "val2-from-file"})},
			}

			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(1))

			bytes := director.UpdateRuntimeConfigArgsForCall(0)
			Expect(bytes).To(Equal([]byte("name1: val1-from-kv\nname2: val2-from-file\n")))
		})

		It("uploads remote releases", func() {
			opts.Args.RuntimeConfig = FileBytesArg{
				Bytes: []byte(`
releases:
- name: capi
  sha1: capi-sha1
  url: https://capi-url
  version: 1+capi
- name: consul
  sha1: consul-sha1
  url: https://consul-url
  version: 1+consul
`),
			}

			err := act()
			Expect(err).ToNot(HaveOccurred())

			Expect(uploadReleaseCmd.RunCallCount()).To(Equal(2))

			Expect(uploadReleaseCmd.RunArgsForCall(0)).To(Equal(UploadReleaseOpts{
				Name:    "capi",
				Args:    UploadReleaseArgs{URL: URLArg("https://capi-url")},
				SHA1:    "capi-sha1",
				Version: VersionArg(semver.MustNewVersionFromString("1+capi")),
			}))

			Expect(uploadReleaseCmd.RunArgsForCall(1)).To(Equal(UploadReleaseOpts{
				Name:    "consul",
				Args:    UploadReleaseArgs{URL: URLArg("https://consul-url")},
				SHA1:    "consul-sha1",
				Version: VersionArg(semver.MustNewVersionFromString("1+consul")),
			}))
		})

		It("returns error and does not deploy if uploading release fails", func() {
			opts.Args.RuntimeConfig = FileBytesArg{
				Bytes: []byte(`
releases:
- name: capi
  sha1: capi-sha1
  url: https://capi-url
  version: 1+capi
`),
			}
			uploadReleaseCmd.RunReturns(errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(0))
		})

		It("returns an error if release version cannot be parsed", func() {
			opts.Args.RuntimeConfig = FileBytesArg{
				Bytes: []byte(`
name: dep
releases:
- name: capi
  sha1: capi-sha1
  url: https://capi-url
  version: 1+capi+capi
`),
			}

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Expected version '1+capi+capi' to match version format"))

			Expect(uploadReleaseCmd.RunCallCount()).To(Equal(0))
			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(0))
		})

		It("does not update if confirmation is rejected", func() {
			ui.AskedConfirmationErr = errors.New("stop")

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("stop"))

			Expect(director.UpdateRuntimeConfigCallCount()).To(Equal(0))
		})

		It("returns error if updating failed", func() {
			director.UpdateRuntimeConfigReturns(errors.New("fake-err"))

			err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})
	})
})
