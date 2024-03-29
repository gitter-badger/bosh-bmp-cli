package cmds_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cmds "github.com/maximilien/bosh-bmp-cli/cmds"
)

var _ = Describe("task command", func() {
	var (
		args    []string
		options cmds.Options
		cmd     cmds.Command
	)

	BeforeEach(func() {
		args = []string{"bmp", "task"}
		options = cmds.Options{
			Verbose: false,
		}

		cmd = cmds.NewTaskCommand(options)
	})

	Describe("NewTaskCommand", func() {
		It("create new TaskCommand", func() {
			Expect(cmd).ToNot(BeNil())

			cmd2 := cmds.NewTaskCommand(options)
			Expect(cmd2).ToNot(BeNil())
			Expect(cmd2).To(Equal(cmd))
		})
	})

	Describe("#Name", func() {
		It("returns the name of a TaskCommand", func() {
			Expect(cmd.Name()).To(Equal("task"))
		})
	})

	Describe("#Description", func() {
		It("returns the description of a TaskCommand", func() {
			Expect(cmd.Description()).To(Equal(`Show the output of the task: \"option --debug, Get the debug info of the task\"`))
		})
	})

	Describe("#Usage", func() {
		It("returns the usage text of a TaskCommand", func() {
			Expect(cmd.Usage()).To(Equal("bmp task <task-id>"))
		})
	})

	Describe("#Validate", func() {
		It("validates a good TaskCommand", func() {
			validate, err := cmd.Validate()
			Expect(validate).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("#Execute", func() {
		It("executes a good TaskCommand", func() {
			rc, err := cmd.Execute(args)
			Expect(rc).To(Equal(0))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
