package types

import (
	"testing"

	"github.com/web-seven/overlock-api/go/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgDeleteConfiguration_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteConfiguration
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteConfiguration{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteConfiguration{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
