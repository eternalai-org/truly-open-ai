package lighthouse

import (
	"reflect"
	"testing"
)

func TestZipAndUploadFileInMultiplePartsToLightHouseByUrl(t *testing.T) {
	type args struct {
		filename string
		fileDir  string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{filename: "test1", fileDir: ""},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				got, err := ZipAndUploadFileInMultiplePartsToLightHouseByUrl(
					"http://storage.googleapis.com/eternal-ai/ai-models/eternalai-774675271/1737456588-offshorefinancialfreedom33topsecretpublishing2016topsecret.pdf",
					"/tmp/data",
					"da69db7d.010ffac15c0f4081a938b9446f599e14")
				if (err != nil) != tt.wantErr {
					t.Errorf("getListZipFile() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("getListZipFile() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestDownloadHFModelFromLightHouse(t *testing.T) {
	type args struct {
		hash  string
		hfDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{hash: "bafkreicraauqqrju5cw4tvqaqimwxxn6fckxkjs3sohpkeotxcf423ikty", hfDir: "/tmp/download"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DownloadFileFromLightHouse(tt.args.hash, tt.args.hfDir); (err != nil) != tt.wantErr {
				t.Errorf("DownloadHFModelFromLightHouse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
