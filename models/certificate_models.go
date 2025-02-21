package models

import (
	"html/template"
	"time"
)

type FollowUp struct {
	DefaultProperties
	Company         string `gorm:"size:100"`
	Email           string `gorm:"size:100"`
	OrderID         uint64
	FollowUpOption  string `gorm:"size:100"`
	Services        string `gorm:"size:250"`
	FollowUpSubject string `gorm:"size:150"`
	FollowUpContent template.HTML
	FeedbackContent template.HTML
	RenewedOrderID  uint64
	Status          bool
}
type EmergencyLightingCert struct {
	DefaultProperties
	OrderID  uint64
	Name     string `gorm:"size:100"`
	Email    string `gorm:"size:100"`
	Address  string `gorm:"size:250"`
	PostCode string `gorm:"size:20"`
	Phone    string `gorm:"size:20"`

	LandlordName     string    `gorm:"size:100"`
	LandlordAddress  string    `gorm:"size:250"`
	LandlordPostCode string    `gorm:"size:20"`
	AssessedBy       string    `gorm:"size:100"`
	Registration     string    `gorm:"size:100"`
	AssessmentDate   time.Time `form:"AssessmentDate"  time_format:"2006-01-02"`
	ReviewDate       time.Time `form:"ReviewDate"  time_format:"2006-01-02"`

	CoveredThisReport string `gorm:"size:550"`
	RelatedReference  string `gorm:"size:550"`
	Validity          string `gorm:"size:50"`
}

type PatCert struct {
	DefaultProperties
	OrderID  uint64
	PatItems []PatItem
	Name     string `gorm:"size:100"`
	Email    string `gorm:"size:100"`
	Address  string `gorm:"size:250"`
	PostCode string `gorm:"size:20"`
	Phone    string `gorm:"size:20"`

	LandlordName     string    `gorm:"size:100"`
	LandlordAddress  string    `gorm:"size:250"`
	LandlordPostCode string    `gorm:"size:20"`
	AssessedBy       string    `gorm:"size:100"`
	Registration     string    `gorm:"size:100"`
	AssessmentDate   time.Time `form:"AssessmentDate"  time_format:"2006-01-02"`
	ReviewDate       time.Time `form:"ReviewDate"  time_format:"2006-01-02"`
}
type PatItem struct {
	DefaultProperties
	PatCertID uint64
	Location  string `gorm:"size:50"`
	Name      string `gorm:"size:100"`
}

type FireAlarmCert struct {
	DefaultProperties
	OrderID  uint64
	Name     string `gorm:"size:100"`
	Email    string `gorm:"size:100"`
	Address  string `gorm:"size:250"`
	PostCode string `gorm:"size:20"`
	Phone    string `gorm:"size:20"`

	LandlordName     string    `gorm:"size:100"`
	LandlordAddress  string    `gorm:"size:250"`
	LandlordPostCode string    `gorm:"size:20"`
	AssessedBy       string    `gorm:"size:100"`
	Registration     string    `gorm:"size:100"`
	AssessmentDate   time.Time `form:"AssessmentDate"  time_format:"2006-01-02"`
	ReviewDate       time.Time `form:"ReviewDate"  time_format:"2006-01-02"`

	SystemDetails     string `gorm:"size:550"`
	CoveredThisReport string `gorm:"size:550"`
	Agreed            string `gorm:"size:550"`
	Variations        string `gorm:"size:550"`
	GeneralCondition  string `gorm:"size:550"`
	Validity          string `gorm:"size:50"`
}
type GasCert struct {
	DefaultProperties
	GasRecommendations []GasRecommendation
	GasInspections     []GasInspection
	CertificateIssuer  string
	OrderID            uint64
	Name               string `gorm:"size:100"`
	Email              string `gorm:"size:100"`
	Address            string `gorm:"size:250"`
	PostCode           string `gorm:"size:20"`
	Phone              string `gorm:"size:20"`

	LandlordName     string `gorm:"size:100"`
	LandlordAddress  string `gorm:"size:250"`
	LandlordPostCode string `gorm:"size:20"`

	GasEngineerName     string `gorm:"size:100"`
	GasRegistraionID    string `gorm:"size:20"`
	GasSafeRegNo        string `gorm:"size:20"`
	GasEngineerAddress  string `gorm:"size:250"`
	GasEngineerPostCode string `gorm:"size:20"`
	GasEngineerTel      string `gorm:"size:20"`

	AssessmentDate time.Time `form:"AssessmentDate"  time_format:"2006-01-02"`
	ReviewDate     time.Time `form:"ReviewDate"  time_format:"2006-01-02"`

	ReminderSentDate time.Time `form:"ReminderSentDate"  time_format:"2006-01-02"`
	RenewalStatus    bool
	RenewedDate      time.Time `form:"RenewedDate"  time_format:"2006-01-02"`
	RenewdOrderID    uint64

	// APPLIANCE DETAILS

}
type GasRecommendation struct {
	DefaultProperties
	GasCertID           uint64
	DefectIdentified    string `gorm:"size:300"`
	RemedialActionTaken string `gorm:"size:300"`
	NoticeLabelIssued   string `gorm:"size:50"`
}
type GasInspection struct {
	DefaultProperties
	GasCertID                        uint64
	Location                         string `gorm:"size:50"`
	ApplianceType                    string `gorm:"size:50"`
	Make                             string `gorm:"size:50"`
	Model                            string `gorm:"size:50"`
	FlueType                         string `gorm:"size:50"`
	LandlordsAppliance               string `gorm:"size:50"`
	ApplianceInspected               string `gorm:"size:50"`
	CombustionAnalyserReading        string `gorm:"size:50"`
	OperatingPressure                string `gorm:"size:50"`
	SafetyDevicesOperation           string `gorm:"size:50"`
	VentilationProvisionSatisfactory string `gorm:"size:50"`
	VisualConditionSatisfactory      string `gorm:"size:50"`
	FluePerformanceTest              string `gorm:"size:50"`
	ApplianceServiced                string `gorm:"size:50"`
	ApplianceSafeToUse               string `gorm:"size:50"`
	COAlarmFitted                    string `gorm:"size:50"`
	IsCOAlarmInDate                  string `gorm:"size:50"`
	TestingCOAlarmSatisfactory       string `gorm:"size:50"`
}

type Fra struct {
	DefaultProperties
	OrderID                             uint64
	FraMaintenanceRecordKeepingID       uint64
	FraMeansForGivingWarningID          uint64
	FraEmergencyActionPlanID            uint64
	FraPortableFireFightingEquipmentID  uint64
	FraFixedInstallationID              uint64
	FraArsonID                          uint64
	FraGeneralID                        uint64
	FraRecommendations                  []FraRecommendation
	FraHazardIdentificationDoorID       uint64
	FraHazardIdentificationSignageID    uint64
	FraHazardIdentificationLightingID   uint64
	FraHazardIdentificationElectricalID uint64
	FraHazardIdentificationCookingID    uint64

	FraHazardIdentificationDoor       FraHazardIdentificationDoor
	FraHazardIdentificationSignage    FraHazardIdentificationSignage
	FraHazardIdentificationLighting   FraHazardIdentificationLighting
	FraHazardIdentificationElectrical FraHazardIdentificationElectrical
	FraHazardIdentificationCooking    FraHazardIdentificationCooking

	FraMaintenanceRecordKeeping      FraMaintenanceRecordKeeping
	FraMeansForGivingWarning         FraMeansForGivingWarning
	FraEmergencyActionPlan           FraEmergencyActionPlan
	FraPortableFireFightingEquipment FraPortableFireFightingEquipment
	FraFixedInstallation             FraFixedInstallation
	FraArson                         FraArson
	FraGeneral                       FraGeneral

	// Order                           Order
	Name                            string `gorm:"size:255"`
	Address                         string `gorm:"size:255"`
	Description                     string `gorm:"size:512"`
	PremisesDescription             string `gorm:"size:512"`
	NumberOfFloors                  int
	ConstructionDate                time.Time `form:"ConstructionDate"  time_format:"2006-01-02"`
	BriefDescriptionOfAccommodation string

	LandlordName      string `gorm:"size:255"`
	LandlordAddress   string `gorm:"size:255"`
	ResponsiblePerson string `gorm:"size:255"`
	ContactNumber     string `gorm:"size:255"`
	Email             string `gorm:"size:255"`

	SleepingOccupants string `gorm:"size:100"`
	DisableOccupants  string `gorm:"size:100"`
	ChildrenOccupants string `gorm:"size:100"`
	ElderlyOccupants  string `gorm:"size:100"`

	AssessmentDate    time.Time `form:"AssessmentDate"  time_format:"2006-01-02"`
	AssessmentType    string    `gorm:"size:100"`
	AssessedBy        string    `gorm:"size:100"`
	Registration      string    `gorm:"size:100"`
	AssessmentPurpose string    `gorm:"type:text"`
	ReviewDate        time.Time `form:"ReviewDate"  time_format:"2006-01-02"`

	PotentialRiskScore string `gorm:"size:24"`
}
type FraMeansForGivingWarning struct {
	DefaultProperties
	FraID uint64
	// Means for Giving Warning
	MeansForGivingWarning_2_0_Comment string `gorm:"type:text"`
	// MeansForGivingWarning_2_0_Recommendation    string `gorm:"type:text"`
	FireAlarmDetectionSystem_2_0_Comment string `gorm:"type:text"`
	// FireAlarmDetectionSystem_2_0_Recommendation string `gorm:"type:text"`

	BuildingHaveMeansGivingWarning_2_1           string `gorm:"size:24"`
	BuildingHaveMeansGivingWarning_2_1_Risklevel string `gorm:"size:24"`
	MeansGivingWarningAppropriate_2_2            string `gorm:"size:24"`
	MeansGivingWarningAppropriate_2_2_Risklevel  string `gorm:"size:24"`
	AlarmAudibleThroughout_2_3                   string `gorm:"size:24"`
	AlarmAudibleThroughout_2_3_Risklevel         string `gorm:"size:24"`
	CallPointsSatisfactory_2_4                   string `gorm:"size:24"`
	CallPointsSatisfactory_2_4_Risklevel         string `gorm:"size:24"`
	DetectionSufficient_2_5                      string `gorm:"size:24"`
	DetectionSufficient_2_5_Risklevel            string `gorm:"size:24"`
	DetectorsCallPointsObstructed_2_6            string `gorm:"size:24"`
	DetectorsCallPointsObstructed_2_6_Risklevel  string `gorm:"size:24"`
	OccupiersAware_2_7                           string `gorm:"size:24"`
	OccupiersAware_2_7_Risklevel                 string `gorm:"size:24"`
	OccupiersTrained_2_8                         string `gorm:"size:24"`
	OccupiersTrained_2_8_Risklevel               string `gorm:"size:24"`
	FireAlarmSystemTested_2_9                    string `gorm:"size:24"`
	FireAlarmSystemTested_2_9_Risklevel          string `gorm:"size:24"`
	FireAlarmSystemServiced_2_10                 string `gorm:"size:24"`
	FireAlarmSystemServiced_2_10_Risklevel       string `gorm:"size:24"`
	BatteryBackUp_2_11                           string `gorm:"size:24"`
	BatteryBackUp_2_11_Risklevel                 string `gorm:"size:24"`
}
type FraEmergencyActionPlan struct {
	DefaultProperties
	FraID uint64
	// Emergency Action Plan
	EmergencyActionPlan_3_0_Comment string `gorm:"type:text"`
	// EmergencyActionPlan_3_0_Recommendation string `gorm:"type:text"`

	WrittenFireEmergencyActionPlan_3_1           string `gorm:"size:24"`
	WrittenFireEmergencyActionPlan_3_1_Risklevel string `gorm:"size:24"`
	PointSpecifiedCommunalBoard_3_2              string `gorm:"size:24"`
	PointSpecifiedCommunalBoard_3_2_Risklevel    string `gorm:"size:24"`
	ActionDiscoveringFire_3_3                    string `gorm:"size:24"`
	ActionDiscoveringFire_3_3_Risklevel          string `gorm:"size:24"`
	HowRaiseAlarm_3_4                            string `gorm:"size:24"`
	HowRaiseAlarm_3_4_Risklevel                  string `gorm:"size:24"`
	ActionHearingFireAlarm_3_5                   string `gorm:"size:24"`
	ActionHearingFireAlarm_3_5_Risklevel         string `gorm:"size:24"`
	ProcedureAlertingMembers_3_6                 string `gorm:"size:24"`
	ProcedureAlertingMembers_3_6_Risklevel       string `gorm:"size:24"`
	EvacuationProcedur_3_7                       string `gorm:"size:24"`
	EvacuationProcedur_3_7_Risklevel             string `gorm:"size:24"`
	LocationAppropriateUse_3_8                   string `gorm:"size:24"`
	LocationAppropriateUse_3_8_Risklevel         string `gorm:"size:24"`
	ImportanceClosingFireDoors_3_9               string `gorm:"size:24"`
	ImportanceClosingFireDoors_3_9_Risklevel     string `gorm:"size:24"`
	IsolationMachinery_3_10                      string `gorm:"size:24"`
	IsolationMachinery_3_10_Risklevel            string `gorm:"size:24"`
	ReasonNotUsingLifts_3_11                     string `gorm:"size:24"`
	ReasonNotUsingLifts_3_11_Risklevel           string `gorm:"size:24"`
	InformationSpecificHazards_3_12              string `gorm:"size:24"`
	InformationSpecificHazards_3_12_Risklevel    string `gorm:"size:24"`
}
type FraPortableFireFightingEquipment struct {
	DefaultProperties
	FraID uint64
	//Portable fire-fighting equipment
	PortableFireFightingEquipment_4_0_Comment string `gorm:"type:text"`
	// PortableFireFightingEquipment_4_0_Recommendation string `gorm:"type:text"`

	SufficientExtinguishers_4_1                    string `gorm:"size:24"`
	SufficientExtinguishers_4_1_Risklevel          string `gorm:"size:24"`
	ExtinguishersCorrectlyPositioned_4_2           string `gorm:"size:24"`
	ExtinguishersCorrectlyPositioned_4_2_Risklevel string `gorm:"size:24"`
	ExtinguishersAppropriate_4_3                   string `gorm:"size:24"`
	ExtinguishersAppropriate_4_3_Risklevel         string `gorm:"size:24"`
	FireBlankets_4_4                               string `gorm:"size:24"`
	FireBlankets_4_4_Risklevel                     string `gorm:"size:24"`
	ExtinguishersObstructed_4_5                    string `gorm:"size:24"`
	ExtinguishersObstructed_4_5_Risklevel          string `gorm:"size:24"`
	ExtinguisherSigns_4_6                          string `gorm:"size:24"`
	ExtinguisherSigns_4_6_Risklevel                string `gorm:"size:24"`
	ExtinguishersServiced_4_7                      string `gorm:"size:24"`
	ExtinguishersServiced_4_7_Risklevel            string `gorm:"size:24"`
}
type FraFixedInstallation struct {
	DefaultProperties
	FraID uint64
	//Fixed Installations
	FixedInstallations_5_0_Comment string `gorm:"type:text"`
	// FixedInstallations_5_0_Recommendation string `gorm:"type:text"`

	SprinklerSystem_5_1                    string `gorm:"size:24"`
	SprinklerSystem_5_1_Risklevel          string `gorm:"size:24"`
	DryWetRiser_5_2                        string `gorm:"size:24"`
	DryWetRiser_5_2_Risklevel              string `gorm:"size:24"`
	GasFloodingSystem_5_3                  string `gorm:"size:24"`
	GasFloodingSystem_5_3_Risklevel        string `gorm:"size:24"`
	SmokeControlSystem_5_4                 string `gorm:"size:24"`
	SmokeControlSystem_5_4_Risklevel       string `gorm:"size:24"`
	OtherFixedInstallation_5_5             string `gorm:"size:24"`
	OtherFixedInstallation_5_5_Risklevel   string `gorm:"size:24"`
	HoseReelsProvided_5_6                  string `gorm:"size:24"`
	HoseReelsProvided_5_6_Risklevel        string `gorm:"size:24"`
	FixedInstallationsTested_5_7           string `gorm:"size:24"`
	FixedInstallationsTested_5_7_Risklevel string `gorm:"size:24"`
	ProtectionSystemsSignage_5_8           string `gorm:"size:24"`
	ProtectionSystemsSignage_5_8_Risklevel string `gorm:"size:24"`
	ProtectionSystemsAlarms_5_9            string `gorm:"size:24"`
	ProtectionSystemsAlarms_5_9_Risklevel  string `gorm:"size:24"`
}
type FraArson struct {
	DefaultProperties
	FraID uint64
	//Arson
	Arson_6_0_Comment        string `gorm:"type:text"`
	Arson_6_0_Recommendation string `gorm:"type:text"`

	SecureDuringDarkness_6_1                string `gorm:"size:24"`
	SecureDuringDarkness_6_1_Risklevel      string `gorm:"size:24"`
	StandardExternalLighting_6_2            string `gorm:"size:24"`
	StandardExternalLighting_6_2_Risklevel  string `gorm:"size:24"`
	ExternalRubbish_6_3                     string `gorm:"size:24"`
	ExternalRubbish_6_3_Risklevel           string `gorm:"size:24"`
	WheelieBinsSecurelyPlaced_6_4           string `gorm:"size:24"`
	WheelieBinsSecurelyPlaced_6_4_Risklevel string `gorm:"size:24"`
	CCTVAvailable_6_5                       string `gorm:"size:24"`
	CCTVAvailable_6_5_Risklevel             string `gorm:"size:24"`
	VideoDoorbellAvailable_6_6              string `gorm:"size:24"`
	VideoDoorbellAvailable_6_6_Risklevel    string `gorm:"size:24"`
}
type FraMaintenanceRecordKeeping struct {
	DefaultProperties
	FraID uint64
	//7.0 Maintenance and Record Keeping
	MaintenanceAndRecordKeeping_7_0_Comment string `gorm:"type:text"`
	// MaintenanceAndRecordKeeping_7_0_Recommendation string `gorm:"type:text"`

	SufficientRiskAssessment_7_1                  string `gorm:"size:24"`
	SufficientRiskAssessment_7_1_Risklevel        string `gorm:"size:24"`
	RiskAssessmentReviewed_7_2                    string `gorm:"size:24"`
	RiskAssessmentReviewed_7_2_Risklevel          string `gorm:"size:24"`
	SufficientEmergencyActionPlan_7_3             string `gorm:"size:24"`
	SufficientEmergencyActionPlan_7_3_Risklevel   string `gorm:"size:24"`
	LogBookKept_7_4                               string `gorm:"size:24"`
	LogBookKept_7_4_Risklevel                     string `gorm:"size:24"`
	FireAlarmTestingRecord_7_5                    string `gorm:"size:24"`
	FireAlarmTestingRecord_7_5_Risklevel          string `gorm:"size:24"`
	EmergencyLightTestingRecord_7_6               string `gorm:"size:24"`
	EmergencyLightTestingRecord_7_6_Risklevel     string `gorm:"size:24"`
	MeansEscapeRouteRecord_7_7                    string `gorm:"size:24"`
	MeansEscapeRouteRecord_7_7_Risklevel          string `gorm:"size:24"`
	FireExtinguisherEquipmentRecord_7_8           string `gorm:"size:24"`
	FireExtinguisherEquipmentRecord_7_8_Risklevel string `gorm:"size:24"`
	FixedInstallationTesting_7_9                  string `gorm:"size:24"`
	FixedInstallationTesting_7_9_Risklevel        string `gorm:"size:24"`
	GeneralFireTraining_7_10                      string `gorm:"size:24"`
	GeneralFireTraining_7_10_Risklevel            string `gorm:"size:24"`
	FireDrillsRecord_7_11                         string `gorm:"size:24"`
	FireDrillsRecord_7_11_Risklevel               string `gorm:"size:24"`
}
type FraHazardIdentificationDoor struct {
	// Hazard identification
	DefaultProperties
	FraID uint64

	// Doors
	SufficientExits_1_1                   string `gorm:"size:24"`
	SufficientExits_1_1_Risklevel         string `gorm:"size:24"`
	ImmediatelyOpened_1_2                 string `gorm:"size:24"`
	ImmediatelyOpened_1_2_Risklevel       string `gorm:"size:24"`
	AppropriateDevices_1_3                string `gorm:"size:24"`
	AppropriateDevices_1_3_Risklevel      string `gorm:"size:24"`
	InternalDoorsWedged_1_4               string `gorm:"size:24"`
	InternalDoorsWedged_1_4_Risklevel     string `gorm:"size:24"`
	AutomaticDoorFastenings_1_5           string `gorm:"size:24"`
	AutomaticDoorFastenings_1_5_Risklevel string `gorm:"size:24"`

	HazardIdentification_1_1_0_Comment string `gorm:"type:text"`
	// HazardIdentification_1_1_0_Recommendation string `gorm:"type:text"`
	Doors_1_0_Comment string `gorm:"type:text"`
	// Doors_1_0_Recommendation                  string `gorm:"type:text"`
}

type FraHazardIdentificationSignage struct {
	// Hazard identification
	DefaultProperties
	FraID uint64
	// Signage
	SufficientExitSignage_2_1            string `gorm:"size:24"`
	SufficientExitSignage_2_1_Risklevel  string `gorm:"size:24"`
	EscapeRoutesIdentified_2_2           string `gorm:"size:24"`
	EscapeRoutesIdentified_2_2_Risklevel string `gorm:"size:24"`
	SignageComply_2_3                    string `gorm:"size:24"`
	SignageComply_2_3_Risklevel          string `gorm:"size:24"`
	AppropriateLabel_2_4                 string `gorm:"size:24"`
	AppropriateLabel_2_4_Risklevel       string `gorm:"size:24"`
	ExternalSignage_2_5                  string `gorm:"size:24"`
	ExternalSignage_2_5_Risklevel        string `gorm:"size:24"`
	AdequateNotices_2_6                  string `gorm:"size:24"`
	AdequateNotices_2_6_Risklevel        string `gorm:"size:24"`
	IlluminatedExitSigns_2_7             string `gorm:"size:24"`
	IlluminatedExitSigns_2_7_Risklevel   string `gorm:"size:24"`

	Signage_2_0_Comment string `gorm:"type:text"`
	// Signage_2_0_Recommendation string `gorm:"type:text"`
}

type FraHazardIdentificationLighting struct {
	// Hazard identification
	DefaultProperties
	FraID uint64
	// Lighting
	Lighting_3_0_Comment string `gorm:"type:text"`
	// Lighting_3_0_Recommendation string `gorm:"type:text"`

	EmergencyLighting_3_1                           string `gorm:"size:24"`
	EmergencyLighting_3_1_Risklevel                 string `gorm:"size:24"`
	EscapeRoutesAdequatelyIlluminated_3_2           string `gorm:"size:24"`
	EscapeRoutesAdequatelyIlluminated_3_2_Risklevel string `gorm:"size:24"`
	ExistingArtificialLighting_3_3                  string `gorm:"size:24"`
	ExistingArtificialLighting_3_3_Risklevel        string `gorm:"size:24"`
	EmergencyLightingCharging_3_4                   string `gorm:"size:24"`
	EmergencyLightingCharging_3_4_Risklevel         string `gorm:"size:24"`
	LightingMaintained_3_5                          string `gorm:"size:24"`
	LightingMaintained_3_5_Risklevel                string `gorm:"size:24"`
}

type FraHazardIdentificationElectrical struct {
	// Hazard identification
	DefaultProperties
	FraID uint64
	// Electrical
	Electrical_4_0_Comment string `gorm:"type:text"`
	// Electrical_4_0_Recommendation string `gorm:"type:text"`

	ReasonableMeasures_4_1               string `gorm:"size:24"`
	ReasonableMeasures_4_1_Risklevel     string `gorm:"size:24"`
	ElectricalInstallation_4_2           string `gorm:"size:24"`
	ElectricalInstallation_4_2_Risklevel string `gorm:"size:24"`
	ElectricalEquipment_4_3              string `gorm:"size:24"`
	ElectricalEquipment_4_3_Risklevel    string `gorm:"size:24"`

	// Smoking
	Smoking_5_0_Comment string `gorm:"type:text"`
	// Smoking_5_0_Recommendation string `gorm:"type:text"`

	PeopleSmoke_5_1                    string `gorm:"size:24"`
	PeopleSmoke_5_1_Risklevel          string `gorm:"size:24"`
	SuitableArrangements_5_2           string `gorm:"size:24"`
	SuitableArrangements_5_2_Risklevel string `gorm:"size:24"`
	SignageAvailable_5_3               string `gorm:"size:24"`
	SignageAvailable_5_3_Risklevel     string `gorm:"size:24"`

	// Portable heating & heating installations
	PortableHeating_6_0_Comment string `gorm:"type:text"`
	// PortableHeating_6_0_Recommendation string `gorm:"type:text"`

	PortableHeatersAvoided_6_1           string `gorm:"size:24"`
	PortableHeatersAvoided_6_1_Risklevel string `gorm:"size:24"`
	SuitableMeasures_6_2                 string `gorm:"size:24"`
	SuitableMeasures_6_2_Risklevel       string `gorm:"size:24"`
	RegularMaintenance_6_3               string `gorm:"size:24"`
	RegularMaintenance_6_3_Risklevel     string `gorm:"size:24"`
}

type FraHazardIdentificationCooking struct {
	// Hazard identification
	DefaultProperties
	FraID uint64
	// Cooking 7.0 `gorm:"type:text"`
	Cooking_7_0_Comment string `gorm:"type:text"`
	// Cooking_7_0_Recommendation string `gorm:"type:text"`

	ReasonableMeasuresPreventFires_7_1           string `gorm:"size:24"`
	ReasonableMeasuresPreventFires_7_1_Risklevel string `gorm:"size:24"`
	FiltersCleanedChangedRegularly_7_2           string `gorm:"size:24"`
	FiltersCleanedChangedRegularly_7_2_Risklevel string `gorm:"size:24"`
	AppliancesRegularlyMaintained_7_3            string `gorm:"size:24"`
	AppliancesRegularlyMaintained_7_3_Risklevel  string `gorm:"size:24"`
	SuitableFireBlanketAvailable_7_4             string `gorm:"size:24"`
	SuitableFireBlanketAvailable_7_4_Risklevel   string `gorm:"size:24"`
	GasSafetyRecordAvailable_7_5                 string `gorm:"size:24"`
	GasSafetyRecordAvailable_7_5_Risklevel       string `gorm:"size:24"`

	// Furniture & Furnishings (Fire safety Regulations 1993)
	FurnitureFurnishings_8_0_Comment string `gorm:"type:text"`
	// FurnitureFurnishings_8_0_Recommendation string `gorm:"type:text"`

	MeetSpecifiedIgnitionRequirements_8_1                string `gorm:"size:24"`
	MeetSpecifiedIgnitionRequirements_8_1_Risklevel      string `gorm:"size:24"`
	UpholsteryCompositesCigaretteResistant_8_2           string `gorm:"size:24"`
	UpholsteryCompositesCigaretteResistant_8_2_Risklevel string `gorm:"size:24"`
	CoversResistant_8_3                                  string `gorm:"size:24"`
	CoversResistant_8_3_Risklevel                        string `gorm:"size:24"`
	PermanentLabelFurniture_8_4                          string `gorm:"size:24"`
	PermanentLabelFurniture_8_4_Risklevel                string `gorm:"size:24"`
	DisplayLabelFurniture_8_5                            string `gorm:"size:24"`
	DisplayLabelFurniture_8_5_Risklevel                  string `gorm:"size:24"`
	FirstSupplierFurnitureCompliance_8_6                 string `gorm:"size:24"`
	FirstSupplierFurnitureCompliance_8_6_Risklevel       string `gorm:"size:24"`

	// Housekeeping
	Housekeeping_9_0_Comment string `gorm:"type:text"`
	// Housekeeping_9_0_Recommendation string `gorm:"type:text"`

	StandardHousekeepingAdequate_9_1                    string `gorm:"size:24"`
	StandardHousekeepingAdequate_9_1_Risklevel          string `gorm:"size:24"`
	CombustibleMaterialsIgnitionSources_9_2             string `gorm:"size:24"`
	CombustibleMaterialsIgnitionSources_9_2_Risklevel   string `gorm:"size:24"`
	HazardousMaterialsStoredAppropriately_9_3           string `gorm:"size:24"`
	HazardousMaterialsStoredAppropriately_9_3_Risklevel string `gorm:"size:24"`
}
type FraGeneral struct {
	DefaultProperties
	FraID uint64
	// General
	General_10_0_Comment string `gorm:"type:text"`
	// General_10_0_Recommendation string `gorm:"type:text"`

	AlternativeExit_10_1                           string `gorm:"size:24"`
	AlternativeExit_10_1_Risklevel                 string `gorm:"size:24"`
	DeadEndSituations_10_2                         string `gorm:"size:24"`
	DeadEndSituations_10_2_Risklevel               string `gorm:"size:24"`
	EscapeRoutesFree_10_3                          string `gorm:"size:24"`
	EscapeRoutesFree_10_3_Risklevel                string `gorm:"size:24"`
	CombustibleMaterialsEscapeRoute_10_4           string `gorm:"size:24"`
	CombustibleMaterialsEscapeRoute_10_4_Risklevel string `gorm:"size:24"`
	SuitableArrangementsInnerRoom_10_5             string `gorm:"size:24"`
	SuitableArrangementsInnerRoom_10_5_Risklevel   string `gorm:"size:24"`
	EscapeRoutesAdequatelyProtected_10_6           string `gorm:"size:24"`
	EscapeRoutesAdequatelyProtected_10_6_Risklevel string `gorm:"size:24"`

	// Others
	Others_11_0_Comment string `gorm:"type:text"`
	// Others_11_0_Recommendation string `gorm:"type:text"`
	Others_11_1           string `gorm:"type:text"`
	Others_11_1_Risklevel string `gorm:"size:24"`
}
type FraRecommendation struct {
	DefaultProperties
	FraID uint64
	// Recommendation
	Action       string `gorm:"type:text"`
	Priority     string `gorm:"size:24"`
	Location     string `gorm:"size:24"`
	Ref          string `gorm:"size:250"`
	RiskCategory string `gorm:"size:150"`
	Compliance   string `gorm:"size:250"`
}

type Eicr struct {
	DefaultProperties
	EngineerID       uint64
	Engineer         Engineer
	EicrObservations []EicrObservation
	EicrPage3        EicrPage3

	//eicr page 1
	OrderID uint64
	Name    string `gorm:"size:255"`
	Address string `gorm:"size:255"`

	LandlordName     string    `gorm:"size:255"`
	LandlordAddress  string    `gorm:"size:255"`
	ContactNumber    string    `gorm:"size:255"`
	Email            string    `gorm:"size:255"`
	AssessmentDate   time.Time `form:"AssessmentDate"  time_format:"2006-01-02"`
	AssessmentReason string    `gorm:"size:512"`
	ReviewDate       time.Time `form:"ReviewDate"  time_format:"2006-01-02"`

	EstimatedWiringAge     string `gorm:"size:50"`
	EvidenceAlterations    string `gorm:"size:25"`
	AlterationsAge         string `gorm:"size:25"`
	InstallationRecords    string `gorm:"size:25"`
	LastInspectionDate     string `gorm:"size:50"`
	Extent                 string `gorm:"type:text"`
	Limitations            string `gorm:"type:text"`
	AgreedWith             string `gorm:"size:255"`
	OperationalLimitations string `gorm:"type:text"`
	AssessmentSummary      string `gorm:"size:50"`
	SatisfactoryYear       string `gorm:"size:50"`

	HasObservations string `gorm:"size:10"`
	C1Items         string `gorm:"size:50"`
	C2Items         string `gorm:"size:50"`
	C3Items         string `gorm:"size:50"`
	FiItems         string `gorm:"size:50"`
}

type EicrObservation struct {
	DefaultProperties
	EicrID      uint64
	ItemNumber  int
	Observation string `gorm:"size:512"`
	Code        string `gorm:"size:25"`
}

type EicrPage3 struct {
	DefaultProperties
	EicrID                     uint64
	GeneralCondition           string `gorm:"type:text"`
	EarthingArrangements       string `gorm:"size:15"` //tns,tncs,tt system
	LiveConductors             string `gorm:"size:50"` // 1-phase (2-wire) 2-phase (3 wire ), Other: etc
	ConfirmationSupplyPolarity string `gorm:"size:10"`
	// Nature of Supply Parameters
	NominalVoltage                  string `gorm:"size:20"`
	NominalFrequency                string `gorm:"size:20"`
	ProspectiveFaultCurrent         string `gorm:"size:20"`
	ExternalEarthFaultLoopImpedance string `gorm:"size:20"`
	// Supply Protective Device
	BsEn         string `gorm:"size:20"`
	BsEnType     string `gorm:"size:20"`
	RatedCurrent string `gorm:"size:20"`
	// 11 PARTICULARS OF INSTALLATION REFERRED TO IN THE REPORT
	// Means of Earthing
	DistributorFacility        string `gorm:"size:10"`
	InstallationEarthElectrode string `gorm:"size:10"`
	//Details of Installation Earth Electrode (where applicable)
	EarthType         string `gorm:"size:20"`
	EarthLocation     string `gorm:"size:50"`
	ResistanceEarth   string `gorm:"size:20"`
	MethodMeasurement string `gorm:"size:50"`

	//Main Switch / Switch-Fuse / Circuit-Breaker / RCD
	FuseboxLocation string `gorm:"size:50"`
	FuseboxBsEn     string `gorm:"size:50"`
	NumberOfPoles   string `gorm:"size:50"`
	CurrentRating   string `gorm:"size:50"`
	FuseRating      string `gorm:"size:50"`
	VoltageRating   string `gorm:"size:50"`
	// If RCD main switch:
	RcdType                       string `gorm:"size:50"`
	RatedResidualOperatingCurrent string `gorm:"size:50"`
	RatedTimeDelay                string `gorm:"size:50"`
	MeasuredOperatingTime         string `gorm:"size:50"`
	//Earthing and Protective Bonding Conductors
	EarthingConductorMaterial           string `gorm:"size:50"`
	EarthingConductorCsa                string `gorm:"size:50"`
	EarthingConductorConnectionVerified string `gorm:"size:50"`
	MainConductorMaterial               string `gorm:"size:50"`
	MainConductorCsa                    string `gorm:"size:50"`
	MainConductorConnectionVerified     string `gorm:"size:50"`

	//Bonding of extraneous-conductive parts
	ToWaterInstallationPipes string `gorm:"size:50"`
	ToGasInstallationPipes   string `gorm:"size:50"`
	ToOilInstallationPipes   string `gorm:"size:50"`
	TolightningProtection    string `gorm:"size:50"`
	ToOtherService           string `gorm:"size:50"`
}

// 12 INSPECTION SCHEDULE
type InspectionSchedule struct {
	DefaultProperties
	EicrID uint64
	//page 1
	ServiceCable_1_1_1                    string `gorm:"size:10"`
	ServiceHead_1_1_2                     string `gorm:"size:10"`
	EarthingArrangement_1_1_3             string `gorm:"size:10"`
	MeterTails_1_1_4                      string `gorm:"size:10"`
	MeteringEquipment1_1_5                string `gorm:"size:10"`
	Isolator_1_1_6                        string `gorm:"size:10"`
	ConsumerIsolator_1_2                  string `gorm:"size:10"`
	ConsumerMeterTails_1_3                string `gorm:"size:10"`
	OtherSources_2_0                      string `gorm:"size:10"`
	PresenceAndCondition_3_1              string `gorm:"size:10"`
	PresenceAndCondition_3_2              string `gorm:"size:10"`
	ProvisionOfEarthing_3_3               string `gorm:"size:10"`
	ConfirmationOfEarthing_3_4            string `gorm:"size:10"`
	AccessibilityAndConditionEarthing_3_5 string `gorm:"size:10"`
	ConfirmationOfMain_3_6                string `gorm:"size:10"`
	ConditionAndAccessibilityMain_3_7     string `gorm:"size:10"`
	AccessibilityAndConditionOther_3_3    string `gorm:"size:10"`
	//4.0 CONSUMER UNIT(S) / DISTRIBUTION BOARD(S)
	AdequacyOfWorking_4_1            string `gorm:"size:10"`
	SecurityOfFixing_4_2             string `gorm:"size:10"`
	IpRatingEtc_4_3                  string `gorm:"size:10"`
	FireRatingetc_4_4                string `gorm:"size:10"`
	EnclosureNotDamaged_4_5          string `gorm:"size:10"`
	PresenceOfMain_4_6               string `gorm:"size:10"`
	OperationOfMain_4_7              string `gorm:"size:10"`
	ManualOperation_4_8              string `gorm:"size:10"`
	CorrectIdentification_4_9        string `gorm:"size:10"`
	PresenceOfRcd_4_10               string `gorm:"size:10"`
	PresenceOfAlternativeSupply_4_11 string `gorm:"size:10"`
	PresenceOfOther_4_12             string `gorm:"size:10"`

	// Compatibility of protective devices
	CompatibilityOfProtective_4_13             string `gorm:"size:10"`
	SinglePoleSwitching_4_14                   string `gorm:"size:10"`
	ProtectionAgainstMechanical_4_15           string `gorm:"size:10"`
	ProtectionAgainstElectromagnetic_4_16      string `gorm:"size:10"`
	RcdsProvidedForFaultProtection_4_17        string `gorm:"size:10"`
	RcdsProvidedForAdditionalProtection_4_18   string `gorm:"size:10"`
	ConfirmationOfIndicationOfSpd_4_19         string `gorm:"size:10"`
	ConfirmationOfAllConductorConnections_4_20 string `gorm:"size:10"`
	AdequateSwitchedAlternative_4_21           string `gorm:"size:10"`
	AdequateGeneratingSetParallel_4_22         string `gorm:"size:10"`

	// Section 5
	IdentificationOfConductors_5_1          string `gorm:"size:10"`
	CablesCorrectlySupported_5_2            string `gorm:"size:10"`
	ConditionOfInsulation_5_3               string `gorm:"size:10"`
	NonSheathedCablesProtected_5_4          string `gorm:"size:10"`
	IntegrityOfConduitAndTrunking_5_4_1     string `gorm:"size:10"`
	AdequacyOfCablesForCurrent_5_5          string `gorm:"size:10"`
	CoordinationBetweenConductors_5_6       string `gorm:"size:10"`
	AdequacyOfProtectiveDevices_5_7         string `gorm:"size:10"`
	PresenceAndAdequacyOfCpc_5_8            string `gorm:"size:10"`
	AppropriateWiringSystems_5_9            string `gorm:"size:10"`
	ConcealedCablesInPrescribedZones_5_10   string `gorm:"size:10"`
	ConcealedCablesProtected_5_11           string `gorm:"size:10"`
	RcdForSockets_32AOrLess_5_12_1          string `gorm:"size:10"`
	RcdForMobileEquipment_32A_5_12_2        string `gorm:"size:10"`
	RcdForCablesLessThan50mmDepth_5_12_3    string `gorm:"size:10"`
	RcdForCablesInWalls_5_12_4              string `gorm:"size:10"`
	RcdForFinalCircuits_5_12_5              string `gorm:"size:10"`
	ProvisionOfFireBarriers_5_13            string `gorm:"size:10"`
	BandIICablesSeparated_5_14              string `gorm:"size:10"`
	CablesSeparatedFromComms_5_15           string `gorm:"size:10"`
	CablesSeparatedFromNonElectrical_5_16   string `gorm:"size:10"`
	ConnectionsSoundlyMade_5_17_1           string `gorm:"size:10"`
	NoBasicInsulationVisible_5_17_2         string `gorm:"size:10"`
	LiveConductorsAdequatelyEnclosed_5_17_3 string `gorm:"size:10"`
	AdequateEntryConnection_5_17_4          string `gorm:"size:10"`
	ConditionOfAccessories_5_18             string `gorm:"size:10"`
	SuitabilityOfAccessories_5_19           string `gorm:"size:10"`
	AdequacyOfWorkingSpace_5_20             string `gorm:"size:10"`
	SinglePoleSwitchingInLine_5_21          string `gorm:"size:10"`

	// Section 6
	AdditionalRcdProtectionForLVCircuits_6_1 string `gorm:"size:10"`
	SelvOrPelvRequirementsMet_6_2            string `gorm:"size:10"`
	ShaverSupplyUnitsCompliant_6_3           string `gorm:"size:10"`
	PresenceOfSupplementaryBonding_6_4       string `gorm:"size:10"`
	LVSocketOutletsDistance_6_5              string `gorm:"size:10"`
	SuitabilityForExternalInfluences_6_6     string `gorm:"size:10"`
	SuitabilityOfAccessoriesForZone_6_7      string `gorm:"size:10"`
	SuitabilityOfEquipmentForLocation_6_8    string `gorm:"size:10"`
}
type Engineer struct {
	DefaultProperties
	Title            string `gorm:"size:100"`
	Name             string `gorm:"size:100"`
	Registration     string `gorm:"size:100"`
	GasID            string `gorm:"size:100"`
	GasSafeNo        string `gorm:"size:100"`
	EicrTradingTitle string `gorm:"size:250"`
	GasTradingTitle  string `gorm:"size:250"`
	FraTradingTitle  string `gorm:"size:250"`
	Address          string `gorm:"size:250"`
	PostCode         string `gorm:"size:15"`
	Phone            string `gorm:"size:25"`
	Position         string `gorm:"size:50"`
	Signature        string `gorm:"type:text"`

	//Test Instrument SRL number
	MultiFunctional          string `gorm:"size:25"`
	InsulationResistance     string `gorm:"size:25"`
	Continuity               string `gorm:"size:25"`
	EarthElectrodeResistance string `gorm:"size:25"`
	EarthFaultLoopImpedance  string `gorm:"size:25"`
	RCD                      string `gorm:"size:25"`
}
