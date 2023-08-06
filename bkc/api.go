// Bunkercoin related functionality
package bkc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// ChainInfo defines information about the Blockchain
type ChainInfo struct {
	BlockCount int
	Hashrate   float64
	Difficulty float64
}

// getFromAPI is a helper function that gets a piece of data from the API at apiurl/path
func getFromAPI(apiurl string, path string) ([]byte, error) {
	requrl, err := url.JoinPath(apiurl, path)
	if err != nil {
		log.Println("error while getting'" + path + "'from the API")
		return []byte(""), err
	}

	resp, err := http.Get(requrl)
	if err != nil {
		log.Println("error while getting'" + path + "'from the API")
		return []byte(""), err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("bad HTTP Status code while getting'" + path + "'from the API")
		return []byte(""), fmt.Errorf("bad HTTP status code (%v)", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error while getting'" + path + "'from the API")
		return []byte(""), err
	}
	return data, nil
}

// GetBlockCount returns the number of blocks in the chain using the API at apiurl
func GetBlockCount(apiurl string) (int, error) {
	data, err := getFromAPI(apiurl, "/getblockcount")
	if err != nil {
		log.Println("error while getting the blockcount")
		return 0, err
	}

	blockcount, err := strconv.Atoi(string(data))
	if err != nil {
		log.Println("error while converting the blockcount data")
		return 0, err
	}

	return blockcount, nil
}

// GetHashrate returns the network hashrate of the chain in Hashes/s using the API at apiurl
func GetHashrate(apiurl string) (float64, error) {
	data, err := getFromAPI(apiurl, "/getnetworkhashps")
	if err != nil {
		log.Println("error while getting the network hashrate")
		return 0.0, err
	}

	hashrate, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		log.Println("error while converting the network hashrate data")
		return 0.0, err
	}

	return hashrate, nil
}

// GetDifficulty returns the network difficulty of the chain using the API at apiurl
func GetDifficulty(apiurl string) (float64, error) {
	data, err := getFromAPI(apiurl, "/getdifficulty")
	if err != nil {
		log.Println("error while getting the network difficulty")
		return 0.0, err
	}

	diff, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		log.Println("error while converting the network difficulty data")
		return 0.0, err
	}

	return diff, nil
}

// A combination of all the Get* functions, which returns a ChainInfo struct
func GetChainInfo(apiurl string) (ChainInfo, error) {
	var chaininfo ChainInfo
	var err error

	chaininfo.BlockCount, err = GetBlockCount(apiurl)
	if err != nil {
		return ChainInfo{}, err
	}

	chaininfo.Hashrate, err = GetHashrate(apiurl)
	if err != nil {
		return ChainInfo{}, err
	}

	chaininfo.Difficulty, err = GetDifficulty(apiurl)
	if err != nil {
		return ChainInfo{}, err
	}

	return chaininfo, nil
}
