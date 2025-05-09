// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package webrtc

// SCTPTransportState indicates the state of the SCTP transport.
type SCTPTransportState int

const (
	// SCTPTransportStateUnknown is the enum's zero-value.
	SCTPTransportStateUnknown SCTPTransportState = iota

	// SCTPTransportStateConnecting indicates the SCTPTransport is in the
	// process of negotiating an association. This is the initial state of the
	// SCTPTransportState when an SCTPTransport is created.
	SCTPTransportStateConnecting

	// SCTPTransportStateConnected indicates the negotiation of an
	// association is completed.
	SCTPTransportStateConnected

	// SCTPTransportStateClosed indicates a SHUTDOWN or ABORT chunk is
	// received or when the SCTP association has been closed intentionally,
	// such as by closing the peer connection or applying a remote description
	// that rejects data or changes the SCTP port.
	SCTPTransportStateClosed
)

// This is done this way because of a linter.
const (
	sctpTransportStateConnectingStr = "connecting"
	sctpTransportStateConnectedStr  = "connected"
	sctpTransportStateClosedStr     = "closed"
)

func newSCTPTransportState(raw string) SCTPTransportState {
	switch raw {
	case sctpTransportStateConnectingStr:
		return SCTPTransportStateConnecting
	case sctpTransportStateConnectedStr:
		return SCTPTransportStateConnected
	case sctpTransportStateClosedStr:
		return SCTPTransportStateClosed
	default:
		return SCTPTransportStateUnknown
	}
}

func (s SCTPTransportState) String() string {
	switch s {
	case SCTPTransportStateConnecting:
		return sctpTransportStateConnectingStr
	case SCTPTransportStateConnected:
		return sctpTransportStateConnectedStr
	case SCTPTransportStateClosed:
		return sctpTransportStateClosedStr
	default:
		return ErrUnknownType.Error()
	}
}
