package cryptoamount

import "testing"

func TestCryptoAmount_ToString(t *testing.T) {
	type args struct {
		decimals int
	}
	tests := []struct {
		name string
		a    CryptoAmount
		args args
		want string
	}{
		{
			name: "Test with 6 decimals and 0.100000",
			a:    100000,
			args: args{decimals: 6},
			want: "0.100000",
		},
		{
			name: "Test with 18 decimals and 0.123456",
			a:    123456000000000000,
			args: args{decimals: 18},
			want: "0.123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.ToString(tt.args.decimals); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
