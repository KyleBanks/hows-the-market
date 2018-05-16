VERSION = 0.0.1

RELEASE_PKG = .
RELEASE_PLATFORMS = darwin/386 darwin/amd64 linux/386 linux/amd64 linux/arm 
INSTALL_PKG = $(RELEASE_PKG)

# Remote includes require 'mmake' 
# # github.com/tj/mmake
include github.com/KyleBanks/make/go/install
include github.com/KyleBanks/make/go/sanity
include github.com/KyleBanks/make/go/release
include github.com/KyleBanks/make/go/coverage
include github.com/KyleBanks/make/git/precommit
