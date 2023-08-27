package sirv

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func (c *Client) SearchFiles(ctx context.Context, payload FileSearchPayload) (*FileSearchResponse, error) {
	var searchResp FileSearchResponse
	err := c.makeRequest(ctx, http.MethodPost, c.BaseURL+"/files/search", payload, &searchResp)
	if err != nil {
		return nil, err
	}

	return &searchResp, nil
}

func (c *Client) ScrollFilesSearch(ctx context.Context, payload FileSearchScrollPayload) (*FileSearchScrollResponse, error) {
	var scrollResp FileSearchScrollResponse
	err := c.makeRequest(ctx, http.MethodPost, c.BaseURL+"/files/search/scroll", payload, &scrollResp)
	if err != nil {
		return nil, err
	}

	return &scrollResp, nil
}

func (c *Client) ReadFolderContents(ctx context.Context, dirname string, continuation string) (*FolderContents, error) {
	reqUrl := fmt.Sprintf("%s/files/readdir?dirname=%s", c.BaseURL, url.PathEscape(dirname))

	if continuation != "" {
		reqUrl = reqUrl + "&continuation=" + url.PathEscape(continuation)
	}

	var contents FolderContents
	err := c.makeRequest(ctx, http.MethodGet, reqUrl, nil, &contents)
	if err != nil {
		return nil, err
	}

	return &contents, nil
}

func (c *Client) GetFileInfo(ctx context.Context, filename string) (*FileInfo, error) {
	reqUrl := fmt.Sprintf("%s/files/stat?filename=%s", c.BaseURL, url.PathEscape(filename))

	var fileInfo FileInfo
	err := c.makeRequest(ctx, http.MethodGet, reqUrl, nil, &fileInfo)
	if err != nil {
		return nil, err
	}

	return &fileInfo, nil
}

func (c *Client) DownloadFile(ctx context.Context, filename Filename, destPath string) error {
	reqUrl := fmt.Sprintf("%s/files/download?filename=%s", c.BaseURL, url.PathEscape(string(filename)))

	resp, err := c.makeHTTPRequest(ctx, http.MethodGet, reqUrl, nil, "application/json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func (c *Client) DeleteFileOrEmptyFolder(ctx context.Context, filename Filename) error {
	reqUrl := fmt.Sprintf("%s/files/delete?filename=%s", c.BaseURL, url.PathEscape(string(filename)))

	resp, err := c.makeHTTPRequest(ctx, http.MethodPost, reqUrl, nil, "application/json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *Client) UploadFile(ctx context.Context, dstFilename, localPath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	reqUrl := fmt.Sprintf("%s/files/upload?filename=%s", c.BaseURL, url.PathEscape(dstFilename))
	_, err = c.makeHTTPRequest(ctx, http.MethodPost, reqUrl, file, "")
	return err
}

func (c *Client) CreateEmptyFolder(ctx context.Context, dirname string) error {
	reqUrl := fmt.Sprintf("%s/files/mkdir?dirname=%s", c.BaseURL, url.PathEscape(dirname))

	resp, err := c.makeHTTPRequest(ctx, http.MethodPost, reqUrl, nil, "application/json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d. Response: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

func (c *Client) RenameFileOrFolder(ctx context.Context, from, to string) error {
	reqUrl := fmt.Sprintf("%s/files/rename?from=%s&to=%s", c.BaseURL, url.PathEscape(from), url.PathEscape(to))

	resp, err := c.makeHTTPRequest(ctx, http.MethodPost, reqUrl, nil, "application/json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
