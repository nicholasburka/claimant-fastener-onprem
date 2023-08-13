// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirEncounter struct {
	models.ResourceBase
	// The set of accounts that may be used for billing for this Encounter
	// https://hl7.org/fhir/r4/search.html#reference
	Account datatypes.JSON `gorm:"column:account;type:text;serializer:json" json:"account,omitempty"`
	// The appointment that scheduled this encounter
	// https://hl7.org/fhir/r4/search.html#reference
	Appointment datatypes.JSON `gorm:"column:appointment;type:text;serializer:json" json:"appointment,omitempty"`
	// The ServiceRequest that initiated this encounter
	// https://hl7.org/fhir/r4/search.html#reference
	BasedOn datatypes.JSON `gorm:"column:basedOn;type:text;serializer:json" json:"basedOn,omitempty"`
	// Classification of patient encounter
	// https://hl7.org/fhir/r4/search.html#token
	Class datatypes.JSON `gorm:"column:class;type:text;serializer:json" json:"class,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): Date first version of the resource instance was recorded
	   * [CarePlan](careplan.html): Time period plan covers
	   * [CareTeam](careteam.html): Time period team covers
	   * [ClinicalImpression](clinicalimpression.html): When the assessment was documented
	   * [Composition](composition.html): Composition editing time
	   * [Consent](consent.html): When this Consent was created or indexed
	   * [DiagnosticReport](diagnosticreport.html): The clinically relevant time of the report
	   * [Encounter](encounter.html): A date within the period the Encounter lasted
	   * [EpisodeOfCare](episodeofcare.html): The provided date search value falls within the episode of care's period
	   * [FamilyMemberHistory](familymemberhistory.html): When history was recorded or last updated
	   * [Flag](flag.html): Time period when flag is active
	   * [Immunization](immunization.html): Vaccination  (non)-Administration Date
	   * [List](list.html): When the list was prepared
	   * [Observation](observation.html): Obtained date/time. If the obtained element is a period, a date that falls in the period
	   * [Procedure](procedure.html): When the procedure was performed
	   * [RiskAssessment](riskassessment.html): When was assessment made?
	   * [SupplyRequest](supplyrequest.html): When the request was made
	*/
	// https://hl7.org/fhir/r4/search.html#date
	Date *time.Time `gorm:"column:date;type:datetime" json:"date,omitempty"`
	// The diagnosis or procedure relevant to the encounter
	// https://hl7.org/fhir/r4/search.html#reference
	Diagnosis datatypes.JSON `gorm:"column:diagnosis;type:text;serializer:json" json:"diagnosis,omitempty"`
	// Episode(s) of care that this encounter should be recorded against
	// https://hl7.org/fhir/r4/search.html#reference
	EpisodeOfCare datatypes.JSON `gorm:"column:episodeOfCare;type:text;serializer:json" json:"episodeOfCare,omitempty"`
	/*
	   Multiple Resources:

	   * [AllergyIntolerance](allergyintolerance.html): External ids for this item
	   * [CarePlan](careplan.html): External Ids for this plan
	   * [CareTeam](careteam.html): External Ids for this team
	   * [Composition](composition.html): Version-independent identifier for the Composition
	   * [Condition](condition.html): A unique identifier of the condition record
	   * [Consent](consent.html): Identifier for this record (external references)
	   * [DetectedIssue](detectedissue.html): Unique id for the detected issue
	   * [DeviceRequest](devicerequest.html): Business identifier for request/order
	   * [DiagnosticReport](diagnosticreport.html): An identifier for the report
	   * [DocumentManifest](documentmanifest.html): Unique Identifier for the set of documents
	   * [DocumentReference](documentreference.html): Master Version Specific Identifier
	   * [Encounter](encounter.html): Identifier(s) by which this encounter is known
	   * [EpisodeOfCare](episodeofcare.html): Business Identifier(s) relevant for this EpisodeOfCare
	   * [FamilyMemberHistory](familymemberhistory.html): A search by a record identifier
	   * [Goal](goal.html): External Ids for this goal
	   * [ImagingStudy](imagingstudy.html): Identifiers for the Study, such as DICOM Study Instance UID and Accession number
	   * [Immunization](immunization.html): Business identifier
	   * [List](list.html): Business identifier
	   * [MedicationAdministration](medicationadministration.html): Return administrations with this external identifier
	   * [MedicationDispense](medicationdispense.html): Returns dispenses with this external identifier
	   * [MedicationRequest](medicationrequest.html): Return prescriptions with this external identifier
	   * [MedicationStatement](medicationstatement.html): Return statements with this external identifier
	   * [NutritionOrder](nutritionorder.html): Return nutrition orders with this external identifier
	   * [Observation](observation.html): The unique id for a particular observation
	   * [Procedure](procedure.html): A unique identifier for a procedure
	   * [RiskAssessment](riskassessment.html): Unique identifier for the assessment
	   * [ServiceRequest](servicerequest.html): Identifiers assigned to this order
	   * [SupplyDelivery](supplydelivery.html): External identifier
	   * [SupplyRequest](supplyrequest.html): Business Identifier for SupplyRequest
	   * [VisionPrescription](visionprescription.html): Return prescriptions with this external identifier
	*/
	// https://hl7.org/fhir/r4/search.html#token
	Identifier datatypes.JSON `gorm:"column:identifier;type:text;serializer:json" json:"identifier,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	LastUpdated *time.Time `gorm:"column:lastUpdated;type:datetime" json:"lastUpdated,omitempty"`
	// Length of encounter in days
	// https://hl7.org/fhir/r4/search.html#quantity
	Length datatypes.JSON `gorm:"column:length;type:text;serializer:json" json:"length,omitempty"`
	// Location the encounter takes place
	// https://hl7.org/fhir/r4/search.html#reference
	Location datatypes.JSON `gorm:"column:location;type:text;serializer:json" json:"location,omitempty"`
	// Time period during which the patient was present at the location
	// https://hl7.org/fhir/r4/search.html#date
	LocationPeriod *time.Time `gorm:"column:locationPeriod;type:datetime" json:"locationPeriod,omitempty"`
	// Another Encounter this encounter is part of
	// https://hl7.org/fhir/r4/search.html#reference
	PartOf datatypes.JSON `gorm:"column:partOf;type:text;serializer:json" json:"partOf,omitempty"`
	// Persons involved in the encounter other than the patient
	// https://hl7.org/fhir/r4/search.html#reference
	Participant datatypes.JSON `gorm:"column:participant;type:text;serializer:json" json:"participant,omitempty"`
	// Role of participant in encounter
	// https://hl7.org/fhir/r4/search.html#token
	ParticipantType datatypes.JSON `gorm:"column:participantType;type:text;serializer:json" json:"participantType,omitempty"`
	// Persons involved in the encounter other than the patient
	// https://hl7.org/fhir/r4/search.html#reference
	Practitioner datatypes.JSON `gorm:"column:practitioner;type:text;serializer:json" json:"practitioner,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	Profile datatypes.JSON `gorm:"column:profile;type:text;serializer:json" json:"profile,omitempty"`
	// Coded reason the encounter takes place
	// https://hl7.org/fhir/r4/search.html#token
	ReasonCode datatypes.JSON `gorm:"column:reasonCode;type:text;serializer:json" json:"reasonCode,omitempty"`
	// Reason the encounter takes place (reference)
	// https://hl7.org/fhir/r4/search.html#reference
	ReasonReference datatypes.JSON `gorm:"column:reasonReference;type:text;serializer:json" json:"reasonReference,omitempty"`
	// The organization (facility) responsible for this encounter
	// https://hl7.org/fhir/r4/search.html#reference
	ServiceProvider datatypes.JSON `gorm:"column:serviceProvider;type:text;serializer:json" json:"serviceProvider,omitempty"`
	// Identifies where the resource comes from
	// https://hl7.org/fhir/r4/search.html#uri
	SourceUri string `gorm:"column:sourceUri;type:text" json:"sourceUri,omitempty"`
	// Wheelchair, translator, stretcher, etc.
	// https://hl7.org/fhir/r4/search.html#token
	SpecialArrangement datatypes.JSON `gorm:"column:specialArrangement;type:text;serializer:json" json:"specialArrangement,omitempty"`
	// planned | arrived | triaged | in-progress | onleave | finished | cancelled +
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// The patient or group present at the encounter
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	Tag datatypes.JSON `gorm:"column:tag;type:text;serializer:json" json:"tag,omitempty"`
	// Text search against the narrative
	// https://hl7.org/fhir/r4/search.html#string
	Text datatypes.JSON `gorm:"column:text;type:text;serializer:json" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirEncounter) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"account":            "reference",
		"appointment":        "reference",
		"basedOn":            "reference",
		"class":              "token",
		"date":               "date",
		"diagnosis":          "reference",
		"episodeOfCare":      "reference",
		"identifier":         "token",
		"language":           "token",
		"lastUpdated":        "date",
		"length":             "quantity",
		"location":           "reference",
		"locationPeriod":     "date",
		"partOf":             "reference",
		"participant":        "reference",
		"participantType":    "token",
		"practitioner":       "reference",
		"profile":            "reference",
		"reasonCode":         "token",
		"reasonReference":    "reference",
		"serviceProvider":    "reference",
		"sourceUri":          "uri",
		"specialArrangement": "token",
		"status":             "token",
		"subject":            "reference",
		"tag":                "token",
		"text":               "string",
		"type":               "special",
	}
	return searchParameters
}
func (s *FhirEncounter) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
	s.ResourceRaw = datatypes.JSON(resourceRaw)
	// unmarshal the raw resource (bytes) into a map
	var resourceRawMap map[string]interface{}
	err := json.Unmarshal(resourceRaw, &resourceRawMap)
	if err != nil {
		return err
	}
	if len(fhirPathJs) == 0 {
		return fmt.Errorf("fhirPathJs script is empty")
	}
	vm := goja.New()
	// setup the global window object
	vm.Set("window", vm.NewObject())
	// set the global FHIR Resource object
	vm.Set("fhirResource", resourceRawMap)
	// compile the fhirpath library
	fhirPathJsProgram, err := goja.Compile("fhirpath.min.js", fhirPathJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting Account
	accountResult, err := vm.RunString(` 
							AccountResult = window.fhirpath.evaluate(fhirResource, 'Encounter.account')
						
							if(AccountResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(AccountResult)
							}
						 `)
	if err == nil && accountResult.String() != "undefined" {
		s.Account = []byte(accountResult.String())
	}
	// extracting Appointment
	appointmentResult, err := vm.RunString(` 
							AppointmentResult = window.fhirpath.evaluate(fhirResource, 'Encounter.appointment')
						
							if(AppointmentResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(AppointmentResult)
							}
						 `)
	if err == nil && appointmentResult.String() != "undefined" {
		s.Appointment = []byte(appointmentResult.String())
	}
	// extracting BasedOn
	basedOnResult, err := vm.RunString(` 
							BasedOnResult = window.fhirpath.evaluate(fhirResource, 'Encounter.basedOn')
						
							if(BasedOnResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(BasedOnResult)
							}
						 `)
	if err == nil && basedOnResult.String() != "undefined" {
		s.BasedOn = []byte(basedOnResult.String())
	}
	// extracting Class
	classResult, err := vm.RunString(` 
							ClassResult = window.fhirpath.evaluate(fhirResource, 'Encounter.class')
							ClassProcessed = ClassResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(ClassProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ClassProcessed)
							}
						 `)
	if err == nil && classResult.String() != "undefined" {
		s.Class = []byte(classResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.recordedDate | CarePlan.period | CareTeam.period | ClinicalImpression.date | Composition.date | Consent.dateTime | DiagnosticReport.effective | Encounter.period | EpisodeOfCare.period | FamilyMemberHistory.date | Flag.period | (Immunization.occurrenceDateTime) | List.date | Observation.effective | Procedure.performed | (RiskAssessment.occurrenceDateTime) | SupplyRequest.authoredOn')[0]")
	if err == nil && dateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, dateResult.String())
		if err == nil {
			s.Date = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", dateResult.String())
			if err == nil {
				s.Date = &d
			}
		}
	}
	// extracting Diagnosis
	diagnosisResult, err := vm.RunString(` 
							DiagnosisResult = window.fhirpath.evaluate(fhirResource, 'Encounter.diagnosis.condition')
						
							if(DiagnosisResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(DiagnosisResult)
							}
						 `)
	if err == nil && diagnosisResult.String() != "undefined" {
		s.Diagnosis = []byte(diagnosisResult.String())
	}
	// extracting EpisodeOfCare
	episodeOfCareResult, err := vm.RunString(` 
							EpisodeOfCareResult = window.fhirpath.evaluate(fhirResource, 'Encounter.episodeOfCare')
						
							if(EpisodeOfCareResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(EpisodeOfCareResult)
							}
						 `)
	if err == nil && episodeOfCareResult.String() != "undefined" {
		s.EpisodeOfCare = []byte(episodeOfCareResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString(` 
							IdentifierResult = window.fhirpath.evaluate(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')
							IdentifierProcessed = IdentifierResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(IdentifierProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(IdentifierProcessed)
							}
						 `)
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString(` 
							LanguageResult = window.fhirpath.evaluate(fhirResource, 'language')
							LanguageProcessed = LanguageResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(LanguageProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LanguageProcessed)
							}
						 `)
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting LastUpdated
	lastUpdatedResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'meta.lastUpdated')[0]")
	if err == nil && lastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, lastUpdatedResult.String())
		if err == nil {
			s.LastUpdated = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", lastUpdatedResult.String())
			if err == nil {
				s.LastUpdated = &d
			}
		}
	}
	// extracting Length
	lengthResult, err := vm.RunString("JSON.stringify(window.fhirpath.evaluate(fhirResource, 'Encounter.length'))")
	if err == nil && lengthResult.String() != "undefined" {
		s.Length = []byte(lengthResult.String())
	}
	// extracting Location
	locationResult, err := vm.RunString(` 
							LocationResult = window.fhirpath.evaluate(fhirResource, 'Encounter.location.location')
						
							if(LocationResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(LocationResult)
							}
						 `)
	if err == nil && locationResult.String() != "undefined" {
		s.Location = []byte(locationResult.String())
	}
	// extracting LocationPeriod
	locationPeriodResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'Encounter.location.period')[0]")
	if err == nil && locationPeriodResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, locationPeriodResult.String())
		if err == nil {
			s.LocationPeriod = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", locationPeriodResult.String())
			if err == nil {
				s.LocationPeriod = &d
			}
		}
	}
	// extracting PartOf
	partOfResult, err := vm.RunString(` 
							PartOfResult = window.fhirpath.evaluate(fhirResource, 'Encounter.partOf')
						
							if(PartOfResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(PartOfResult)
							}
						 `)
	if err == nil && partOfResult.String() != "undefined" {
		s.PartOf = []byte(partOfResult.String())
	}
	// extracting Participant
	participantResult, err := vm.RunString(` 
							ParticipantResult = window.fhirpath.evaluate(fhirResource, 'Encounter.participant.individual')
						
							if(ParticipantResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ParticipantResult)
							}
						 `)
	if err == nil && participantResult.String() != "undefined" {
		s.Participant = []byte(participantResult.String())
	}
	// extracting ParticipantType
	participantTypeResult, err := vm.RunString(` 
							ParticipantTypeResult = window.fhirpath.evaluate(fhirResource, 'Encounter.participant.type')
							ParticipantTypeProcessed = ParticipantTypeResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(ParticipantTypeProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ParticipantTypeProcessed)
							}
						 `)
	if err == nil && participantTypeResult.String() != "undefined" {
		s.ParticipantType = []byte(participantTypeResult.String())
	}
	// extracting Practitioner
	practitionerResult, err := vm.RunString(` 
							PractitionerResult = window.fhirpath.evaluate(fhirResource, 'Encounter.participant.individual.where(resolve() is Practitioner)')
						
							if(PractitionerResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(PractitionerResult)
							}
						 `)
	if err == nil && practitionerResult.String() != "undefined" {
		s.Practitioner = []byte(practitionerResult.String())
	}
	// extracting Profile
	profileResult, err := vm.RunString(` 
							ProfileResult = window.fhirpath.evaluate(fhirResource, 'meta.profile')
						
							if(ProfileResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ProfileResult)
							}
						 `)
	if err == nil && profileResult.String() != "undefined" {
		s.Profile = []byte(profileResult.String())
	}
	// extracting ReasonCode
	reasonCodeResult, err := vm.RunString(` 
							ReasonCodeResult = window.fhirpath.evaluate(fhirResource, 'Encounter.reasonCode')
							ReasonCodeProcessed = ReasonCodeResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(ReasonCodeProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ReasonCodeProcessed)
							}
						 `)
	if err == nil && reasonCodeResult.String() != "undefined" {
		s.ReasonCode = []byte(reasonCodeResult.String())
	}
	// extracting ReasonReference
	reasonReferenceResult, err := vm.RunString(` 
							ReasonReferenceResult = window.fhirpath.evaluate(fhirResource, 'Encounter.reasonReference')
						
							if(ReasonReferenceResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ReasonReferenceResult)
							}
						 `)
	if err == nil && reasonReferenceResult.String() != "undefined" {
		s.ReasonReference = []byte(reasonReferenceResult.String())
	}
	// extracting ServiceProvider
	serviceProviderResult, err := vm.RunString(` 
							ServiceProviderResult = window.fhirpath.evaluate(fhirResource, 'Encounter.serviceProvider')
						
							if(ServiceProviderResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(ServiceProviderResult)
							}
						 `)
	if err == nil && serviceProviderResult.String() != "undefined" {
		s.ServiceProvider = []byte(serviceProviderResult.String())
	}
	// extracting SourceUri
	sourceUriResult, err := vm.RunString("window.fhirpath.evaluate(fhirResource, 'meta.source')[0]")
	if err == nil && sourceUriResult.String() != "undefined" {
		s.SourceUri = sourceUriResult.String()
	}
	// extracting SpecialArrangement
	specialArrangementResult, err := vm.RunString(` 
							SpecialArrangementResult = window.fhirpath.evaluate(fhirResource, 'Encounter.hospitalization.specialArrangement')
							SpecialArrangementProcessed = SpecialArrangementResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(SpecialArrangementProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SpecialArrangementProcessed)
							}
						 `)
	if err == nil && specialArrangementResult.String() != "undefined" {
		s.SpecialArrangement = []byte(specialArrangementResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString(` 
							StatusResult = window.fhirpath.evaluate(fhirResource, 'Encounter.status')
							StatusProcessed = StatusResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(StatusProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(StatusProcessed)
							}
						 `)
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString(` 
							SubjectResult = window.fhirpath.evaluate(fhirResource, 'Encounter.subject')
						
							if(SubjectResult.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(SubjectResult)
							}
						 `)
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	// extracting Tag
	tagResult, err := vm.RunString(` 
							TagResult = window.fhirpath.evaluate(fhirResource, 'meta.tag')
							TagProcessed = TagResult.reduce((accumulator, currentValue) => {
								if (currentValue.coding) {
									//CodeableConcept
									currentValue.coding.map((coding) => {
										accumulator.push({
											"code": coding.code,	
											"system": coding.system,
											"text": currentValue.text
										})
									})
								} else if (currentValue.value) {
									//ContactPoint, Identifier
									accumulator.push({
										"code": currentValue.value,
										"system": currentValue.system,
										"text": currentValue.type?.text
									})
								} else if (currentValue.code) {
									//Coding
									accumulator.push({
										"code": currentValue.code,
										"system": currentValue.system,
										"text": currentValue.display
									})
								} else if ((typeof currentValue === 'string') || (typeof currentValue === 'boolean')) {
									//string, boolean
									accumulator.push({
										"code": currentValue,
									})
								}
								return accumulator
							}, [])
						
				
							if(TagProcessed.length == 0) {
								"undefined"
							}
 							else {
								JSON.stringify(TagProcessed)
							}
						 `)
	if err == nil && tagResult.String() != "undefined" {
		s.Tag = []byte(tagResult.String())
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirEncounter) TableName() string {
	return "fhir_encounter"
}
