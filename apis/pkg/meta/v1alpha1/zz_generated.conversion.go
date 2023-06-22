// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/crossplane/crossplane/apis/pkg/meta/v1"
	v11 "k8s.io/api/rbac/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GeneratedFromHubConverter struct{}

func (c *GeneratedFromHubConverter) Configuration(source *v1.Configuration) *Configuration {
	var pV1alpha1Configuration *Configuration
	if source != nil {
		var v1alpha1Configuration Configuration
		v1alpha1Configuration.TypeMeta = c.v1TypeMetaToV1TypeMeta((*source).TypeMeta)
		v1alpha1Configuration.ObjectMeta = ConvertObjectMeta((*source).ObjectMeta)
		v1alpha1Configuration.Spec = c.v1ConfigurationSpecToV1alpha1ConfigurationSpec((*source).Spec)
		pV1alpha1Configuration = &v1alpha1Configuration
	}
	return pV1alpha1Configuration
}
func (c *GeneratedFromHubConverter) Provider(source *v1.Provider) *Provider {
	var pV1alpha1Provider *Provider
	if source != nil {
		var v1alpha1Provider Provider
		v1alpha1Provider.TypeMeta = c.v1TypeMetaToV1TypeMeta((*source).TypeMeta)
		v1alpha1Provider.ObjectMeta = ConvertObjectMeta((*source).ObjectMeta)
		v1alpha1Provider.Spec = c.v1ProviderSpecToV1alpha1ProviderSpec((*source).Spec)
		pV1alpha1Provider = &v1alpha1Provider
	}
	return pV1alpha1Provider
}
func (c *GeneratedFromHubConverter) pV1CrossplaneConstraintsToPV1alpha1CrossplaneConstraints(source *v1.CrossplaneConstraints) *CrossplaneConstraints {
	var pV1alpha1CrossplaneConstraints *CrossplaneConstraints
	if source != nil {
		var v1alpha1CrossplaneConstraints CrossplaneConstraints
		v1alpha1CrossplaneConstraints.Version = (*source).Version
		pV1alpha1CrossplaneConstraints = &v1alpha1CrossplaneConstraints
	}
	return pV1alpha1CrossplaneConstraints
}
func (c *GeneratedFromHubConverter) v1ConfigurationSpecToV1alpha1ConfigurationSpec(source v1.ConfigurationSpec) ConfigurationSpec {
	var v1alpha1ConfigurationSpec ConfigurationSpec
	v1alpha1ConfigurationSpec.MetaSpec = c.v1MetaSpecToV1alpha1MetaSpec(source.MetaSpec)
	return v1alpha1ConfigurationSpec
}
func (c *GeneratedFromHubConverter) v1ControllerSpecToV1alpha1ControllerSpec(source v1.ControllerSpec) ControllerSpec {
	var v1alpha1ControllerSpec ControllerSpec
	var pString *string
	if source.Image != nil {
		xstring := *source.Image
		pString = &xstring
	}
	v1alpha1ControllerSpec.Image = pString
	var v1PolicyRuleList []v11.PolicyRule
	if source.PermissionRequests != nil {
		v1PolicyRuleList = make([]v11.PolicyRule, len(source.PermissionRequests))
		for i := 0; i < len(source.PermissionRequests); i++ {
			v1PolicyRuleList[i] = c.v1PolicyRuleToV1PolicyRule(source.PermissionRequests[i])
		}
	}
	v1alpha1ControllerSpec.PermissionRequests = v1PolicyRuleList
	return v1alpha1ControllerSpec
}
func (c *GeneratedFromHubConverter) v1DependencyToV1alpha1Dependency(source v1.Dependency) Dependency {
	var v1alpha1Dependency Dependency
	var pString *string
	if source.Provider != nil {
		xstring := *source.Provider
		pString = &xstring
	}
	v1alpha1Dependency.Provider = pString
	var pString2 *string
	if source.Configuration != nil {
		xstring2 := *source.Configuration
		pString2 = &xstring2
	}
	v1alpha1Dependency.Configuration = pString2
	v1alpha1Dependency.Version = source.Version
	return v1alpha1Dependency
}
func (c *GeneratedFromHubConverter) v1MetaSpecToV1alpha1MetaSpec(source v1.MetaSpec) MetaSpec {
	var v1alpha1MetaSpec MetaSpec
	v1alpha1MetaSpec.Crossplane = c.pV1CrossplaneConstraintsToPV1alpha1CrossplaneConstraints(source.Crossplane)
	var v1alpha1DependencyList []Dependency
	if source.DependsOn != nil {
		v1alpha1DependencyList = make([]Dependency, len(source.DependsOn))
		for i := 0; i < len(source.DependsOn); i++ {
			v1alpha1DependencyList[i] = c.v1DependencyToV1alpha1Dependency(source.DependsOn[i])
		}
	}
	v1alpha1MetaSpec.DependsOn = v1alpha1DependencyList
	return v1alpha1MetaSpec
}
func (c *GeneratedFromHubConverter) v1PolicyRuleToV1PolicyRule(source v11.PolicyRule) v11.PolicyRule {
	var v1PolicyRule v11.PolicyRule
	var stringList []string
	if source.Verbs != nil {
		stringList = make([]string, len(source.Verbs))
		for i := 0; i < len(source.Verbs); i++ {
			stringList[i] = source.Verbs[i]
		}
	}
	v1PolicyRule.Verbs = stringList
	var stringList2 []string
	if source.APIGroups != nil {
		stringList2 = make([]string, len(source.APIGroups))
		for j := 0; j < len(source.APIGroups); j++ {
			stringList2[j] = source.APIGroups[j]
		}
	}
	v1PolicyRule.APIGroups = stringList2
	var stringList3 []string
	if source.Resources != nil {
		stringList3 = make([]string, len(source.Resources))
		for k := 0; k < len(source.Resources); k++ {
			stringList3[k] = source.Resources[k]
		}
	}
	v1PolicyRule.Resources = stringList3
	var stringList4 []string
	if source.ResourceNames != nil {
		stringList4 = make([]string, len(source.ResourceNames))
		for l := 0; l < len(source.ResourceNames); l++ {
			stringList4[l] = source.ResourceNames[l]
		}
	}
	v1PolicyRule.ResourceNames = stringList4
	var stringList5 []string
	if source.NonResourceURLs != nil {
		stringList5 = make([]string, len(source.NonResourceURLs))
		for m := 0; m < len(source.NonResourceURLs); m++ {
			stringList5[m] = source.NonResourceURLs[m]
		}
	}
	v1PolicyRule.NonResourceURLs = stringList5
	return v1PolicyRule
}
func (c *GeneratedFromHubConverter) v1ProviderSpecToV1alpha1ProviderSpec(source v1.ProviderSpec) ProviderSpec {
	var v1alpha1ProviderSpec ProviderSpec
	v1alpha1ProviderSpec.Controller = c.v1ControllerSpecToV1alpha1ControllerSpec(source.Controller)
	v1alpha1ProviderSpec.MetaSpec = c.v1MetaSpecToV1alpha1MetaSpec(source.MetaSpec)
	return v1alpha1ProviderSpec
}
func (c *GeneratedFromHubConverter) v1TypeMetaToV1TypeMeta(source v12.TypeMeta) v12.TypeMeta {
	var v1TypeMeta v12.TypeMeta
	v1TypeMeta.Kind = source.Kind
	v1TypeMeta.APIVersion = source.APIVersion
	return v1TypeMeta
}

type GeneratedToHubConverter struct{}

func (c *GeneratedToHubConverter) Configuration(source *Configuration) *v1.Configuration {
	var pV1Configuration *v1.Configuration
	if source != nil {
		var v1Configuration v1.Configuration
		v1Configuration.TypeMeta = c.v1TypeMetaToV1TypeMeta((*source).TypeMeta)
		v1Configuration.ObjectMeta = ConvertObjectMeta((*source).ObjectMeta)
		v1Configuration.Spec = c.v1alpha1ConfigurationSpecToV1ConfigurationSpec((*source).Spec)
		pV1Configuration = &v1Configuration
	}
	return pV1Configuration
}
func (c *GeneratedToHubConverter) Provider(source *Provider) *v1.Provider {
	var pV1Provider *v1.Provider
	if source != nil {
		var v1Provider v1.Provider
		v1Provider.TypeMeta = c.v1TypeMetaToV1TypeMeta((*source).TypeMeta)
		v1Provider.ObjectMeta = ConvertObjectMeta((*source).ObjectMeta)
		v1Provider.Spec = c.v1alpha1ProviderSpecToV1ProviderSpec((*source).Spec)
		pV1Provider = &v1Provider
	}
	return pV1Provider
}
func (c *GeneratedToHubConverter) pV1alpha1CrossplaneConstraintsToPV1CrossplaneConstraints(source *CrossplaneConstraints) *v1.CrossplaneConstraints {
	var pV1CrossplaneConstraints *v1.CrossplaneConstraints
	if source != nil {
		var v1CrossplaneConstraints v1.CrossplaneConstraints
		v1CrossplaneConstraints.Version = (*source).Version
		pV1CrossplaneConstraints = &v1CrossplaneConstraints
	}
	return pV1CrossplaneConstraints
}
func (c *GeneratedToHubConverter) v1PolicyRuleToV1PolicyRule(source v11.PolicyRule) v11.PolicyRule {
	var v1PolicyRule v11.PolicyRule
	var stringList []string
	if source.Verbs != nil {
		stringList = make([]string, len(source.Verbs))
		for i := 0; i < len(source.Verbs); i++ {
			stringList[i] = source.Verbs[i]
		}
	}
	v1PolicyRule.Verbs = stringList
	var stringList2 []string
	if source.APIGroups != nil {
		stringList2 = make([]string, len(source.APIGroups))
		for j := 0; j < len(source.APIGroups); j++ {
			stringList2[j] = source.APIGroups[j]
		}
	}
	v1PolicyRule.APIGroups = stringList2
	var stringList3 []string
	if source.Resources != nil {
		stringList3 = make([]string, len(source.Resources))
		for k := 0; k < len(source.Resources); k++ {
			stringList3[k] = source.Resources[k]
		}
	}
	v1PolicyRule.Resources = stringList3
	var stringList4 []string
	if source.ResourceNames != nil {
		stringList4 = make([]string, len(source.ResourceNames))
		for l := 0; l < len(source.ResourceNames); l++ {
			stringList4[l] = source.ResourceNames[l]
		}
	}
	v1PolicyRule.ResourceNames = stringList4
	var stringList5 []string
	if source.NonResourceURLs != nil {
		stringList5 = make([]string, len(source.NonResourceURLs))
		for m := 0; m < len(source.NonResourceURLs); m++ {
			stringList5[m] = source.NonResourceURLs[m]
		}
	}
	v1PolicyRule.NonResourceURLs = stringList5
	return v1PolicyRule
}
func (c *GeneratedToHubConverter) v1TypeMetaToV1TypeMeta(source v12.TypeMeta) v12.TypeMeta {
	var v1TypeMeta v12.TypeMeta
	v1TypeMeta.Kind = source.Kind
	v1TypeMeta.APIVersion = source.APIVersion
	return v1TypeMeta
}
func (c *GeneratedToHubConverter) v1alpha1ConfigurationSpecToV1ConfigurationSpec(source ConfigurationSpec) v1.ConfigurationSpec {
	var v1ConfigurationSpec v1.ConfigurationSpec
	v1ConfigurationSpec.MetaSpec = c.v1alpha1MetaSpecToV1MetaSpec(source.MetaSpec)
	return v1ConfigurationSpec
}
func (c *GeneratedToHubConverter) v1alpha1ControllerSpecToV1ControllerSpec(source ControllerSpec) v1.ControllerSpec {
	var v1ControllerSpec v1.ControllerSpec
	var pString *string
	if source.Image != nil {
		xstring := *source.Image
		pString = &xstring
	}
	v1ControllerSpec.Image = pString
	var v1PolicyRuleList []v11.PolicyRule
	if source.PermissionRequests != nil {
		v1PolicyRuleList = make([]v11.PolicyRule, len(source.PermissionRequests))
		for i := 0; i < len(source.PermissionRequests); i++ {
			v1PolicyRuleList[i] = c.v1PolicyRuleToV1PolicyRule(source.PermissionRequests[i])
		}
	}
	v1ControllerSpec.PermissionRequests = v1PolicyRuleList
	return v1ControllerSpec
}
func (c *GeneratedToHubConverter) v1alpha1DependencyToV1Dependency(source Dependency) v1.Dependency {
	var v1Dependency v1.Dependency
	var pString *string
	if source.Provider != nil {
		xstring := *source.Provider
		pString = &xstring
	}
	v1Dependency.Provider = pString
	var pString2 *string
	if source.Configuration != nil {
		xstring2 := *source.Configuration
		pString2 = &xstring2
	}
	v1Dependency.Configuration = pString2
	v1Dependency.Version = source.Version
	return v1Dependency
}
func (c *GeneratedToHubConverter) v1alpha1MetaSpecToV1MetaSpec(source MetaSpec) v1.MetaSpec {
	var v1MetaSpec v1.MetaSpec
	v1MetaSpec.Crossplane = c.pV1alpha1CrossplaneConstraintsToPV1CrossplaneConstraints(source.Crossplane)
	var v1DependencyList []v1.Dependency
	if source.DependsOn != nil {
		v1DependencyList = make([]v1.Dependency, len(source.DependsOn))
		for i := 0; i < len(source.DependsOn); i++ {
			v1DependencyList[i] = c.v1alpha1DependencyToV1Dependency(source.DependsOn[i])
		}
	}
	v1MetaSpec.DependsOn = v1DependencyList
	return v1MetaSpec
}
func (c *GeneratedToHubConverter) v1alpha1ProviderSpecToV1ProviderSpec(source ProviderSpec) v1.ProviderSpec {
	var v1ProviderSpec v1.ProviderSpec
	v1ProviderSpec.Controller = c.v1alpha1ControllerSpecToV1ControllerSpec(source.Controller)
	v1ProviderSpec.MetaSpec = c.v1alpha1MetaSpecToV1MetaSpec(source.MetaSpec)
	return v1ProviderSpec
}
