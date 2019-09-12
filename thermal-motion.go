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

const ThermalMotionKey = "thermal-motion"

type ThermalMotion struct {
	DynamicThreshold bool   `mapstructure:"min-secs"`
	TempThresh       uint16 `mapstructure:"temp-thresh"`
	DeltaThresh      uint16 `mapstructure:"delta-thresh"`
	CountThresh      int    `mapstructure:"count-thresh"`
	FrameCompareGap  int    `mapstructure:"frame-compare-gap"`
	UseOneDiffOnly   bool   `mapstructure:"use-one-diff-only"`
	TriggerFrames    int    `mapstructure:"trigger-frames"`
	WarmerOnly       bool   `mapstructure:"warmer-only"`
	EdgePixels       int    `mapstructure:"edge-pixels"`
	Verbose          bool   `mapstructure:"verbose"`
}

func DefaultThermalMotion() ThermalMotion {
	return ThermalMotion{
		DynamicThreshold: true,
		TempThresh:       2900,
		DeltaThresh:      50,
		CountThresh:      3,
		FrameCompareGap:  45,
		Verbose:          false,
		TriggerFrames:    2,
		UseOneDiffOnly:   true,
		WarmerOnly:       true,
		EdgePixels:       1,
	}
}
