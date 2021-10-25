package dh

import "testing"

func Test_Dh(t *testing.T) {

	client, err := GenerateKey()
	if err != nil {
		t.Error(err)
		return
	}

	server, err := GenerateKey()
	if err != nil {
		t.Error(err)
		return
	}

	key1, err := ComputeKey(client.PubKey, server.PriKey)
	if err != nil {
		t.Error(err)
		return
	}

	key2, err := ComputeKey(server.PubKey, client.PriKey)
	if err != nil {
		t.Error(err)
		return
	}

	if key1.Cmp(key2) == 0 {
		t.Log("succeed.")
	} else {
		t.Error("failed.")
	}

}
