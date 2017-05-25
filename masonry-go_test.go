package main_test

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Masonry CLI", func() {
	Describe("When the CLI is run with no commands", func() {
		It("should list the available commands", func() {
			output := Masonry("", "")
			Eventually(output.Out.Contents).Should(ContainSubstring("Install compliance dependencies"))
			Eventually(output.Out.Contents).Should(ContainSubstring("Create Documentation"))
		})
	})

	Describe("Base Docs Commands", func() {
		Describe("When the CLI is run with the docs command", func() {
			It("should list the available doc commands", func() {
				output := Masonry("docs", "")
				Eventually(output.Out.Contents).Should(ContainSubstring("Create Gitbook Documentation"))
			})
		})
	})

	Describe("Gitbook Docs Commands", func() {

		var exportTempDir string
		BeforeEach(func() {
			exportTempDir, _ = ioutil.TempDir("", "exports")
		})

		Describe("Gitbook Commands", func() {
			Describe("When the CLI is run with the `docs gitbook` command", func() {
				It("should let the user know that they have not described a certification and show how to use the command", func() {
					output := Masonry("docs", "gitbook")
					Eventually(output.Err.Contents).Should(ContainSubstring("Error: Missing Certification Argument\n"))
				})
			})

			Describe("When the CLI is run with the `docs gitbook` command without opencontrols dir", func() {
				It("should let the user know that there is no opencontrols/certifications directory", func() {
					output := Masonry("docs", "gitbook", "LATO")
					Eventually(output.Err.Contents).Should(ContainSubstring("Error: `" + filepath.Join("opencontrols", "certifications") + "` directory does exist\n"))
				})
			})
		})

		Describe("When the CLI is run with the `docs gitbook` command with a certification and no markdown", func() {
			It("should create the documentation but warn users that there is no markdown dir", func() {
				output := Masonry(
					"docs", "gitbook", "LATO",
					"-e", exportTempDir,
					"-o", filepath.Join("fixtures", "opencontrol_fixtures"),
					"-m", "sdfds").Wait(1 * time.Second)
				Eventually(output.Out.Contents).Should(ContainSubstring("Warning: markdown directory does not exist\n"))
				Eventually(output.Out.Contents).Should(ContainSubstring("New Gitbook Documentation Created\n"))
			})
		})

		Describe("When the CLI is run with the `docs gitbook` command with a certification", func() {
			It("should create the documentation without warning the user", func() {
				exportTempDir, _ := ioutil.TempDir("", "exports")
				output := Masonry(
					"docs", "gitbook", "LATO",
					"-e", exportTempDir,
					"-o", filepath.Join("fixtures", "opencontrol_fixtures_with_markdown"),
					"-m", filepath.Join("fixtures", "opencontrol_fixtures_with_markdown", "markdowns")).Wait(1 * time.Second)
				Eventually(output.Out.Contents).ShouldNot(ContainSubstring("Warning: markdown directory does not exist\n"))
				Eventually(output.Out.Contents).Should(ContainSubstring("New Gitbook Documentation Created\n"))
			})
		})
		AfterEach(func() {
			_ = os.RemoveAll(exportTempDir)
		})
	})

	Describe("Diff Commands", func() {
		Describe("When the diff command is run", func() {
			It("should let the user know that they have not described a certification and show how to use the command", func() {
				output := Masonry("diff")
				Eventually(output.Err.Contents).Should(ContainSubstring("Error: Missing Certification Argument\n"))
			})
		})
		Describe("When the CLI is run with the `diff` command without opencontrols dir", func() {
			It("should let the user know that there is no opencontrols/certifications directory", func() {
				output := Masonry("diff", "LATO")
				Eventually(output.Err.Contents).Should(ContainSubstring("Error: `" + filepath.Join("opencontrols", "certifications") + "` directory does exist\n"))
			})
		})
		Describe("When the CLI is run with the `diff` command with a certification", func() {
			It("should print the number of missing controls", func() {
				output := Masonry(
					"diff", "LATO",
					"-o", filepath.Join("fixtures", "opencontrol_fixtures")).Wait(1 * time.Second)
				Eventually(output.Out.Contents).Should(ContainSubstring("Number of missing controls:"))
			})
		})
	})
})

func Masonry(args ...string) *Session {
	path, err := Build("github.com/opencontrol/compliance-masonry")
	Expect(err).NotTo(HaveOccurred())
	cmd := exec.Command(path, args...)
	stdin, err := cmd.StdinPipe()
	Expect(err).ToNot(HaveOccurred())
	buffer := bufio.NewWriter(stdin)
	_, _ = buffer.WriteString(strings.Join(args, " "))
	_ = buffer.Flush()
	session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}
