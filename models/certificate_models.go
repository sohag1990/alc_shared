package models

import (
	"time"

	"github.com/sohag1990/alc_shared/wp_models"
)

type FollowUp struct {
	DefaultProperties
	OrderID         uint64
	FollowUpOption  string `gorm:"size:100"`
	Services        string `gorm:"size:250"`
	FollowUpSubject string `gorm:"size:150"`
	FollowUpContent string `gorm:"type:text"`
	FeedbackContent string `gorm:"type:text"`
	Status          bool
	LineItems       []wp_models.LineItem
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
	CertificateIssuer string
	OrderID           uint64
	// Order                           Order
	Name                            string
	Address                         string
	Description                     string
	NumberOfFloors                  int
	ConstructionDate                time.Time `form:"ConstructionDate"  time_format:"2006-01-02"`
	BriefDescriptionOfAccommodation string

	LandlordName      string
	LandlordAddress   string
	ResponsiblePerson string
	ContactNumber     string
	Email             string

	SleepingOccupants string `gorm:"size:100"`
	DisableOccupants  string `gorm:"size:100"`
	ChildrenOccupants string `gorm:"size:100"`
	ElderlyOccupants  string `gorm:"size:100"`

	AssessmentDate    time.Time `form:"AssessmentDate"  time_format:"2006-01-02"`
	AssessmentType    string    `gorm:"size:100"`
	AssessedBy        string    `gorm:"size:100"`
	Registration      string    `gorm:"size:100"`
	AssessmentPurpose string    `gorm:"size:250"`
	ReviewDate        time.Time `form:"ReviewDate"  time_format:"2006-01-02"`

	PotentialRiskScore string `gorm:"size:24"`

	// Hazard identification
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

	HazardIdentification_1_1_0_Comment        string
	HazardIdentification_1_1_0_Recommendation string
	Doors_1_0_Comment                         string
	Doors_1_0_Recommendation                  string

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

	Signage_2_0_Comment        string
	Signage_2_0_Recommendation string

	// Lighting
	Lighting_3_0_Comment        string
	Lighting_3_0_Recommendation string

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

	// Electrical
	Electrical_4_0_Comment        string
	Electrical_4_0_Recommendation string

	ReasonableMeasures_4_1               string `gorm:"size:24"`
	ReasonableMeasures_4_1_Risklevel     string `gorm:"size:24"`
	ElectricalInstallation_4_2           string `gorm:"size:24"`
	ElectricalInstallation_4_2_Risklevel string `gorm:"size:24"`
	ElectricalEquipment_4_3              string `gorm:"size:24"`
	ElectricalEquipment_4_3_Risklevel    string `gorm:"size:24"`

	// Smoking
	Smoking_5_0_Comment        string
	Smoking_5_0_Recommendation string

	PeopleSmoke_5_1                    string `gorm:"size:24"`
	PeopleSmoke_5_1_Risklevel          string `gorm:"size:24"`
	SuitableArrangements_5_2           string `gorm:"size:24"`
	SuitableArrangements_5_2_Risklevel string `gorm:"size:24"`
	SignageAvailable_5_3               string `gorm:"size:24"`
	SignageAvailable_5_3_Risklevel     string `gorm:"size:24"`

	// Portable heating & heating installations
	PortableHeating_6_0_Comment        string
	PortableHeating_6_0_Recommendation string

	PortableHeatersAvoided_6_1           string `gorm:"size:24"`
	PortableHeatersAvoided_6_1_Risklevel string `gorm:"size:24"`
	SuitableMeasures_6_2                 string `gorm:"size:24"`
	SuitableMeasures_6_2_Risklevel       string `gorm:"size:24"`
	RegularMaintenance_6_3               string `gorm:"size:24"`
	RegularMaintenance_6_3_Risklevel     string `gorm:"size:24"`

	// Cooking 7.0
	Cooking_7_0_Comment        string
	Cooking_6_0_Recommendation string

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
	FurnitureFurnishings_8_0_Comment        string
	FurnitureFurnishings_8_0_Recommendation string

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
	Housekeeping_9_0_Comment        string
	Housekeeping_9_0_Recommendation string

	StandardHousekeepingAdequate_9_1                    string `gorm:"size:24"`
	StandardHousekeepingAdequate_9_1_Risklevel          string `gorm:"size:24"`
	CombustibleMaterialsIgnitionSources_9_2             string `gorm:"size:24"`
	CombustibleMaterialsIgnitionSources_9_2_Risklevel   string `gorm:"size:24"`
	HazardousMaterialsStoredAppropriately_9_3           string `gorm:"size:24"`
	HazardousMaterialsStoredAppropriately_9_3_Risklevel string `gorm:"size:24"`

	// General
	General_10_0_Comment        string
	General_10_0_Recommendation string

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
	Others_11_0_Comment        string
	Others_11_0_Recommendation string
	Others_11_1                string `gorm:"size:250"`
	Others_11_1_Risklevel      string `gorm:"size:24"`

	// Means for Giving Warning
	MeansForGivingWarning_2_0_Comment           string
	MeansForGivingWarning_2_0_Recommendation    string
	FireAlarmDetectionSystem_2_0_Comment        string
	FireAlarmDetectionSystem_2_0_Recommendation string

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

	// Emergency Action Plan
	EmergencyActionPlan_3_0_Comment        string
	EmergencyActionPlan_3_0_Recommendation string

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

	//Portable fire-fighting equipment
	PortableFireFightingEquipment_4_0_Comment        string
	PortableFireFightingEquipment_4_0_Recommendation string

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

	//Fixed Installations
	FixedInstallations_5_0_Comment        string
	FixedInstallations_5_0_Recommendation string

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

	//Arson
	Arson_6_0_Comment        string
	Arson_6_0_Recommendation string

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

	//7.0 Maintenance and Record Keeping
	MaintenanceAndRecordKeeping_7_0_Comment        string
	MaintenanceAndRecordKeeping_7_0_Recommendation string

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
