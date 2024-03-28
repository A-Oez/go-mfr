package json_model

type ServiceRequests struct {
	OdataMetadata string  `json:"odata.metadata"`
	Value         []Value `json:"value"`
}

type Value struct {
	Appointments                       []Appointment `json:"Appointments"`
	Steps                              []Step        `json:"Steps"`
	Id                                 string        `json:"Id"`
	Name                               string        `json:"Name"`
	ExternalId                         string        `json:"ExternalId"`
	InvoiceId                          string        `json:"InvoiceId"`
	ClosedAt                           string        `json:"ClosedAt"`
	ReleasedAt                         string        `json:"ReleasedAt"`
	WorkDoneAt                         string        `json:"WorkDoneAt"`
	TargetTimeInMinutes                string        `json:"TargetTimeInMinutes"`
	DateModified                       string        `json:"DateModified"`
	DateOfCreation                     string        `json:"DateOfCreation"`
	DueDateRangeStart                  string        `json:"DueDateRangeStart"`
	DueDateRangeEnd                    string        `json:"DueDateRangeEnd"`
	PortalLink                         string        `json:"PortalLink"`
	CostCenterId                       string        `json:"CostCenterId"`
	Description                        string        `json:"Description"`
	State                              string        `json:"State"`
	CustomValues                       []string      `json:"CustomValues"`
	CurrentOwnerId                     string        `json:"CurrentOwnerId"`
	CustomerId                         string        `json:"CustomerId"`
	ParentServiceRequestId             string        `json:"ParentServiceRequestId"`
	Location                           string        `json:"Location"`
	Version                            int           `json:"Version"`
	IsTemplate                         bool          `json:"IsTemplate"`
	IsTemplateMobile                   bool          `json:"IsTemplateMobile"`
	CreateFromServiceRequestTemplateId string        `json:"CreateFromServiceRequestTemplateId"`
	Type                               string        `json:"Type"`
}

type Appointment struct {
	Id                  string   `json:"Id"`
	Version             int      `json:"Version"`
	State               string   `json:"State"`
	Type                string   `json:"Type"`
	CreatedAt           string   `json:"CreatedAt"`
	EndDateTime         string   `json:"EndDateTime"`
	StartDateTime       string   `json:"StartDateTime"`
	DrivingDistanceFrom int      `json:"DrivingDistanceFrom"`
	DrivingDistanceTo   int      `json:"DrivingDistanceTo,omitempty"`
	WasReadOnClientSide bool     `json:"WasReadOnClientSide"`
	ContactIds          []string `json:"ContactIds"`
	ServiceRequestId    string   `json:"ServiceRequestId"`
	ContactId           string   `json:"ContactId"`
	Note                string   `json:"Note,omitempty"`
	ExternalId          string   `json:"ExternalId,omitempty"`
}

type Step struct {
	Id                  string `json:"Id"`
	MobileId            string `json:"MobileId"`
	Version             int    `json:"Version"`
	Name                string `json:"Name"`
	IsDone              bool   `json:"IsDone"`
	HasError            bool   `json:"HasError"`
	TrackingId          string `json:"TrackingId"`
	Type                string `json:"Type"`
	SortOrder           int    `json:"SortOrder"`
	Data                string `json:"Data"`
	DateModifiedOffline string `json:"DateModifiedOffline"`
	ServiceRequestId    string `json:"ServiceRequestId"`
	Description         string `json:"Description"`
	Comment             string `json:"Comment"`
	InternalComment     string `json:"InternalComment"`
	ServiceObjectId     string `json:"ServiceObjectId"`
	StepListTemplateId  string `json:"StepListTemplateId"`
	ParentId            string `json:"ParentId"`
}

type StepData struct {
	Fields []StepDataField `json:"fields"`
	Custom interface{}     `json:"custom"`
}

type StepDataField struct {
	IsRequired bool   `json:"isRequired"`
	Name       string `json:"name"`
	Result     string `json:"result"`
	Type       string `json:"type"`
	Unit       string `json:"unit"`
}
