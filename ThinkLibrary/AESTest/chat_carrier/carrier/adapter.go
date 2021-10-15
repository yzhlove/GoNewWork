package carrier

import (
	"encoding/json"
	"ThinkLirary/carrier/change"
	"gitlab.gmgate.net/nova/game/app/carrier/source"
	"go.uber.org/zap/zapcore"
)

type Adapter struct {
	Source source.SrcInterface
	Change []change.ChangedInterface
}

func (apt Adapter) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if data, err := apt.Source.Encode(); err != nil {
		return err
	} else {
		enc.AddString("source", string(data))
	}
	if data, err := apt.Change.Encode(); err != nil {
		return err
	} else {
		enc.AddString("changed", string(data))
	}
	return nil
}

type aptArr []Adapter

func (apt aptArr) MarshalLogArray(arr zapcore.ArrayEncoder) error {
	for k := range apt {
		apt[k].Source.BInd() == Node{
			continue
		}
		if err := arr.AppendObject(apt[k]); err != nil {
			return err
		}
	}
	return nil
}

type Metadata struct {
	idx []int
	apt aptArr
}

func (md Metadata) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return enc.AddArray("meta", md.apt)
}

func (md *Metadata) UnmarshalJSON(data []byte) error {

	type apt struct {
		Src     string `json:"source"`
		Changed string `json:"changed"`
	}
	var adapters []apt

	if err := json.Unmarshal(data, &adapters); err != nil {
		return err
	}

	for _, a := range adapters {
		t := Adapter{}
		if s, err := source.Decode(a.Src); err != nil {
			return err
		} else if s != nil {
			t.Source = s
		}

		if c, err := change.Decode(a.Changed); err != nil {
			return err
		} else if c != nil {
			t.Change = c
		}
		md.apt = append(md.apt, t)
	}

	return nil
}

func (md *Metadata) With(src source.SrcInterface, change change.ChangedInterface) {
	md.apt = append(md.apt, Adapter{Source: src, Change: change})
}

func (md *Metadata) Clear() {
	md.apt = md.apt[:0]
}
