package domain_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/domain"
	"github.com/sirockin/cucumber-screenplay-go/features"
)

func TestDomainFeatures(t *testing.T){
	features.Test(t, domain.New(), []string{"../features"})
}
