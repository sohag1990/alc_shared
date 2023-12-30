package models

import "time"

type Fra struct {
	DefaultProperties
	OrderID                         uint64
	Order                           Order
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

	SleepingOccupants string
	DisableOccupants  string
	ChildrenOccupants string
	ElderlyOccupants  string

	AssessmentDate    time.Time `form:"AssessmentDate"  time_format:"2006-01-02"`
	AssessmentType    string
	AssessedBy        string
	Registration      string
	AssessmentPurpose string
	ReviewDate        time.Time `form:"ReviewDate"  time_format:"2006-01-02"`

	ProbabilityScore         string
	ProbabilityDescription   string
	PotentialRiskScore       string
	PotentialRiskDescription string

	// Hazard identification
	// Doors
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

	HazardIdentification_1_1_0_Comment        string
	HazardIdentification_1_1_0_Recommendation string
	Doors_1_0_Comment                         string
	Doors_1_0_Recommendation                  string

	// Signage
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
	IlluminatedExitSigns_2_7             string
	IlluminatedExitSigns_2_7_Risklevel   string

	Signage_2_0_Comment        string
	Signage_2_0_Recommendation string

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

	// Cooking 7.0
	Cooking_7_0_Comment        string
	Cooking_6_0_Recommendation string

	ReasonableMeasuresPreventFires_7_1           string
	ReasonableMeasuresPreventFires_7_1_Risklevel string
	FiltersCleanedChangedRegularly_7_2           string
	FiltersCleanedChangedRegularly_7_2_Risklevel string
	AppliancesRegularlyMaintained_7_3            string
	AppliancesRegularlyMaintained_7_3_Risklevel  string
	SuitableFireBlanketAvailable_7_4             string
	SuitableFireBlanketAvailable_7_4_Risklevel   string
	GasSafetyRecordAvailable_7_5                 string
	GasSafetyRecordAvailable_7_5_Risklevel       string

	// Furniture & Furnishings (Fire safety Regulations 1993)
	FurnitureFurnishings_8_0_Comment        string
	FurnitureFurnishings_8_0_Recommendation string

	MeetSpecifiedIgnitionRequirements_8_1                string
	MeetSpecifiedIgnitionRequirements_8_1_Risklevel      string
	UpholsteryCompositesCigaretteResistant_8_2           string
	UpholsteryCompositesCigaretteResistant_8_2_Risklevel string
	CoversResistant_8_3                                  string
	CoversResistant_8_3_Risklevel                        string
	PermanentLabelFurniture_8_4                          string
	PermanentLabelFurniture_8_4_Risklevel                string
	DisplayLabelFurniture_8_5                            string
	DisplayLabelFurniture_8_5_Risklevel                  string
	FirstSupplierFurnitureCompliance_8_6                 string
	FirstSupplierFurnitureCompliance_8_6_Risklevel       string

	// Housekeeping
	Housekeeping_9_0_Comment        string
	Housekeeping_9_0_Recommendation string

	StandardHousekeepingAdequate_9_1                    string
	StandardHousekeepingAdequate_9_1_Risklevel          string
	CombustibleMaterialsIgnitionSources_9_2             string
	CombustibleMaterialsIgnitionSources_9_2_Risklevel   string
	HazardousMaterialsStoredAppropriately_9_3           string
	HazardousMaterialsStoredAppropriately_9_3_Risklevel string

	// General
	General_10_0_Comment        string
	General_10_0_Recommendation string

	AlternativeExit_10_1                           string
	AlternativeExit_10_1_Risklevel                 string
	DeadEndSituations_10_2                         string
	DeadEndSituations_10_2_Risklevel               string
	EscapeRoutesFree_10_3                          string
	EscapeRoutesFree_10_3_Risklevel                string
	CombustibleMaterialsEscapeRoute_10_4           string
	CombustibleMaterialsEscapeRoute_10_4_Risklevel string
	SuitableArrangementsInnerRoom_10_5             string
	SuitableArrangementsInnerRoom_10_5_Risklevel   string
	EscapeRoutesAdequatelyProtected_10_6           string
	EscapeRoutesAdequatelyProtected_10_6_Risklevel string

	// Others
	Others_11_0_Comment        string
	Others_11_0_Recommendation string
	Others_11_1                string
	Others_11_1_Risklevel      string

	// Means for Giving Warning
	MeansForGivingWarning_2_0_Comment           string
	MeansForGivingWarning_2_0_Recommendation    string
	FireAlarmDetectionSystem_2_0_Comment        string
	FireAlarmDetectionSystem_2_0_Recommendation string

	BuildingHaveMeansGivingWarning_2_1           string
	BuildingHaveMeansGivingWarning_2_1_Risklevel string
	MeansGivingWarningAppropriate_2_2            string
	MeansGivingWarningAppropriate_2_2_Risklevel  string
	AlarmAudibleThroughout_2_3                   string
	AlarmAudibleThroughout_2_3_Risklevel         string
	CallPointsSatisfactory_2_4                   string
	CallPointsSatisfactory_2_4_Risklevel         string
	DetectionSufficient_2_5                      string
	DetectionSufficient_2_5_Risklevel            string
	DetectorsCallPointsObstructed_2_6            string
	DetectorsCallPointsObstructed_2_6_Risklevel  string
	OccupiersAware_2_7                           string
	OccupiersAware_2_7_Risklevel                 string
	OccupiersTrained_2_8                         string
	OccupiersTrained_2_8_Risklevel               string
	FireAlarmSystemTested_2_9                    string
	FireAlarmSystemTested_2_9_Risklevel          string
	FireAlarmSystemServiced_2_10                 string
	FireAlarmSystemServiced_2_10_Risklevel       string
	BatteryBackUp_2_11                           string
	BatteryBackUp_2_11_Risklevel                 string

	// Emergency Action Plan
	EmergencyActionPlan_3_0_Comment        string
	EmergencyActionPlan_3_0_Recommendation string

	WrittenFireEmergencyActionPlan_3_1           string
	WrittenFireEmergencyActionPlan_3_1_Risklevel string
	PointSpecifiedCommunalBoard_3_2              string
	PointSpecifiedCommunalBoard_3_2_Risklevel    string
	ActionDiscoveringFire_3_3                    string
	ActionDiscoveringFire_3_3_Risklevel          string
	HowRaiseAlarm_3_4                            string
	HowRaiseAlarm_3_4_Risklevel                  string
	ActionHearingFireAlarm_3_5                   string
	ActionHearingFireAlarm_3_5_Risklevel         string
	ProcedureAlertingMembers_3_6                 string
	ProcedureAlertingMembers_3_6_Risklevel       string
	EvacuationProcedur_3_7                       string
	EvacuationProcedur_3_7_Risklevel             string
	LocationAppropriateUse_3_8                   string
	LocationAppropriateUse_3_8_Risklevel         string
	ImportanceClosingFireDoors_3_9               string
	ImportanceClosingFireDoors_3_9_Risklevel     string
	IsolationMachinery_3_10                      string
	IsolationMachinery_3_10_Risklevel            string
	ReasonNotUsingLifts_3_11                     string
	ReasonNotUsingLifts_3_11_Risklevel           string
	InformationSpecificHazards_3_12              string
	InformationSpecificHazards_3_12_Risklevel    string

	//Portable fire-fighting equipment
	PortableFireFightingEquipment_4_0_Comment        string
	PortableFireFightingEquipment_4_0_Recommendation string

	SufficientExtinguishers_4_1                    string
	SufficientExtinguishers_4_1_Risklevel          string
	ExtinguishersCorrectlyPositioned_4_2           string
	ExtinguishersCorrectlyPositioned_4_2_Risklevel string
	ExtinguishersAppropriate_4_3                   string
	ExtinguishersAppropriate_4_3_Risklevel         string
	FireBlankets_4_4                               string
	FireBlankets_4_4_Risklevel                     string
	ExtinguishersObstructed_4_5                    string
	ExtinguishersObstructed_4_5_Risklevel          string
	ExtinguisherSigns_4_6                          string
	ExtinguisherSigns_4_6_Risklevel                string
	ExtinguishersServiced_4_7                      string
	ExtinguishersServiced_4_7_Risklevel            string

	//Fixed Installations
	FixedInstallations_5_0_Comment        string
	FixedInstallations_5_0_Recommendation string

	SprinklerSystem_5_1                    string
	SprinklerSystem_5_1_Risklevel          string
	DryWetRiser_5_2                        string
	DryWetRiser_5_2_Risklevel              string
	GasFloodingSystem_5_3                  string
	GasFloodingSystem_5_3_Risklevel        string
	SmokeControlSystem_5_4                 string
	SmokeControlSystem_5_4_Risklevel       string
	OtherFixedInstallation_5_5             string
	OtherFixedInstallation_5_5_Risklevel   string
	HoseReelsProvided_5_6                  string
	HoseReelsProvided_5_6_Risklevel        string
	FixedInstallationsTested_5_7           string
	FixedInstallationsTested_5_7_Risklevel string
	ProtectionSystemsSignage_5_8           string
	ProtectionSystemsSignage_5_8_Risklevel string
	ProtectionSystemsAlarms_5_9            string
	ProtectionSystemsAlarms_5_9_Risklevel  string

	//Arson
	Arson_6_0_Comment        string
	Arson_6_0_Recommendation string

	SecureDuringDarkness_6_1                string
	SecureDuringDarkness_6_1_Risklevel      string
	StandardExternalLighting_6_2            string
	StandardExternalLighting_6_2_Risklevel  string
	ExternalRubbish_6_3                     string
	ExternalRubbish_6_3_Risklevel           string
	WheelieBinsSecurelyPlaced_6_4           string
	WheelieBinsSecurelyPlaced_6_4_Risklevel string
	CCTVAvailable_6_5                       string
	CCTVAvailable_6_5_Risklevel             string
	VideoDoorbellAvailable_6_6              string
	VideoDoorbellAvailable_6_6_Risklevel    string

	//7.0 Maintenance and Record Keeping
	MaintenanceAndRecordKeeping_7_0_Comment        string
	MaintenanceAndRecordKeeping_7_0_Recommendation string

	SufficientRiskAssessment_7_1                  string
	SufficientRiskAssessment_7_1_Risklevel        string
	RiskAssessmentReviewed_7_2                    string
	RiskAssessmentReviewed_7_2_Risklevel          string
	SufficientEmergencyActionPlan_7_3             string
	SufficientEmergencyActionPlan_7_3_Risklevel   string
	LogBookKept_7_4                               string
	LogBookKept_7_4_Risklevel                     string
	FireAlarmTestingRecord_7_5                    string
	FireAlarmTestingRecord_7_5_Risklevel          string
	EmergencyLightTestingRecord_7_6               string
	EmergencyLightTestingRecord_7_6_Risklevel     string
	MeansEscapeRouteRecord_7_7                    string
	MeansEscapeRouteRecord_7_7_Risklevel          string
	FireExtinguisherEquipmentRecord_7_8           string
	FireExtinguisherEquipmentRecord_7_8_Risklevel string
	FixedInstallationTesting_7_9                  string
	FixedInstallationTesting_7_9_Risklevel        string
	GeneralFireTraining_7_10                      string
	GeneralFireTraining_7_10_Risklevel            string
	FireDrillsRecord_7_11                         string
	FireDrillsRecord_7_11_Risklevel               string
}
