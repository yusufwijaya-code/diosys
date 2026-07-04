package gdrive_helper

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Config holds the Google Drive OAuth credentials and the destination folder.
type Config struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
	RootFolderID string
}

// Client wraps the Google Drive service together with the configured root folder.
type Client struct {
	service      *drive.Service
	rootFolderID string
}

// UploadResult mirrors the ldksyahid-app contract returned after an upload.
type UploadResult struct {
	FileName string `json:"fileName"`
	GdriveID string `json:"gdriveID"`
}

// NewClient builds a Drive client from a refresh-token based OAuth credential.
func NewClient(cfg Config) (*Client, error) {
	if cfg.ClientID == "" || cfg.ClientSecret == "" || cfg.RefreshToken == "" {
		return nil, fmt.Errorf("incomplete google drive credentials")
	}

	oauthConfig := &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{drive.DriveScope},
	}

	ctx := context.Background()
	tokenSource := oauthConfig.TokenSource(ctx, &oauth2.Token{RefreshToken: cfg.RefreshToken})

	service, err := drive.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, fmt.Errorf("failed to create google drive service: %w", err)
	}

	return &Client{service: service, rootFolderID: cfg.RootFolderID}, nil
}

// UploadImage stores an uploaded file in the root folder and makes it publicly
// readable, returning the stored file name and the Google Drive file ID.
func (c *Client) UploadImage(fileHeader *multipart.FileHeader, prefix string) (UploadResult, error) {
	src, err := fileHeader.Open()
	if err != nil {
		return UploadResult{}, fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	extension := filepath.Ext(fileHeader.Filename)
	fileName := fmt.Sprintf("%s-%s%s", prefix, uuid.NewString(), extension)

	file := &drive.File{Name: fileName}
	if c.rootFolderID != "" {
		file.Parents = []string{c.rootFolderID}
	}

	created, err := c.service.Files.Create(file).Media(src).Fields("id", "name").Do()
	if err != nil {
		return UploadResult{}, fmt.Errorf("failed to upload file to google drive: %w", err)
	}

	permission := &drive.Permission{Role: "reader", Type: "anyone"}
	if _, err := c.service.Permissions.Create(created.Id, permission).Do(); err != nil {
		return UploadResult{}, fmt.Errorf("failed to set public permission: %w", err)
	}

	return UploadResult{FileName: created.Name, GdriveID: created.Id}, nil
}

// DeleteFile removes a file from Google Drive. A missing file is treated as success.
func (c *Client) DeleteFile(gdriveID string) error {
	if gdriveID == "" {
		return nil
	}
	return c.service.Files.Delete(gdriveID).Do()
}

// PublicURL builds a directly embeddable image URL for a Google Drive file ID.
func PublicURL(gdriveID string) string {
	if gdriveID == "" {
		return ""
	}
	return fmt.Sprintf("https://drive.google.com/thumbnail?id=%s&sz=w1000", gdriveID)
}

// DownloadURL builds a direct download URL for any Google Drive file (e.g. PDF).
func DownloadURL(gdriveID string) string {
	if gdriveID == "" {
		return ""
	}
	return fmt.Sprintf("https://drive.google.com/uc?export=download&id=%s", gdriveID)
}
