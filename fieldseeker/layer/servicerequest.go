package layer

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ServiceRequest struct {
	ObjectID                uint      `field:"OBJECTID"`
	Received                time.Time `field:"RECDATETIME"`
	Source                  string    `field:"SOURCE"`
	EnteredBy               string    `field:"ENTRYTECH"`
	Priority                string    `field:"PRIORITY"`
	Supervisor              string    `field:"SUPERVISOR"`
	AssignedTo              string    `field:"ASSIGNEDTECH"`
	Status                  string    `field:"STATUS"`
	AnonymousCaller         int16     `field:"CLRANON"`
	CallerName              string    `field:"CLRFNAME"`
	CallerPhone             string    `field:"CLRPHONE1"`
	CallerAlternatePhone    string    `field:"CLRPHONE2"`
	CallerEmail             string    `field:"CLREMAIL"`
	CallerCompany           string    `field:"CLRCOMPANY"`
	CallerAddress           string    `field:"CLRADDR1"`
	CallerAddress2          string    `field:"CLRADDR2"`
	CallerCity              string    `field:"CLRCITY"`
	CallerState             string    `field:"CLRSTATE"`
	CallerZip               string    `field:"CLRZIP"`
	CallerOther             string    `field:"CLROTHER"`
	CallerContactPreference string    `field:"CLRCONTPREF"`
	RequestCompany          string    `field:"REQCOMPANY"`
	RequestAddress          string    `field:"REQADDR1"`
	RequestAddress2         string    `field:"REQADDR2"`
	RequestCity             string    `field:"REQCITY"`
	RequestState            string    `field:"REQSTATE"`
	RequestZip              string    `field:"REQZIP"`
	RequestCrossStreet      string    `field:"REQCROSSST"`
	RequestSubdivision      string    `field:"REQSUBDIV"`
	RequestMapGrID          string    `field:"REQMAPGRID"`
	PermissionToEnter       int16     `field:"REQPERMISSION"`
	RequestTarget           string    `field:"REQTARGET"`
	RequestDescription      string    `field:"REQDESCR"`
	NotesForFieldTechnician string    `field:"REQNOTESFORTECH"`
	NotesForCustomer        string    `field:"REQNOTESFORCUST"`
	RequestFieldNotes       string    `field:"REQFLDNOTES"`
	RequestProgramActions   string    `field:"REQPROGRAMACTIONS"`
	Closed                  time.Time `field:"DATETIMECLOSED"`
	ClosedBy                string    `field:"TECHCLOSED"`
	Sr                      int32     `field:"SR_NUMBER"`
	Reviewed                int16     `field:"REVIEWED"`
	ReviewedBy              string    `field:"REVIEWEDBY"`
	ReviewedDate            time.Time `field:"REVIEWEDDATE"`
	Accepted                int16     `field:"ACCEPTED"`
	AcceptedDate            time.Time `field:"ACCEPTEDDATE"`
	RejectedBy              string    `field:"REJECTEDBY"`
	RejectedDate            time.Time `field:"REJECTEDDATE"`
	RejectedReason          string    `field:"REJECTEDREASON"`
	DueDate                 time.Time `field:"DUEDATE"`
	AcceptedBy              string    `field:"ACCEPTEDBY"`
	Comments                string    `field:"COMMENTS"`
	EstimatedCompletionDate time.Time `field:"ESTCOMPLETEDATE"`
	NextAction              string    `field:"NEXTACTION"`
	RecordStatus            int16     `field:"RECORDSTATUS"`
	GlobalID                uuid.UUID `field:"GlobalID"`
	CreatedUser             string    `field:"created_user"`
	CreatedDate             time.Time `field:"created_date"`
	LastEditedUser          string    `field:"last_edited_user"`
	LastEditedDate          time.Time `field:"last_edited_date"`
	FirstResponseDate       time.Time `field:"FIRSTRESPONSEDATE"`
	ResponseDayCount        int16     `field:"RESPONSEDAYCOUNT"`
	VerifyCorrectLocation   string    `field:"ALLOWED"`
	Xvalue                  string    `field:"XVALUE"`
	Yvalue                  string    `field:"YVALUE"`
	ValidX                  string    `field:"VALIDX"`
	ValidY                  string    `field:"VALIDY"`
	ExternalID              string    `field:"EXTERNALID"`
	ExternalError           string    `field:"EXTERNALERROR"`
	PointlocID              uuid.UUID `field:"POINTLOCID"`
	Notified                int16     `field:"NOTIFIED"`
	NotifiedDate            time.Time `field:"NOTIFIEDDATE"`
	Scheduled               int16     `field:"SCHEDULED"`
	ScheduledDate           time.Time `field:"SCHEDULEDDATE"`
	Dog                     int32     `field:"DOG"`
	SchedulePeriod          string    `field:"schedule_period"`
	ScheduleNotes           string    `field:"schedule_notes"`
	PreferSpeakingSpanish   int32     `field:"Spanish"`
	CreationDate            time.Time `field:"CreationDate"`
	Creator                 string    `field:"Creator"`
	EditDate                time.Time `field:"EditDate"`
	Editor                  string    `field:"Editor"`
	IssuesReported          string    `field:"ISSUESREPORTED"`
	Jurisdiction            string    `field:"JURISDICTION"`
	NotificationTimestamp   string    `field:"NOTIFICATIONTIMESTAMP"`
	Zone                    string    `field:"ZONE"`
	Zone2                   string    `field:"ZONE2"`
	Geometry                json.RawMessage
}

func (x *ServiceRequest) GetGeometry() json.RawMessage  { return x.Geometry }
func (x *ServiceRequest) SetGeometry(m json.RawMessage) { x.Geometry = m }
