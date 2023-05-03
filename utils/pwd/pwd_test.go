package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$kfFvlCaGClfa6mQECybdCOR/qAij3LtOG.6GoPeidsWm48khwWVSK", "12345"))
	fmt.Println(CheckPwd("$2a$04$kfFvlCaGClfa6mQECybdCOR/qAij3LtOG.6GoPeidsWm48khwWVSK", "123456"))
}
