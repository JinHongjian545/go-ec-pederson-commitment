package go_ec_pederson_commitment

import (
	"fmt"
	"testing"
)

func TestParamsGenToString(t *testing.T) {
	GString, HString := ParamsGenToString()
	fmt.Println(GString)
	fmt.Println(HString)
}

func TestRandomGenToNumberString(t *testing.T) {
	rString := RandomGenToNumberString()
	fmt.Println(rString)
}

func TestCommitToString(t *testing.T) {
	type args struct {
		GString string
		HString string
		rString string
		secret  []byte
	}
	tests := []struct {
		name             string
		args             args
		wantCommitString string
		wantErr          bool
	}{
		{
			name: "commit",
			args: args{
				"364732713438655479616d68564b713549766652396c744a7872344e7276545f2d5f5f736f514f79496d38",
				"49756256713151674b63304144486738416a50376c48756b4a7268566d6b5551574b6a626e6e456a62796b",
				"398891056723960618120827012073372943324213998153307115797995179191814376821",
				[]byte("this is the secret message to commit"),
			},
			wantCommitString: "",
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCommitString, err := CommitToString(tt.args.GString, tt.args.HString, tt.args.rString, tt.args.secret)
			fmt.Println("commitment: " + gotCommitString)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommitToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
/*			if gotCommitString != tt.wantCommitString {
				t.Errorf("CommitToString() = %v, want %v", gotCommitString, tt.wantCommitString)
			}*/
		})
	}
}

func TestOpenByString(t *testing.T) {
	GString := "364732713438655479616d68564b713549766652396c744a7872344e7276545f2d5f5f736f514f79496d38"
	HString := "49756256713151674b63304144486738416a50376c48756b4a7268566d6b5551574b6a626e6e456a62796b"
	rString := "398891056723960618120827012073372943324213998153307115797995179191814376821"
	commString := "7841426c71536e32476b77626a4a4f5836645078586175684f67546a50326650597a4c31695338434e7934"
	commString2 := "7841426c71536e32476b77626a4a4f5836645078586175684f67546a50326650597a4c31695338434e7934"
	commString3 := "22134sfdaaaa12"

	fmt.Println(OpenByString(commString, GString, HString, rString, []byte("this is the secret message to commit")))
	fmt.Println(OpenByString(commString2, GString, HString, rString, []byte("test ")))
	fmt.Println(OpenByString(commString3, GString, HString, rString, []byte("test ")))
}
