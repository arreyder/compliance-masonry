package diff

import (
	"fmt"
	"github.com/opencontrol/compliance-masonry/lib"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
)

// Inventory maintains the inventory of all the controls within a given workspace.
type Inventory struct {
	common.Workspace
	masterControlList       map[string]common.Control
	actualSatisfiedControls map[string]common.Satisfies
	MissingControlList      map[string]common.Control
}

// retrieveMasterControlsList will gather the list of controls needed for a given certification.
func (i *Inventory) retrieveMasterControlsList() {
	standardKeys := i.GetCertification().GetSortedStandards()
	for _, standardKey := range standardKeys {
		for _, controlKey := range i.GetCertification().GetControlKeysFor(standardKey) {
			key := standardAndControlString(standardKey, controlKey)
			if _, exists := i.masterControlList[key]; !exists {
				standard, found := i.GetStandard(standardKey)
				if !found {
					continue
				}
				i.masterControlList[key] = standard.GetControl(controlKey)
			}
		}
	}
}

// findCompletedControls will find the list of all documented, but not completed controls found within the workspace.
func (i *Inventory) findCompletedControls() {
	for _, component := range i.GetAllComponents() {
		for _, satisfiedControl := range component.GetAllSatisfies() {
			// We're looking for components that we've not established some level of compliance for.
			if "completed" != satisfiedControl.GetImplementationStatus() {
				key := standardAndControlString(satisfiedControl.GetStandardKey(), satisfiedControl.GetControlKey())
				if _, exists := i.actualSatisfiedControls[key]; !exists {
					i.actualSatisfiedControls[key] = satisfiedControl
				}
			}
		}
	}
}

// findDocumentedControls will find the list of all documented controls found within the workspace.
func (i *Inventory) findDocumentedControls() {
	for _, component := range i.GetAllComponents() {
		for _, satisfiedControl := range component.GetAllSatisfies() {
			key := standardAndControlString(satisfiedControl.GetStandardKey(), satisfiedControl.GetControlKey())
			if _, exists := i.actualSatisfiedControls[key]; !exists {
				i.actualSatisfiedControls[key] = satisfiedControl
			}
		}
	}
}

// calculateNonDocumentedControls will compute the diff between the master list of controls and the documented controls.
func (i *Inventory) calculateNonDocumentedControls() {
	for standardAndControlKey, control := range i.masterControlList {
		if _, exists := i.actualSatisfiedControls[standardAndControlKey]; !exists {
			i.MissingControlList[standardAndControlKey] = control
		}
	}
}

// standardAndControlString makes a string from the standard and the control.
// This is helpful for functions that want to create unique keys consistently.
func standardAndControlString(standard string, control string) string {
	return standard + "@" + control
}

// Config contains the settings for how to compute the gap analysis
type Config struct {
	Certification  string
	OpencontrolDir string
}

// ComputeGapAnalysis will compute the gap analysis and return the inventory of the controls for the
// opencontrol workspace if successful. Otherwise, it will return a list of error messages.
// TODO: fix the error return to return of type error. This was used because existing code returned that type
// TODO: e.g. GetCertification
func ComputeGapAnalysis(config Config) (Inventory, []error) {
	// Initialize inventory with certification
	certificationPath, errs := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" {
		return Inventory{}, errs
	}
	workspace, _ := lib.LoadData(config.OpencontrolDir, certificationPath)
	i := Inventory{
		Workspace:               workspace,
		masterControlList:       make(map[string]common.Control),
		actualSatisfiedControls: make(map[string]common.Satisfies),
		MissingControlList:      make(map[string]common.Control),
	}
	if i.GetCertification() == nil || i.GetAllComponents() == nil {
		return Inventory{}, []error{fmt.Errorf("Unable to load data in %s for certification %s", config.OpencontrolDir, config.Certification)}
	}

	// Gather list of all controls for certification
	i.retrieveMasterControlsList()
	// Find the documented controls.
	i.findDocumentedControls()
	// Calculate the Missing controls / Non documented
	i.calculateNonDocumentedControls()

	return i, nil
}

func ComputeGapAnalysisTODO(config Config) (Inventory, []error) {
	// Initialize inventory with certification
	certificationPath, errs := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" {
		return Inventory{}, errs
	}
	workspace, _ := lib.LoadData(config.OpencontrolDir, certificationPath)
	i := Inventory{
		Workspace:               workspace,
		masterControlList:       make(map[string]common.Control),
		actualSatisfiedControls: make(map[string]common.Satisfies),
		MissingControlList:      make(map[string]common.Control),
	}
	if i.GetCertification() == nil || i.GetAllComponents() == nil {
		return Inventory{}, []error{fmt.Errorf("Unable to load data in %s for certification %s", config.OpencontrolDir, config.Certification)}
	}

	// Gather list of all controls for certification
	i.retrieveMasterControlsList()
	// Find the documented controls.
	i.findCompletedControls()
	// Calculate the Missing controls / Non documented
	i.calculateNonDocumentedControls()

	return i, nil
}
