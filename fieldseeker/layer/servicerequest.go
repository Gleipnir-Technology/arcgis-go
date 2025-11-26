package layer

import (
	"time"

	"github.com/google/uuid"
)

type ServiceRequestServiceRequestTargetType string

const (
	ServiceRequestServiceRequestTargetRequestfreemosquitofish                              ServiceRequestServiceRequestTargetType = "mosquitofish"
	ServiceRequestServiceRequestTargetReportaneglectedpoolorspa                            ServiceRequestServiceRequestTargetType = "neglected pool or spa"
	ServiceRequestServiceRequestTargetReportstandingwater                                  ServiceRequestServiceRequestTargetType = "standing water"
	ServiceRequestServiceRequestTargetReportmosquitopresence                               ServiceRequestServiceRequestTargetType = "mosquito presence"
	ServiceRequestServiceRequestTargetReportdaybitingmosquitoes                            ServiceRequestServiceRequestTargetType = "biting mosquitoes"
	ServiceRequestServiceRequestTargetRequestarepresentativeforeducationaloroutreachevents ServiceRequestServiceRequestTargetType = "event"
	ServiceRequestServiceRequestTargetWebRequestforFish                                    ServiceRequestServiceRequestTargetType = "fish"
	ServiceRequestServiceRequestTargetWebReportofMosquitoes                                ServiceRequestServiceRequestTargetType = "mosquito"
	ServiceRequestServiceRequestTargetWebReportofMosquitoSource                            ServiceRequestServiceRequestTargetType = "source"
	ServiceRequestServiceRequestTargetWebReportofDeadBird                                  ServiceRequestServiceRequestTargetType = "bird"
)

type ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type string

const (
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5AlysiaDavis       ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Alysia Davis"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5AlejandraGill     ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Alejandra Gill"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5AndreaTroupin     ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Andrea Troupin"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5BrendaRodriguez   ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Brenda Rodriguez"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5BryanFerguson     ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Bryan Ferguson"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5BryanRuiz         ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Bryan Ruiz"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5ConlinReis        ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Conlin Reis"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5CarlosRodriguez   ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Carlos Rodriguez"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5ErickArriga       ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Erick Arriga"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5LandonMcGill      ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Landon McGill"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5MarcoMartinez     ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Marco Martinez"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5MarkNakata        ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Mark Nakata"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5MarioSanchez      ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Mario Sanchez"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5PabloOrtega       ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Juan Pablo Ortega"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5RyanSpratt        ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Ryan Spratt"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5TedMcGill         ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Ted McGill"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5BenjaminSperry    ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Benjamin Sperry"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5ZacheryBarragan   ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Zachery Barragan"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5ArturoGarciaTrejo ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Arturo Garcia-Trejo"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5JesusSolano       ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Jesus Jolano"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5YajairaGodinez    ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Yajaira Godinez"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5JakeMaldonado     ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Jake Maldonado"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5RafaelRamirez     ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Rafael Ramirez"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5LisaSalgado       ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Lisa Salgado"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5KoryWilson        ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Kory Wilson"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5CarlosPalacios    ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Carlos Palacios"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5FatimaHidalgo     ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Fatima Hidalgo"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5AaronFredrick     ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Aaron Fredrick"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5JoshMalone        ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Josh Malone"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5JorgePerez        ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Jorge Perez"
	ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5LauraRomos        ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type = "Laura Romos"
)

type ServiceRequestNotInUITFType int16

const (
	ServiceRequestNotInUITFTrue  ServiceRequestNotInUITFType = 1
	ServiceRequestNotInUITFFalse ServiceRequestNotInUITFType = 0
)

type ServiceRequestServiceRequestNextActionType string

const (
	ServiceRequestServiceRequestNextActionNightspray ServiceRequestServiceRequestNextActionType = "Night spray"
	ServiceRequestServiceRequestNextActionSitevisit  ServiceRequestServiceRequestNextActionType = "Site visit"
)

type ServiceRequestServiceRequestIssuesType string

const (
	ServiceRequestServiceRequestIssuesBeehiveRelated           ServiceRequestServiceRequestIssuesType = "Beehive Related"
	ServiceRequestServiceRequestIssuesUnsanitaryAccumulations  ServiceRequestServiceRequestIssuesType = "Unsanitary Accumulations"
	ServiceRequestServiceRequestIssuesRoosterorNoise           ServiceRequestServiceRequestIssuesType = "Rooster or Noise"
	ServiceRequestServiceRequestIssuesRatsAttracted            ServiceRequestServiceRequestIssuesType = "Rats Attracted"
	ServiceRequestServiceRequestIssuesOdor                     ServiceRequestServiceRequestIssuesType = "Odor"
	ServiceRequestServiceRequestIssuesNumberofAnimalsOverLimit ServiceRequestServiceRequestIssuesType = "Number of Animals Over Limit"
	ServiceRequestServiceRequestIssuesLocation                 ServiceRequestServiceRequestIssuesType = "Location"
	ServiceRequestServiceRequestIssuesViolation                ServiceRequestServiceRequestIssuesType = "Violation"
	ServiceRequestServiceRequestIssuesInadequateEnclosure      ServiceRequestServiceRequestIssuesType = "Inadequate Enclosure"
	ServiceRequestServiceRequestIssuesEscapedAnimal            ServiceRequestServiceRequestIssuesType = "Escaped Animal"
	ServiceRequestServiceRequestIssuesIllegalAnimal            ServiceRequestServiceRequestIssuesType = "Illegal Animal"
)

type ServiceRequestServiceRequestscheduleperiod3f40c046afd14abd8bf4389650d29a49Type string

const (
	ServiceRequestServiceRequestscheduleperiod3f40c046afd14abd8bf4389650d29a49AM ServiceRequestServiceRequestscheduleperiod3f40c046afd14abd8bf4389650d29a49Type = "AM"
	ServiceRequestServiceRequestscheduleperiod3f40c046afd14abd8bf4389650d29a49PM ServiceRequestServiceRequestscheduleperiod3f40c046afd14abd8bf4389650d29a49Type = "PM"
)

type ServiceRequestServiceRequestSpanishaaa3dc669f9a45278ecdc9f76db33879Type int32

const (
	ServiceRequestServiceRequestSpanishaaa3dc669f9a45278ecdc9f76db33879NoUknown ServiceRequestServiceRequestSpanishaaa3dc669f9a45278ecdc9f76db33879Type = 0
	ServiceRequestServiceRequestSpanishaaa3dc669f9a45278ecdc9f76db33879Yes      ServiceRequestServiceRequestSpanishaaa3dc669f9a45278ecdc9f76db33879Type = 1
)

type ServiceRequestServiceRequestPriorityType string

const (
	ServiceRequestServiceRequestPriorityLow                     ServiceRequestServiceRequestPriorityType = "Low"
	ServiceRequestServiceRequestPriorityMedium                  ServiceRequestServiceRequestPriorityType = "Medium"
	ServiceRequestServiceRequestPriorityHigh                    ServiceRequestServiceRequestPriorityType = "High"
	ServiceRequestServiceRequestPriorityFollowupVisit           ServiceRequestServiceRequestPriorityType = "Follow up Visit"
	ServiceRequestServiceRequestPriorityHTCResponse             ServiceRequestServiceRequestPriorityType = "HTC Response"
	ServiceRequestServiceRequestPriorityDiseaseAcitivtyResponse ServiceRequestServiceRequestPriorityType = "Disease Activity Response"
)

type ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aType string

const (
	ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aRickAlverez   ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aType = "Rick Alverez"
	ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aBryanFerguson ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aType = "Bryan Ferguson"
	ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aBryanRuiz     ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aType = "Bryan Ruiz"
	ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aAndreaTroupin ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aType = "Andrea Troupin"
	ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aConlinReis    ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aType = "Conlin Reis"
)

type ServiceRequestServiceRequestRegionType string

const (
	ServiceRequestServiceRequestRegionFlorida ServiceRequestServiceRequestRegionType = "FL"
	ServiceRequestServiceRequestRegionIdaho   ServiceRequestServiceRequestRegionType = "ID"
)

type ServiceRequestServiceRequestRejectedReasonType string

const (
	ServiceRequestServiceRequestRejectedReasonDistance ServiceRequestServiceRequestRejectedReasonType = "Distance"
	ServiceRequestServiceRequestRejectedReasonWorkload ServiceRequestServiceRequestRejectedReasonType = "Workload"
)

type ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696Type int32

const (
	ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696Unknown   ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696Type = 0
	ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696Yes       ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696Type = 1
	ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696No        ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696Type = 2
	ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696Agressive ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696Type = 3
)

type ServiceRequestServiceRequestSourceType string

const (
	ServiceRequestServiceRequestSourcePhone         ServiceRequestServiceRequestSourceType = "Phone"
	ServiceRequestServiceRequestSourceEmail         ServiceRequestServiceRequestSourceType = "Email"
	ServiceRequestServiceRequestSourceWebsite       ServiceRequestServiceRequestSourceType = "Website"
	ServiceRequestServiceRequestSourceDropin        ServiceRequestServiceRequestSourceType = "Drop-in"
	ServiceRequestServiceRequestSource2025PoolsList ServiceRequestServiceRequestSourceType = "2025_pools"
)

type ServiceRequestServiceRequestStatusType string

const (
	ServiceRequestServiceRequestStatusAssigned       ServiceRequestServiceRequestStatusType = "Assigned"
	ServiceRequestServiceRequestStatusClosed         ServiceRequestServiceRequestStatusType = "Closed"
	ServiceRequestServiceRequestStatusFieldRectified ServiceRequestServiceRequestStatusType = "FieldRectified"
	ServiceRequestServiceRequestStatusOpen           ServiceRequestServiceRequestStatusType = "Open"
	ServiceRequestServiceRequestStatusRejected       ServiceRequestServiceRequestStatusType = "Rejected"
	ServiceRequestServiceRequestStatusUnverified     ServiceRequestServiceRequestStatusType = "Unverified"
	ServiceRequestServiceRequestStatusAccepted       ServiceRequestServiceRequestStatusType = "Accepted"
)

type ServiceRequestServiceRequestContactPreferencesType string

const (
	ServiceRequestServiceRequestContactPreferencesNone  ServiceRequestServiceRequestContactPreferencesType = "None"
	ServiceRequestServiceRequestContactPreferencesCall  ServiceRequestServiceRequestContactPreferencesType = "Call"
	ServiceRequestServiceRequestContactPreferencesEmail ServiceRequestServiceRequestContactPreferencesType = "Email"
	ServiceRequestServiceRequestContactPreferencesText  ServiceRequestServiceRequestContactPreferencesType = "Text"
)

type ServiceRequest struct {
	ObjectID                uint                                                                           `field:"OBJECTID"`
	Received                time.Time                                                                      `field:"RECDATETIME"`
	Source                  ServiceRequestServiceRequestSourceType                                         `field:"SOURCE"`
	EnteredBy               string                                                                         `field:"ENTRYTECH"`
	Priority                ServiceRequestServiceRequestPriorityType                                       `field:"PRIORITY"`
	Supervisor              ServiceRequestServiceRequestSUPERVISOReba07b90c8854fe680807aa775403b9aType     `field:"SUPERVISOR"`
	AssignedTo              ServiceRequestServiceRequestASSIGNEDTECH71d0d685868f4b7a87e23661a3ee67c5Type   `field:"ASSIGNEDTECH"`
	Status                  ServiceRequestServiceRequestStatusType                                         `field:"STATUS"`
	AnonymousCaller         ServiceRequestNotInUITFType                                                    `field:"CLRANON"`
	CallerName              string                                                                         `field:"CLRFNAME"`
	CallerPhone             string                                                                         `field:"CLRPHONE1"`
	CallerAlternatePhone    string                                                                         `field:"CLRPHONE2"`
	CallerEmail             string                                                                         `field:"CLREMAIL"`
	CallerCompany           string                                                                         `field:"CLRCOMPANY"`
	CallerAddress           string                                                                         `field:"CLRADDR1"`
	CallerAddress2          string                                                                         `field:"CLRADDR2"`
	CallerCity              string                                                                         `field:"CLRCITY"`
	CallerState             ServiceRequestServiceRequestRegionType                                         `field:"CLRSTATE"`
	CallerZip               string                                                                         `field:"CLRZIP"`
	CallerOther             string                                                                         `field:"CLROTHER"`
	CallerContactPreference ServiceRequestServiceRequestContactPreferencesType                             `field:"CLRCONTPREF"`
	RequestCompany          string                                                                         `field:"REQCOMPANY"`
	RequestAddress          string                                                                         `field:"REQADDR1"`
	RequestAddress2         string                                                                         `field:"REQADDR2"`
	RequestCity             string                                                                         `field:"REQCITY"`
	RequestState            ServiceRequestServiceRequestRegionType                                         `field:"REQSTATE"`
	RequestZip              string                                                                         `field:"REQZIP"`
	RequestCrossStreet      string                                                                         `field:"REQCROSSST"`
	RequestSubdivision      string                                                                         `field:"REQSUBDIV"`
	RequestMapGrID          string                                                                         `field:"REQMAPGRID"`
	PermissionToEnter       ServiceRequestNotInUITFType                                                    `field:"REQPERMISSION"`
	RequestTarget           ServiceRequestServiceRequestTargetType                                         `field:"REQTARGET"`
	RequestDescription      string                                                                         `field:"REQDESCR"`
	NotesForFieldTechnician string                                                                         `field:"REQNOTESFORTECH"`
	NotesForCustomer        string                                                                         `field:"REQNOTESFORCUST"`
	RequestFieldNotes       string                                                                         `field:"REQFLDNOTES"`
	RequestProgramActions   string                                                                         `field:"REQPROGRAMACTIONS"`
	Closed                  time.Time                                                                      `field:"DATETIMECLOSED"`
	ClosedBy                string                                                                         `field:"TECHCLOSED"`
	Sr                      int32                                                                          `field:"SR_NUMBER"`
	Reviewed                ServiceRequestNotInUITFType                                                    `field:"REVIEWED"`
	ReviewedBy              string                                                                         `field:"REVIEWEDBY"`
	ReviewedDate            time.Time                                                                      `field:"REVIEWEDDATE"`
	Accepted                ServiceRequestNotInUITFType                                                    `field:"ACCEPTED"`
	AcceptedDate            time.Time                                                                      `field:"ACCEPTEDDATE"`
	RejectedBy              string                                                                         `field:"REJECTEDBY"`
	RejectedDate            time.Time                                                                      `field:"REJECTEDDATE"`
	RejectedReason          ServiceRequestServiceRequestRejectedReasonType                                 `field:"REJECTEDREASON"`
	DueDate                 time.Time                                                                      `field:"DUEDATE"`
	AcceptedBy              string                                                                         `field:"ACCEPTEDBY"`
	Comments                string                                                                         `field:"COMMENTS"`
	EstimatedCompletionDate time.Time                                                                      `field:"ESTCOMPLETEDATE"`
	NextAction              ServiceRequestServiceRequestNextActionType                                     `field:"NEXTACTION"`
	RecordStatus            int16                                                                          `field:"RECORDSTATUS"`
	GlobalID                uuid.UUID                                                                      `field:"GlobalID"`
	CreatedUser             string                                                                         `field:"created_user"`
	CreatedDate             time.Time                                                                      `field:"created_date"`
	LastEditedUser          string                                                                         `field:"last_edited_user"`
	LastEditedDate          time.Time                                                                      `field:"last_edited_date"`
	FirstResponseDate       time.Time                                                                      `field:"FIRSTRESPONSEDATE"`
	ResponseDayCount        int16                                                                          `field:"RESPONSEDAYCOUNT"`
	VerifyCorrectLocation   string                                                                         `field:"ALLOWED"`
	Xvalue                  string                                                                         `field:"XVALUE"`
	Yvalue                  string                                                                         `field:"YVALUE"`
	ValidX                  string                                                                         `field:"VALIDX"`
	ValidY                  string                                                                         `field:"VALIDY"`
	ExternalID              string                                                                         `field:"EXTERNALID"`
	ExternalError           string                                                                         `field:"EXTERNALERROR"`
	PointlocID              uuid.UUID                                                                      `field:"POINTLOCID"`
	Notified                int16                                                                          `field:"NOTIFIED"`
	NotifiedDate            time.Time                                                                      `field:"NOTIFIEDDATE"`
	Scheduled               int16                                                                          `field:"SCHEDULED"`
	ScheduledDate           time.Time                                                                      `field:"SCHEDULEDDATE"`
	Dog                     ServiceRequestServiceRequestDOG2b95ec9712864fcd88f4f0e31113f696Type            `field:"DOG"`
	SchedulePeriod          ServiceRequestServiceRequestscheduleperiod3f40c046afd14abd8bf4389650d29a49Type `field:"schedule_period"`
	ScheduleNotes           string                                                                         `field:"schedule_notes"`
	PreferSpeakingSpanish   ServiceRequestServiceRequestSpanishaaa3dc669f9a45278ecdc9f76db33879Type        `field:"Spanish"`
	CreationDate            time.Time                                                                      `field:"CreationDate"`
	Creator                 string                                                                         `field:"Creator"`
	EditDate                time.Time                                                                      `field:"EditDate"`
	Editor                  string                                                                         `field:"Editor"`
	IssuesReported          ServiceRequestServiceRequestIssuesType                                         `field:"ISSUESREPORTED"`
	Jurisdiction            string                                                                         `field:"JURISDICTION"`
	NotificationTimestamp   string                                                                         `field:"NOTIFICATIONTIMESTAMP"`
	Zone                    string                                                                         `field:"ZONE"`
	Zone2                   string                                                                         `field:"ZONE2"`
}
