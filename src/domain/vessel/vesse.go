package vessel

type Vessel struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	FlagID uint64 `json:"flag_id"`
}
