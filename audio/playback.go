package audio

import "github.com/CyCoreSystems/ari"

// A Playback is a lifecycle managed audio object
type Playback struct {
	startCh chan struct{}
	stopCh  chan struct{}
	quitCh  chan struct{}

	handle *ari.PlaybackHandle
	err    error
}

// Handle returns the ARI reference to the playback object
func (p *Playback) Handle() *ari.PlaybackHandle {
	return p.handle
}

// StartCh returns the channel that is closed when the playback has started
func (p *Playback) StartCh() <-chan struct{} {
	return p.startCh
}

// StopCh returns the channel that is closed when the playback has stopped
func (p *Playback) StopCh() <-chan struct{} {
	return p.stopCh
}

// Err returns any accumulated errors during playback
func (p *Playback) Err() error {
	return p.err
}

// Cancel cancels the playback
func (p *Playback) Cancel() {
	if p.quitCh != nil {
		close(p.quitCh)
	}
	p.quitCh = nil
}
