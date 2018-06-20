package myGinkgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/myGinkgo/book"
)

func TestMyGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MyGinkgo Suite")
}
