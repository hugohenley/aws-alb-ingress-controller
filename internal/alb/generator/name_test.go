package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NameLB(t *testing.T) {
	nameGen := NameGenerator{
		ALBNamePrefix: "albprefix",
	}

	assert.Equal(t, nameGen.NameLB("namespace", "ingress"), "albprefix-namespace-ingres-1829")
}

func Test_TG(t *testing.T) {
	nameGen := NameGenerator{
		ALBNamePrefix: "albprefix",
	}

	assert.Equal(t, nameGen.NameTG("namespace", "ingress", "svc", "80", "instance", "tcp"), "albprefix-40fd833d587e7bd5076")
}

func Test_NameLBSG(t *testing.T) {
	nameGen := NameGenerator{
		ALBNamePrefix: "albprefix",
	}

	assert.Equal(t, nameGen.NameLBSG("namespace", "ingress"), "albprefix-namespace-ingres-1829")
}

func Test_NameInstanceSG(t *testing.T) {
	nameGen := NameGenerator{
		ALBNamePrefix: "albprefix",
	}

	assert.Equal(t, nameGen.NameInstanceSG("namespace", "ingress"), "instance-albprefix-namespace-ingres-1829")
}
