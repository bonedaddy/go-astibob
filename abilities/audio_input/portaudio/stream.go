package portaudio

import (
	"fmt"

	"github.com/asticode/go-astikit"
	"github.com/gordonklaus/portaudio"
)

type Stream struct {
	b []int32
	l astikit.SeverityLogger
	o StreamOptions
	s *portaudio.Stream
}

type StreamOptions struct {
	BitDepth          int     `toml:"bit_depth"`
	BufferLength      int     `toml:"buffer_length"`
	MaxSilenceLevel   float64 `toml:"max_silence_level"`
	NumInputChannels  int     `toml:"num_input_channels"`
	NumOutputChannels int     `toml:"num_output_channels"`
	SampleRate        int     `toml:"sample_rate"`
}

func (p *PortAudio) NewDefaultStream(o StreamOptions) (s *Stream, err error) {
	// Create stream
	s = &Stream{
		b: make([]int32, o.BufferLength),
		l: p.l,
		o: o,
	}

	// Log
	p.l.Debugf("portaudio: opening default stream %p", s)

	// Open default stream
	if s.s, err = portaudio.OpenDefaultStream(s.o.NumInputChannels, s.o.NumOutputChannels, float64(s.o.SampleRate), len(s.b), s.b); err != nil {
		err = fmt.Errorf("portaudio: opening default stream %p failed: %w", s, err)
		return
	}
	return
}

func (s *Stream) BitDepth() int { return s.o.BitDepth }

func (s *Stream) MaxSilenceLevel() float64 { return s.o.MaxSilenceLevel }

func (s *Stream) NumChannels() int { return s.o.NumInputChannels }

func (s *Stream) SampleRate() int { return s.o.SampleRate }

func (s *Stream) Close() (err error) {
	// Log
	s.l.Debugf("portaudio: closing stream %p", s)

	// Close
	if err = s.s.Close(); err != nil {
		err = fmt.Errorf("portaudio: closing stream %p failed: %w", s, err)
		return
	}
	return
}

func (s *Stream) Start() (err error) {
	// Log
	s.l.Debugf("portaudio: starting stream %p", s)

	// Start
	if err = s.s.Start(); err != nil {
		err = fmt.Errorf("portaudio: starting stream %p failed: %w", s, err)
		return
	}
	return
}

func (s *Stream) Stop() (err error) {
	// Log
	s.l.Debugf("portaudio: stopping stream %p", s)

	// Stop
	if err = s.s.Stop(); err != nil {
		err = fmt.Errorf("portaudio: stopping stream %p failed: %w", s, err)
		return
	}
	return
}

func (s *Stream) Read() (rs []int, err error) {
	// Read
	if err = s.s.Read(); err != nil {
		err = fmt.Errorf("portaudio: reading from stream %p failed: %w", s, err)
		return
	}

	// Clone buffer
	for _, v := range s.b {
		rs = append(rs, int(v))
	}
	return
}
