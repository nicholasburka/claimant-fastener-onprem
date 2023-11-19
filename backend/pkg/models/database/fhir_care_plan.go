// THIS FILE IS GENERATED BY https://github.com/fastenhealth/fasten-onprem/blob/main/backend/pkg/models/database/generate.go
// PLEASE DO NOT EDIT BY HAND

package database

import (
	"encoding/json"
	"fmt"
	goja "github.com/dop251/goja"
	models "github.com/fastenhealth/fasten-onprem/backend/pkg/models"
	datatypes "gorm.io/datatypes"
	"time"
)

type FhirCarePlan struct {
	models.ResourceBase
	// Detail type of activity
	// https://hl7.org/fhir/r4/search.html#token
	ActivityCode datatypes.JSON `gorm:"column:activityCode;type:text;serializer:json" json:"activityCode,omitempty"`
	// Specified date occurs within period specified by CarePlan.activity.detail.scheduled[x]
	// https://hl7.org/fhir/r4/search.html#date
	ActivityDate *time.Time `gorm:"column:activityDate;type:datetime" json:"activityDate,omitempty"`
	// Activity details defined in specific resource
	// https://hl7.org/fhir/r4/search.html#reference
	ActivityReference datatypes.JSON `gorm:"column:activityReference;type:text;serializer:json" json:"activityReference,omitempty"`
	// Fulfills CarePlan
	// https://hl7.org/fhir/r4/search.html#reference
	BasedOn datatypes.JSON `gorm:"column:basedOn;type:text;serializer:json" json:"basedOn,omitempty"`
	// Who's involved in plan?
	// https://hl7.org/fhir/r4/search.html#reference
	CareTeam datatypes.JSON `gorm:"column:careTeam;type:text;serializer:json" json:"careTeam,omitempty"`
	// Type of plan
	// https://hl7.org/fhir/r4/search.html#token
	Category datatypes.JSON `gorm:"column:category;type:text;serializer:json" json:"category,omitempty"`
	// Health issues this plan addresses
	// https://hl7.org/fhir/r4/search.html#reference
	Condition datatypes.JSON `gorm:"column:condition;type:text;serializer:json" json:"condition,omitempty"`
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
	// Encounter created as part of
	// https://hl7.org/fhir/r4/search.html#reference
	Encounter datatypes.JSON `gorm:"column:encounter;type:text;serializer:json" json:"encounter,omitempty"`
	// Desired outcome of plan
	// https://hl7.org/fhir/r4/search.html#reference
	Goal datatypes.JSON `gorm:"column:goal;type:text;serializer:json" json:"goal,omitempty"`
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
	// Instantiates FHIR protocol or definition
	// https://hl7.org/fhir/r4/search.html#reference
	InstantiatesCanonical datatypes.JSON `gorm:"column:instantiatesCanonical;type:text;serializer:json" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	// https://hl7.org/fhir/r4/search.html#uri
	InstantiatesUri string `gorm:"column:instantiatesUri;type:text" json:"instantiatesUri,omitempty"`
	// proposal | plan | order | option
	// https://hl7.org/fhir/r4/search.html#token
	Intent datatypes.JSON `gorm:"column:intent;type:text;serializer:json" json:"intent,omitempty"`
	// Language of the resource content
	// https://hl7.org/fhir/r4/search.html#token
	Language datatypes.JSON `gorm:"column:language;type:text;serializer:json" json:"language,omitempty"`
	// When the resource version last changed
	// https://hl7.org/fhir/r4/search.html#date
	MetaLastUpdated *time.Time `gorm:"column:metaLastUpdated;type:datetime" json:"metaLastUpdated,omitempty"`
	// Profiles this resource claims to conform to
	// https://hl7.org/fhir/r4/search.html#reference
	MetaProfile datatypes.JSON `gorm:"column:metaProfile;type:text;serializer:json" json:"metaProfile,omitempty"`
	// Tags applied to this resource
	// https://hl7.org/fhir/r4/search.html#token
	MetaTag datatypes.JSON `gorm:"column:metaTag;type:text;serializer:json" json:"metaTag,omitempty"`
	// Tags applied to this resource
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	MetaVersionId string `gorm:"column:metaVersionId;type:text" json:"metaVersionId,omitempty"`
	// Part of referenced CarePlan
	// https://hl7.org/fhir/r4/search.html#reference
	PartOf datatypes.JSON `gorm:"column:partOf;type:text;serializer:json" json:"partOf,omitempty"`
	// Matches if the practitioner is listed as a performer in any of the "simple" activities.  (For performers of the detailed activities, chain through the activitydetail search parameter.)
	// https://hl7.org/fhir/r4/search.html#reference
	Performer datatypes.JSON `gorm:"column:performer;type:text;serializer:json" json:"performer,omitempty"`
	// CarePlan replaced by this CarePlan
	// https://hl7.org/fhir/r4/search.html#reference
	Replaces datatypes.JSON `gorm:"column:replaces;type:text;serializer:json" json:"replaces,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	// https://hl7.org/fhir/r4/search.html#token
	Status datatypes.JSON `gorm:"column:status;type:text;serializer:json" json:"status,omitempty"`
	// Who the care plan is for
	// https://hl7.org/fhir/r4/search.html#reference
	Subject datatypes.JSON `gorm:"column:subject;type:text;serializer:json" json:"subject,omitempty"`
	// Text search against the narrative
	// This is a primitive string literal (`keyword` type). It is not a recognized SearchParameter type from https://hl7.org/fhir/r4/search.html, it's Fasten Health-specific
	Text string `gorm:"column:text;type:text" json:"text,omitempty"`
	// A resource type filter
	// https://hl7.org/fhir/r4/search.html#special
	Type datatypes.JSON `gorm:"column:type;type:text;serializer:json" json:"type,omitempty"`
}

func (s *FhirCarePlan) GetSearchParameters() map[string]string {
	searchParameters := map[string]string{
		"activityCode":          "token",
		"activityDate":          "date",
		"activityReference":     "reference",
		"basedOn":               "reference",
		"careTeam":              "reference",
		"category":              "token",
		"condition":             "reference",
		"date":                  "date",
		"encounter":             "reference",
		"goal":                  "reference",
		"id":                    "keyword",
		"identifier":            "token",
		"instantiatesCanonical": "reference",
		"instantiatesUri":       "uri",
		"intent":                "token",
		"language":              "token",
		"metaLastUpdated":       "date",
		"metaProfile":           "reference",
		"metaTag":               "token",
		"metaVersionId":         "keyword",
		"partOf":                "reference",
		"performer":             "reference",
		"replaces":              "reference",
		"sort_date":             "date",
		"source_id":             "keyword",
		"source_resource_id":    "keyword",
		"source_resource_type":  "keyword",
		"source_uri":            "keyword",
		"status":                "token",
		"subject":               "reference",
		"text":                  "keyword",
		"type":                  "special",
	}
	return searchParameters
}
func (s *FhirCarePlan) PopulateAndExtractSearchParameters(resourceRaw json.RawMessage) error {
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
	// compile the searchParametersExtractor library
	searchParametersExtractorJsProgram, err := goja.Compile("searchParameterExtractor.js", searchParameterExtractorJs, true)
	if err != nil {
		return err
	}
	// add the fhirpath library in the goja vm
	_, err = vm.RunProgram(fhirPathJsProgram)
	if err != nil {
		return err
	}
	// add the searchParametersExtractor library in the goja vm
	_, err = vm.RunProgram(searchParametersExtractorJsProgram)
	if err != nil {
		return err
	}
	// execute the fhirpath expression for each search parameter
	// extracting ActivityCode
	activityCodeResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'CarePlan.activity.detail.code')")
	if err == nil && activityCodeResult.String() != "undefined" {
		s.ActivityCode = []byte(activityCodeResult.String())
	}
	// extracting ActivityDate
	activityDateResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'CarePlan.activity.detail.scheduledTiming | CarePlan.activity.detail.scheduledPeriod | CarePlan.activity.detail.scheduledString')")
	if err == nil && activityDateResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, activityDateResult.String())
		if err == nil {
			s.ActivityDate = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", activityDateResult.String())
			if err == nil {
				s.ActivityDate = &d
			}
		}
	}
	// extracting ActivityReference
	activityReferenceResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.activity.reference')")
	if err == nil && activityReferenceResult.String() != "undefined" {
		s.ActivityReference = []byte(activityReferenceResult.String())
	}
	// extracting BasedOn
	basedOnResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.basedOn')")
	if err == nil && basedOnResult.String() != "undefined" {
		s.BasedOn = []byte(basedOnResult.String())
	}
	// extracting CareTeam
	careTeamResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.careTeam')")
	if err == nil && careTeamResult.String() != "undefined" {
		s.CareTeam = []byte(careTeamResult.String())
	}
	// extracting Category
	categoryResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'CarePlan.category')")
	if err == nil && categoryResult.String() != "undefined" {
		s.Category = []byte(categoryResult.String())
	}
	// extracting Condition
	conditionResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.addresses')")
	if err == nil && conditionResult.String() != "undefined" {
		s.Condition = []byte(conditionResult.String())
	}
	// extracting Date
	dateResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'AllergyIntolerance.recordedDate | CarePlan.period | CareTeam.period | ClinicalImpression.date | Composition.date | Consent.dateTime | DiagnosticReport.effectiveDateTime | DiagnosticReport.effectivePeriod | Encounter.period | EpisodeOfCare.period | FamilyMemberHistory.date | Flag.period | (Immunization.occurrenceDateTime) | List.date | Observation.effectiveDateTime | Observation.effectivePeriod | Observation.effectiveTiming | Observation.effectiveInstant | Procedure.performedDateTime | Procedure.performedPeriod | Procedure.performedString | Procedure.performedAge | Procedure.performedRange | (RiskAssessment.occurrenceDateTime) | SupplyRequest.authoredOn')")
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
	// extracting Encounter
	encounterResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.encounter')")
	if err == nil && encounterResult.String() != "undefined" {
		s.Encounter = []byte(encounterResult.String())
	}
	// extracting Goal
	goalResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.goal')")
	if err == nil && goalResult.String() != "undefined" {
		s.Goal = []byte(goalResult.String())
	}
	// extracting Identifier
	identifierResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'AllergyIntolerance.identifier | CarePlan.identifier | CareTeam.identifier | Composition.identifier | Condition.identifier | Consent.identifier | DetectedIssue.identifier | DeviceRequest.identifier | DiagnosticReport.identifier | DocumentManifest.masterIdentifier | DocumentManifest.identifier | DocumentReference.masterIdentifier | DocumentReference.identifier | Encounter.identifier | EpisodeOfCare.identifier | FamilyMemberHistory.identifier | Goal.identifier | ImagingStudy.identifier | Immunization.identifier | List.identifier | MedicationAdministration.identifier | MedicationDispense.identifier | MedicationRequest.identifier | MedicationStatement.identifier | NutritionOrder.identifier | Observation.identifier | Procedure.identifier | RiskAssessment.identifier | ServiceRequest.identifier | SupplyDelivery.identifier | SupplyRequest.identifier | VisionPrescription.identifier')")
	if err == nil && identifierResult.String() != "undefined" {
		s.Identifier = []byte(identifierResult.String())
	}
	// extracting InstantiatesCanonical
	instantiatesCanonicalResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.instantiatesCanonical')")
	if err == nil && instantiatesCanonicalResult.String() != "undefined" {
		s.InstantiatesCanonical = []byte(instantiatesCanonicalResult.String())
	}
	// extracting InstantiatesUri
	instantiatesUriResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'CarePlan.instantiatesUri')")
	if err == nil && instantiatesUriResult.String() != "undefined" {
		s.InstantiatesUri = instantiatesUriResult.String()
	}
	// extracting Intent
	intentResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'CarePlan.intent')")
	if err == nil && intentResult.String() != "undefined" {
		s.Intent = []byte(intentResult.String())
	}
	// extracting Language
	languageResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'language')")
	if err == nil && languageResult.String() != "undefined" {
		s.Language = []byte(languageResult.String())
	}
	// extracting MetaLastUpdated
	metaLastUpdatedResult, err := vm.RunString("extractDateSearchParameters(fhirResource, 'meta.lastUpdated')")
	if err == nil && metaLastUpdatedResult.String() != "undefined" {
		t, err := time.Parse(time.RFC3339, metaLastUpdatedResult.String())
		if err == nil {
			s.MetaLastUpdated = &t
		} else if err != nil {
			d, err := time.Parse("2006-01-02", metaLastUpdatedResult.String())
			if err == nil {
				s.MetaLastUpdated = &d
			}
		}
	}
	// extracting MetaProfile
	metaProfileResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'meta.profile')")
	if err == nil && metaProfileResult.String() != "undefined" {
		s.MetaProfile = []byte(metaProfileResult.String())
	}
	// extracting MetaTag
	metaTagResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'meta.tag')")
	if err == nil && metaTagResult.String() != "undefined" {
		s.MetaTag = []byte(metaTagResult.String())
	}
	// extracting MetaVersionId
	metaVersionIdResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'meta.versionId')")
	if err == nil && metaVersionIdResult.String() != "undefined" {
		s.MetaVersionId = metaVersionIdResult.String()
	}
	// extracting PartOf
	partOfResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.partOf')")
	if err == nil && partOfResult.String() != "undefined" {
		s.PartOf = []byte(partOfResult.String())
	}
	// extracting Performer
	performerResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.activity.detail.performer')")
	if err == nil && performerResult.String() != "undefined" {
		s.Performer = []byte(performerResult.String())
	}
	// extracting Replaces
	replacesResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.replaces')")
	if err == nil && replacesResult.String() != "undefined" {
		s.Replaces = []byte(replacesResult.String())
	}
	// extracting Status
	statusResult, err := vm.RunString("extractTokenSearchParameters(fhirResource, 'CarePlan.status')")
	if err == nil && statusResult.String() != "undefined" {
		s.Status = []byte(statusResult.String())
	}
	// extracting Subject
	subjectResult, err := vm.RunString("extractReferenceSearchParameters(fhirResource, 'CarePlan.subject')")
	if err == nil && subjectResult.String() != "undefined" {
		s.Subject = []byte(subjectResult.String())
	}
	// extracting Text
	textResult, err := vm.RunString("extractSimpleSearchParameters(fhirResource, 'text')")
	if err == nil && textResult.String() != "undefined" {
		s.Text = textResult.String()
	}
	return nil
}

// TableName overrides the table name from fhir_observations (pluralized) to `fhir_observation`. https://gorm.io/docs/conventions.html#TableName
func (s *FhirCarePlan) TableName() string {
	return "fhir_care_plan"
}
