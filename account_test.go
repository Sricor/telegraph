package telegraph

import (
	"reflect"
	"testing"
)

func Test_client_CreateAccount(t *testing.T) {
	type args struct {
		p *CreateAccountParameters
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult *Account
		wantErr    bool
	}{
		// {
		// 	name: "Test1",
		// 	c:    NewClient(),
		// 	args: args{&CreateAccountParameters{ShortName: "Ruz"}},
		// 	wantResult: &Account{
		// 		ShortName:   "Ruz",
		// 		AccessToken: "",
		// 		AuthURL:     "",
		// 	},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.c.CreateAccount(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.CreateAccount() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_client_EditAccountInfo(t *testing.T) {
	type args struct {
		p *EditAccountInfoParameters
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult *Account
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.c.EditAccountInfo(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.EditAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.EditAccountInfo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_client_GetAccountInfo(t *testing.T) {
	type args struct {
		p *GetAccountInfoParameters
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult *Account
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.c.GetAccountInfo(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetAccountInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.GetAccountInfo() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_client_RevokeAccessToken(t *testing.T) {
	type args struct {
		p *RevokeAccessTokenParameters
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult *Account
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.c.RevokeAccessToken(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.RevokeAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.RevokeAccessToken() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
