package packets

type ServerGamePlayerWinCount struct {
	Packet
	UserId int `json:"uid"`
	Wins   int `json:"w"`
}

func NewServerGamePlayerWinCount(userId int, wins int) *ServerGamePlayerWinCount {
	return &ServerGamePlayerWinCount{
		Packet: Packet{Id: PacketIdServerGamePlayerWinCount},
		UserId: userId,
		Wins:   wins,
	}
}
