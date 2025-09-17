package features_test

import (
	"testing"

	"github.com/sirockin/cucumber-screenplay-go/features"
	"github.com/sirockin/cucumber-screenplay-go/features/driver/domain"
)

func TestDomainFeatures(t *testing.T){
	features.Test(t, domain.New(), []string{"."})
}
