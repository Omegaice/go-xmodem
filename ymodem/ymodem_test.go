// portenta-c33-fwuploader-plugin
// Copyright (c) 2023 Arduino LLC.  All right reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package ymodem

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYModem_CRC16(t *testing.T) {
	data := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	require.Equal(t, uint16(39210), CRC16(data))
}

func TestYModem_CRC16Constant(t *testing.T) {
	data := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	require.Equal(t, uint16(43803), CRC16Constant(data, 13))
}
