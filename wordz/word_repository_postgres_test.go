package wordz

import "testing"

func TestWordRepositoryPostgres_FetchWordByNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy Path",
			args: args{
				n: 1,
			},
			want: "SHINE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := WordRepositoryPostgres{}
			if got := r.FetchWordByNumber(tt.args.n); got != tt.want {
				t.Errorf("FetchWordByNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
