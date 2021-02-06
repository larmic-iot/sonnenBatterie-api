package crypto

import "testing"

func TestEncrypt(t *testing.T) {
	type args struct {
		data string
		salt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Encrypt('demo', 'test-salt')", args{data: "demo", salt: "test-salt"}, "6b8b89b180cf95ecdba3e0556562a9d6d2f73b054c4b0d388c6ebf7ea61b8a5f53fb93f021209b796a977261b7aae4d3d7c67d3cd963f58b82170930031f81ce"},
		{"Encrypt('', 'test-salt')", args{data: "", salt: "test-salt"}, "6b36b3512166a7f4d69c458c0dcc34368c586dc02331855ecfe10b51f03bf224fbe79d451ee6d54b5ed2a51fc28a8825c74a9c382854dce685e3ec3102e7200a"},
		{"Encrypt('test-data', '')", args{data: "test-date", salt: ""}, "cddbab1385dba100e0263979b6c84a1b0870a7d122030433aa0fb87e4e618374a7a2f6f991ecd969e3a471553605903924e6e2e597530f2c10a8b7dcca6af3b3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encrypt(tt.args.data, tt.args.salt); got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
