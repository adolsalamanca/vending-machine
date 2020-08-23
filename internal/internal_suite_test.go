package internal_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/MetalBlueberry/golandreporter"
)

func TestInternal(t *testing.T) {
	RegisterFailHandler(Fail)
	golandReporter := golandreporter.NewGolandReporter()
	RunSpecsWithDefaultAndCustomReporters(t, "Internal Suite", []Reporter{golandReporter})
}
