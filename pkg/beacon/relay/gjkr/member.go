package gjkr

import (
	"math/big"

	"github.com/keep-network/keep-core/pkg/beacon/relay/pedersen"
)

type memberCore struct {
	// ID of this group member.
	ID int
	// Group to which this member belongs.
	group *Group
	// DKG Protocol configuration parameters.
	protocolConfig *DKG
}

// CommittingMember represents one member in a threshold key sharing group, after
// it has a full list of `memberIDs` that belong to its threshold group. A
// member in this state has two maps of member shares for each member of the
// group.
type CommittingMember struct {
	*memberCore

	// Pedersen VSS scheme used to calculate commitments.
	vss *pedersen.VSS
	// Polynomial `a` coefficients generated by the member. Polynomial is of
	// degree `dishonestThreshold`, so the number of coefficients equals
	// `dishonestThreshold + 1`
	//
	// This is a private value and should not be exposed.
	secretCoefficients []*big.Int
	// Shares calculated by the current member for themself. They are defined as
	// `s_ii` and `t_ii` respectively across the protocol specification.
	//
	// These are private values and should not be exposed.
	selfSecretShareS, selfSecretShareT *big.Int
	// Shares calculated for the current member by peer group members.
	//
	// receivedSharesS are defined as `s_ji` and receivedSharesT are
	// defined as `t_ji` across the protocol specification.
	receivedSharesS, receivedSharesT map[int]*big.Int
	// Commitments to coefficients received from peer group members.
	receivedCommitments map[int][]*big.Int
}

// SharingMember represents one member in a threshold key sharing group.
type SharingMember struct {
	*CommittingMember

	shareS, shareT *big.Int
	// Public values of each polynomial `a` coefficient defined in secretCoefficients
	// field. It is denoted as `A_ik` in protocol specification.
	publicCoefficients []*big.Int
}
