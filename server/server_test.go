package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/hermeznetwork/hermez-bridge/bridgetree/pb"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

const port = "8080"

func init() {
	// Change dir to project root
	// This is important because we have relative paths to files containing test vectors
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestBridgeMock(t *testing.T) {
	err := RunMockServer()
	require.NoError(t, err)

	err = healthRestCheck(port)
	require.NoError(t, err)

	address := "http://localhost:" + port

	resp, err := http.Get(address + "/api")
	require.NoError(t, err)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	_ = resp.Body.Close()

	var apiResp pb.CheckAPIResponse
	err = protojson.Unmarshal(bodyBytes, &apiResp)
	require.NoError(t, err)

	require.Equal(t, "v1", apiResp.Api)

	resp, err = http.Get(fmt.Sprintf("%s%s/%s", address, "/bridges", "0xeB17ce701E9D92724AA2ABAdA7E4B28830597Dd9"))
	require.NoError(t, err)

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var bridgeResp pb.GetBridgesResponse
	err = protojson.Unmarshal(bodyBytes, &bridgeResp)
	require.NoError(t, err)
	require.Greater(t, len(bridgeResp.Deposits), 0)

	offset := 3
	resp, err = http.Get(fmt.Sprintf("%s%s/%s?offset=%d", address, "/bridges", "0xeB17ce701E9D92724AA2ABAdA7E4B28830597Dd9", offset))
	require.NoError(t, err)

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	err = protojson.Unmarshal(bodyBytes, &bridgeResp)
	require.NoError(t, err)
	require.Equal(t, len(bridgeResp.Deposits), offset-1)

	offset = 1
	resp, err = http.Get(fmt.Sprintf("%s%s/%s?offset=%d", address, "/claims", "0xeB17ce701E9D92724AA2ABAdA7E4B28830597Dd9", offset))
	require.NoError(t, err)

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var claimResp pb.GetClaimsResponse
	err = protojson.Unmarshal(bodyBytes, &claimResp)
	require.NoError(t, err)
	require.Equal(t, len(claimResp.Claims), 2)

	resp, err = http.Get(fmt.Sprintf("%s%s?orig_net=%d&deposit_cnt=%d", address, "/merkle-proofs", 0, 2))
	require.NoError(t, err)

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var proofResp pb.GetProofResponse
	err = protojson.Unmarshal(bodyBytes, &proofResp)
	require.NoError(t, err)
	require.Equal(t, len(proofResp.Proof.MerkleProof), 32)

	resp, err = http.Get(fmt.Sprintf("%s%s?orig_net=%d&deposit_cnt=%d", address, "/claim-status", 0, 2))
	require.NoError(t, err)

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	var claimStatus pb.GetClaimStatusResponse
	err = protojson.Unmarshal(bodyBytes, &claimStatus)
	require.NoError(t, err)
	require.Equal(t, claimStatus.Ready, true)
}
