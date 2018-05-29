package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// HeartbeatEvent : Emitted every 2 seconds after enabling it by calling SetHeartbeat.
// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#heartbeat
type HeartbeatEvent struct {
	// Toggles between every JSON meassage as an "I am alive" indicator.
	// Required: Yes.
	Pulse bool `mapstructure:"pulse"`
	// Current active profile.
	// Required: No.
	CurrentProfile string `mapstructure:"current-profile"`
	// Current active scene.
	// Required: No.
	CurrentScene string `mapstructure:"current-scene"`
	// Current streaming state.
	// Required: No.
	Streaming bool `mapstructure:"streaming"`
	// Total time (in seconds) since the stream started.
	// Required: No.
	TotalStreamTime int `mapstructure:"total-stream-time"`
	// Total bytes sent since the stream started.
	// Required: No.
	TotalStreamBytes int `mapstructure:"total-stream-bytes"`
	// Total frames streamed since the stream started.
	// Required: No.
	TotalStreamFrames int `mapstructure:"total-stream-frames"`
	// Current recording state.
	// Required: No.
	Recording bool `mapstructure:"recording"`
	// Total time (in seconds) since recording started.
	// Required: No.
	TotalRecordTime int `mapstructure:"total-record-time"`
	// Total bytes recorded since the recording started.
	// Required: No.
	TotalRecordBytes int `mapstructure:"total-record-bytes"`
	// Total frames recorded since the recording started.
	// Required: No.
	TotalRecordFrames int `mapstructure:"total-record-frames"`
	_event            `mapstructure:",squash"`
}

// Type returns the event's update type.
func (e HeartbeatEvent) Type() string { return e.UpdateType }

// StreamTC returns the event's stream timecode.
func (e HeartbeatEvent) StreamTC() string { return e.StreamTimecode }

// RecTC returns the event's recording timecode.
func (e HeartbeatEvent) RecTC() string { return e.RecTimecode }
