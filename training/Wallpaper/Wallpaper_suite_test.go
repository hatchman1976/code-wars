package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWallpaper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wallpaper Suite")
}
