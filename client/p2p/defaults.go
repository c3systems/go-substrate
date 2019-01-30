package p2p

// TODO: merge with defaults sub-package?

import "fmt"

// DefaultClientID ...
var DefaultClientID = "go-substrate/0.0.0"

// DefaultAddress ...
var DefaultAddress = "127.0.0.1"

// DefaultMaxPeers ...
var DefaultMaxPeers = 25

// DefaultPort ...
var DefaultPort = 31333

// DefaultProtocolPing ...
var DefaultProtocolPing = "/ipfs/ping/1.0.0"

// DefaultMaxRequestBlocks ...
var DefaultMaxRequestBlocks = 64

// DefaultProtocolBase ...
// TODO: change...
var DefaultProtocolBase = "/substrate"

// DefaultProtocolType ...
var DefaultProtocolType = "/sup"

// DefaultProtocolVersion ...
var DefaultProtocolVersion = 1

// DefaultRole ...
var DefaultRole = "full"

// DefaultMaxQueuedBlocks ...
var DefaultMaxQueuedBlocks = DefaultMaxRequestBlocks * 8

// DefaultMinIdleBlocks ...
var DefaultMinIdleBlocks = 16

// DefaultProtocolDot ...
var DefaultProtocolDot = fmt.Sprintf("%s%s/%d", DefaultProtocolBase, DefaultProtocolType, DefaultProtocolVersion)
