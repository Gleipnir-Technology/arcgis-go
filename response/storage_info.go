package response

type StorageInfo struct {
	StorageFormat string `json:"storageFormat"`
	PacketSize    int    `json:"packetSize"`
}
