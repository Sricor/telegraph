package telegraph

import (
	"reflect"
	"testing"
)

func Test_client_GetPage(t *testing.T) {
	type args struct {
		p *GetPageRequest
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult *Page
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.c.GetPage(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.GetPage() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_client_GetViews(t *testing.T) {
	type args struct {
		p *GetViewsRequest
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult *PageViews
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.c.GetViews(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetViews() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.GetViews() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_client_CreatePage(t *testing.T) {
	type args struct {
		p *CreatePageRequest
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult *Page
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.c.CreatePage(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.CreatePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.CreatePage() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_client_EditPage(t *testing.T) {
	type args struct {
		p *EditPageRequest
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult *Page
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.c.EditPage(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.EditPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.EditPage() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_client_GetPageList(t *testing.T) {
	type args struct {
		p *GetPageListRequest
	}
	tests := []struct {
		name       string
		c          *client
		args       args
		wantResult *PageList
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.c.GetPageList(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("client.GetPageList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("client.GetPageList() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestNodeElement_AddChildren(t *testing.T) {
	type args struct {
		t *NodeElement
	}
	tests := []struct {
		name string
		n    *NodeElement
		args args
		want *NodeElement
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.AddChildren(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeElement.AddChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeElement_AddText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		n    *NodeElement
		args args
		want *NodeElement
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.AddText(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeElement.AddText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeElement_AddImage(t *testing.T) {
	type args struct {
		src         string
		description string
	}
	tests := []struct {
		name string
		n    *NodeElement
		args args
		want *NodeElement
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.AddImage(tt.args.src, tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeElement.AddImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
