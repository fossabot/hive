package hive

import "testing"

func TestGo_GenerateServer(t *testing.T) {
	tests := []struct {
		name    string
		Setup *LanguageSetup
		bee *Bee
		wantErr bool
	}{
		{
			name:    "sub0",
			Setup:   &LanguageSetup{
				Active:  true,
				Repo:    "github.com/benka-me/sub0",
				Files:   []string{
					"sub0.proto",
					"rpc-sub0.proto",
				},
				PkgName: "sub0",
			},
			bee:     &Bee{
				DevLang: DevLang(0),
			},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Go{
				Setup: tt.Setup,
			}
			if err := g.GenerateServer(tt.bee); (err != nil) != tt.wantErr {
				t.Errorf("ClientsFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}