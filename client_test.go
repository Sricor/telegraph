package telegraph

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_SetAPI(t *testing.T) {
	type args struct {
		api string
	}
	tests := []struct {
		name string
		c    *client
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SetAPI(tt.args.api)
		})
	}
}

func Test_client_SetToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		c    *client
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SetToken(tt.args.token)
		})
	}
}

func Test_client_Token(t *testing.T) {
	tests := []struct {
		name       string
		c          *client
		wantResult string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.c.Token(); gotResult != tt.wantResult {
				t.Errorf("client.Token() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_client_request(t *testing.T) {
	type args struct {
		path    string
		payload interface{}
	}
	tests := []struct {
		name         string
		c            *client
		args         args
		wantResponse *Response
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.c.request(tt.args.path, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("client.request() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func Test_client_withToken(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.c.withToken(tt.args.v); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.withToken() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_jsonMarshal(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonMarshal(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonMarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonUnmarshal(t *testing.T) {
	type args struct {
		data []byte
		v    any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := jsonUnmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("jsonUnmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
