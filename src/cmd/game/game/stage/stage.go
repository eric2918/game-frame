package stage

import (
	"frame/pb"

	"github.com/spf13/cast"
)

type Stage struct {
	Basic   map[int64]*pb.StageMission
	Mission map[string]*pb.UnlockMission
}

func New(mission map[string]*pb.UnlockMission, basic map[int64]*pb.StageMission) *Stage {
	return &Stage{Basic: basic, Mission: mission}
}

// Unlock 解锁关卡
func (s *Stage) Unlock(missionIds []string) {
	if s.Mission == nil {
		s.Mission = make(map[string]*pb.UnlockMission)
	}
	for _, missionId := range missionIds {
		if _, ok := s.Mission[missionId]; ok {
			continue
		}
		s.Mission[missionId] = &pb.UnlockMission{
			MissionID: missionId,
			ChapterID: s.Basic[cast.ToInt64(missionId)].ChapterID,
			StarLevel: 0,
		}
	}
}
