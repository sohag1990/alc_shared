package models

import "time"

type Fra struct {
	DefaultProperties
	Name                            string
	Address                         string
	Description                     string
	NumberOfFloors                  int
	ConstructionDate                time.Time
	BriefDescriptionOfAccommodation string

	LandlordName      string
	LandlordAddress   string
	ResponsiblePerson string
	ContactNumber     string
	Email             string

	SleepingOccupants string
	DisableOccupants  string
	ChildrenOccupants string
	ElderlyOccupants  string

	Assessmentdate    time.Time
	AssessmentType    string
	AssessedBy        string
	Registration      string
	AssessmentPurpose string
	ReviewDate        string

	ProbabilityScore         string
	ProbabilityDescription   string
	PotentialRiskScore       string
	PotentialRiskDescription string

	// Hazard identification
	// Doors
	Doors_1_0_Comment        string
	Doors_1_0_Recommendation string

	SufficientExits_1_1                   string
	SufficientExits_1_1_RiskLevel         string
	ImmediatelyOpened_1_2                 string
	ImmediatelyOpened_1_2_RiskLevel       string
	AppropriateDevices_1_3                string
	AppropriateDevices_1_3_RiskLevel      string
	InternalDoorsWedged_1_4               string
	InternalDoorsWedged_1_4_RiskLevel     string
	AutomaticDoorFastenings_1_5           string
	AutomaticDoorFastenings_1_5_RiskLevel string

	// Signage
	Signage_2_0_Comment        string
	Signage_2_0_Recommendation string

	SufficientExitSignage_2_1            string
	SufficientExitSignage_2_1_Risklevel  string
	EscapeRoutesIdentified_2_2           string
	EscapeRoutesIdentified_2_2_Risklevel string
	SignageComply_2_3                    string
	SignageComply_2_3_Risklevel          string
	AppropriateLabel_2_4                 string
	AppropriateLabel_2_4_Risklevel       string
	ExternalSignage_2_5                  string
	ExternalSignage_2_5_Risklevel        string
	AdequateNotices_2_6                  string
	AdequateNotices_2_6_Risklevel        string
	IlluminatedExitSigns                 string
	IlluminatedExitSigns_Risklevel       string

	// Lighting
	Lighting_3_0_Comment        string
	Lighting_3_0_Recommendation string

	EmergencyLighting_3_1                           string
	EmergencyLighting_3_1_Risklevel                 string
	EscapeRoutesAdequatelyIlluminated_3_2           string
	EscapeRoutesAdequatelyIlluminated_3_2_Risklevel string
	ExistingArtificialLighting_3_3                  string
	ExistingArtificialLighting_3_3_Risklevel        string
	EmergencyLightingCharging_3_4                   string
	EmergencyLightingCharging_3_4_Risklevel         string
	LightingMaintained_3_5                          string
	LightingMaintained_3_5_Risklevel                string

	// Electrical
	Electrical_4_0_Comment        string
	Electrical_4_0_Recommendation string

	ReasonableMeasures_4_1               string
	ReasonableMeasures_4_1_Risklevel     string
	ElectricalInstallation_4_2           string
	ElectricalInstallation_4_2_Risklevel string
	ElectricalEquipment_4_3              string
	ElectricalEquipment_4_3_Risklevel    string

	// Smoking
	Smoking_5_0_Comment        string
	Smoking_5_0_Recommendation string

	PeopleSmoke_5_1                    string
	PeopleSmoke_5_1_Risklevel          string
	SuitableArrangements_5_2           string
	SuitableArrangements_5_2_Risklevel string
	SignageAvailable_5_3               string
	SignageAvailable_5_3_Risklevel     string

	// Portable heating & heating installations
	PortableHeating_6_0_Comment        string
	PortableHeating_6_0_Recommendation string

	PortableHeatersAvoided_6_1           string
	PortableHeatersAvoided_6_1_Risklevel string
	SuitableMeasures_6_2                 string
	SuitableMeasures_6_2_Risklevel       string
	RegularMaintenance_6_3               string
	RegularMaintenance_6_3_Risklevel     string
}