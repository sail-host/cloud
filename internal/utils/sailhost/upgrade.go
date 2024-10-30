package sailhost

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

const (
	UpgradeUrl  = "https://api.github.com/repos/sail-host/cloud/tags"
	DownloadUrl = "https://github.com/sail-host/cloud/releases/download/%s/sailhost-%s-%s.tar.gz"
)

func LastVersion() (string, error) {

	type response struct {
		Name       string `json:"name"`
		TarballUrl string `json:"tarball_url"`
		Commit     struct {
			Sha string `json:"sha"`
			Url string `json:"url"`
		} `json:"commit"`
		NodeId string `json:"node_id"`
	}

	var resp []response

	res, err := http.Get(UpgradeUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}

	return resp[0].Name, nil
}

func DownloadLastVersion(downloadPath string) error {
	lastVersion, err := LastVersion()
	if err != nil {
		return err
	}

	url := fmt.Sprintf(DownloadUrl, lastVersion, runtime.GOOS, runtime.GOARCH)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(fmt.Sprintf("%s/bin", downloadPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/bin/sailhost.tar.gz", downloadPath))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		return err
	}

	// Extract tar.gz file and save binary
	tarFile, err := os.Open(fmt.Sprintf("%s/bin/sailhost.tar.gz", downloadPath))
	if err != nil {
		return err
	}
	defer tarFile.Close()

	gzr, err := gzip.NewReader(tarFile)
	if err != nil {
		return err
	}
	defer gzr.Close()
	tr := tar.NewReader(gzr)

	header, err := tr.Next()
	if err != nil {
		return err
	}

	outFile, err := os.OpenFile(fmt.Sprintf("%s/bin/sailhost", downloadPath), os.O_CREATE|os.O_WRONLY, header.FileInfo().Mode())
	if err != nil {
		return err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, tr); err != nil {
		return err
	}

	// Remove tar.gz file
	err = os.Remove(fmt.Sprintf("%s/bin/sailhost.tar.gz", downloadPath))
	if err != nil {
		return err
	}

	// Set permissions
	err = os.Chmod(fmt.Sprintf("%s/bin/sailhost", downloadPath), 0755)
	if err != nil {
		return err
	}

	return nil
}
