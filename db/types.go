package db

import (
	"crypto/sha256"
)

type Hash = [sha256.Size]byte
