// BSD 3-Clause License
//
// Copyright (c) 2019, Guillaume Ballet
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of the copyright holder nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package pcsc

const (
	SCardSuccess                   = 0x00000000 /* No error was encountered. */
	SCardProtocolUndefined = 0x0000                 /* protocol not set */
	SCardProtocolUnSet     = SCardProtocolUndefined /* backward compat */
	SCardProtocolT0        = 0x0001                 /* T=0 active protocol. */
	SCardProtocolT1        = 0x0002                 /* T=1 active protocol. */
	SCardProtocolRaw       = 0x0004                 /* Raw active protocol. */
	SCardProtocolT15       = 0x0008                 /* T=15 protocol. */

	SCardProtocolAny = (SCardProtocolT0 | SCardProtocolT1) /* IFD determines prot. */

	SCardShareExclusive = 0x0001 /* Exclusive mode only */
	SCardShareShared    = 0x0002 /* Shared mode only */
	SCardShareDirect    = 0x0003 /* Raw mode only */

	SCardLeaveCard   = 0x0000 /* Do nothing on close */
	SCardResetCard   = 0x0001 /* Reset on close */
	SCardUnpowerCard = 0x0002 /* Power down on close */
	SCardEjectCard   = 0x0003 /* Eject on close */

	SCardUnknown    = 0x0001 /* Unknown state */
	SCardAbsent     = 0x0002 /* Card is absent */
	SCardPresent    = 0x0004 /* Card is present */
	SCardSwallowed  = 0x0008 /* Card not powered */
	SCardPowever    = 0x0010 /* Card is powered */
	SCardNegotiable = 0x0020 /* Ready for PTS */
	SCardSpecific   = 0x0040 /* PTS has been set */

	PCSCDSockName = "/run/pcscd/pcscd.comm"
)

// List of commands to send to the daemon
const (
	_                                   = iota
	SCardEstablishContext               /* used by SCardEstablishContext() */
	SCardReleaseContext                 /* used by SCardReleaseContext() */
	SCardListReaders                    /* used by SCardListReaders() */
	SCardConnect                        /* used by SCardConnect() */
	SCardReConnect                      /* used by SCardReconnect() */
	SCardDisConnect                     /* used by SCardDisconnect() */
	SCardBeginTransaction               /* used by SCardBeginTransaction() */
	SCardEndTransaction                 /* used by SCardEndTransaction() */
	SCardTransmit                       /* used by SCardTransmit() */
	SCardControl                        /* used by SCardControl() */
	SCardStatus                         /* used by SCardStatus() */
	SCardGetStatusChange                /* not used */
	SCardCancel                         /* used by SCardCancel() */
	SCardCancelTransaction              /* not used */
	SCardGetAttrib                      /* used by SCardGetAttrib() */
	SCardSetAttrib                      /* used by SCardSetAttrib() */
	CommandVersion                      /* get the client/server protocol version */
	CommandGetReaderState               /* get the readers state */
	CommandWaitReaderStateChange        /* wait for a reader state change */
	CommandStopWaitingReaderStateChange /* stop waiting for a reader state change */
)

// Protocol information
const (
	ProtocolVersionMajor = 4 /* IPC major */
	ProtocolVersionMinor = 3 /* IPC minor */
)
