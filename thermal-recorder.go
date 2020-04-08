// go-config - Library for reading cacophony config files.
// Copyright (C) 2018, The Cacophony Project
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package config

const ThermalRecorderKey = "thermal-recorder"

func init() {
	allSections[ThermalRecorderKey] = section{
		key:         ThermalRecorderKey,
		mapToStruct: thermalRecorderMapToStruct,
		validate:    noValidateFunc,
	}
}

type ThermalRecorder struct {
	Active         bool
	OutputDir      string `mapstructure:"output-dir"`
	MinDiskSpaceMB uint64 `mapstructure:"min-disk-space-mb"`
	MinSecs        int    `mapstructure:"min-secs"`
	MaxSecs        int    `mapstructure:"max-secs"`
	PreviewSecs    int    `mapstructure:"preview-secs"`
	LogRate        int    `mapstructure:"log-rate"`
}

func DefaultThermalRecorder() ThermalRecorder {
	return ThermalRecorder{
		Active:         false,
		MaxSecs:        600,
		MinSecs:        10,
		PreviewSecs:    3,
		MinDiskSpaceMB: 200,
		OutputDir:      "/var/spool/cptv",
		LogRate:        0,
	}
}

func thermalRecorderMapToStruct(m map[string]interface{}) (interface{}, error) {
	var s ThermalRecorder
	if err := decodeStructFromMap(&s, m, nil); err != nil {
		return nil, err
	}
	return s, nil
}
