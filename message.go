package gcm

// Message is used by the application server to send a message to
// the GCM server. See the documentation for GCM Architectural
// Overview for more information:
// http://developer.android.com/google/gcm/gcm.html#send-msg
type Message struct {
	RegistrationIDs       []string               `json:"registration_ids"`
	CollapseKey           string                 `json:"collapse_key,omitempty"`
	Data                  map[string]interface{} `json:"data,omitempty"`
	DelayWhileIdle        bool                   `json:"delay_while_idle,omitempty"`
	TimeToLive            int                    `json:"time_to_live,omitempty"`
	RestrictedPackageName string                 `json:"restricted_package_name,omitempty"`
	DryRun                bool                   `json:"dry_run,omitempty"`
}

// NewMessage returns a new Message with the specified payload
// and registration IDs.
func NewMessage(data map[string]interface{}, regIDs ...string) *Message {
	return &Message{RegistrationIDs: regIDs, Data: data}
}

func (message *Message) SetDelayWhileIdle(delay_while_idle bool) {
	message.DelayWhileIdle = delay_while_idle
}

func (message *Message) SetTimeToLive(time_to_live int) {
	message.TimeToLive = time_to_live
}

func (message *Message) SetRestrictedPackageName(restricted_package_name string) {
	message.RestrictedPackageName = restricted_package_name
}

func (message *Message) SetCollapseKey(collapse_key string) {
	message.CollapseKey = collapse_key
}

func (message *Message) SetDryRun(dry_run bool) {
	message.DryRun = dry_run
}
