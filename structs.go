package sirv

import "net/http"

type PlanLimit struct {
	TotalRequestsPerHour   int
	SearchRequests         int
	VideoToSpinConversions int
	SpinToVideoConversions int
	FetchRequests          int
}

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
	Limit      PlanLimit
}

type TokenResponse struct {
	Token   string   `json:"token"`
	Expires int      `json:"expiresIn"`
	Scope   []string `json:"scope"`
}

type AuthPayload struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type AccountInfo struct {
	DateCreated   string `json:"dateCreated"`
	Alias         string `json:"alias"`
	FileSizeLimit int64  `json:"fileSizeLimit"`
	Fetching      struct {
		Enabled bool   `json:"enabled"`
		Type    string `json:"type"`
		HTTP    struct {
			Auth struct {
				Enabled bool `json:"enabled"`
			} `json:"auth"`
			URL string `json:"url"`
		} `json:"http"`
		MaxFileSize int64 `json:"maxFilesize"`
	} `json:"fetching"`
	Minify struct {
		Enabled bool `json:"enabled"`
	} `json:"minify"`
	CDNTempURL string `json:"cdnTempURL"`
	CDNURL     string `json:"cdnURL"`
	Aliases    map[string]struct {
		Prefix string `json:"prefix"`
		CDN    bool   `json:"cdn"`
	} `json:"aliases"`
}

type LimitInfo struct {
	Count     int `json:"count"`
	Limit     int `json:"limit"`
	Remaining int `json:"remaining"`
	Reset     int `json:"reset"`
}

type APILimits struct {
	S3Global                  LimitInfo `json:"s3:global"`
	S3PUT                     LimitInfo `json:"s3:PUT"`
	S3GET                     LimitInfo `json:"s3:GET"`
	S3DELETE                  LimitInfo `json:"s3:DELETE"`
	RestGlobal                LimitInfo `json:"rest:global"`
	RestPostFilesSearch       LimitInfo `json:"rest:post:files:search"`
	RestPostFilesSearchScroll LimitInfo `json:"rest:post:files:search:scroll"`
	RestPostFilesVideo2spin   LimitInfo `json:"rest:post:files:video2spin"`
	RestPostFilesSpin2video   LimitInfo `json:"rest:post:files:spin2video"`
	RestPostFilesFetch        LimitInfo `json:"rest:post:files:fetch"`
	RestPostFilesUpload       LimitInfo `json:"rest:post:files:upload"`
	RestPostFilesDelete       LimitInfo `json:"rest:post:files:delete"`
	RestPostAccount           LimitInfo `json:"rest:post:account"`
	RestPostAccountFetching   LimitInfo `json:"rest:post:account:fetching"`
	RestGetStatsHttp          LimitInfo `json:"rest:get:stats:http"`
	RestGetStatsStorage       LimitInfo `json:"rest:get:stats:storage"`
	RestPostAccountNew        LimitInfo `json:"rest:post:account:new"`
	RestPostUserAccounts      LimitInfo `json:"rest:post:user:accounts"`
	RestGetRestCredentials    LimitInfo `json:"rest:get:rest:credentials"`
	RestPostVideoToSpin       LimitInfo `json:"rest:post:video:toSpin"`
	RestPostUploadToSirv      LimitInfo `json:"rest:post:upload:toSirv"`
	FtpGlobal                 LimitInfo `json:"ftp:global"`
	FtpSTOR                   LimitInfo `json:"ftp:STOR"`
	FtpRETR                   LimitInfo `json:"ftp:RETR"`
	FtpDELE                   LimitInfo `json:"ftp:DELE"`
	FetchFile                 LimitInfo `json:"fetch:file"`
}

type StorageInfo struct {
	Plan              int64       `json:"plan"`
	Burstable         int64       `json:"burstable"`
	Extra             int64       `json:"extra"`
	Used              int64       `json:"used"`
	Files             int64       `json:"files"`
	QuotaExceededDate interface{} `json:"quotaExceededDate"` // The type here is uncertain, if it's a timestamp string you can use time.Time
}

type User struct {
	Role   string `json:"role"`
	UserID string `json:"userId"`
}

type Filename string

type FileSearchPayload struct {
	Query  string            `json:"query,omitempty"`
	Sort   map[string]string `json:"sort,omitempty"`
	From   int               `json:"from,omitempty"`
	Size   int               `json:"size,omitempty"`
	Scroll bool              `json:"scroll,omitempty"`
}

type FileSearchResponse struct {
	Hits     []FileHit `json:"hits"`
	Total    int       `json:"total"`
	Relation string    `json:"_relation"`
	ScrollId string    `json:"scrollId,omitempty"`
}

type FileHit struct {
	Index   string    `json:"_index"`
	Type    string    `json:"_type"`
	Id      string    `json:"_id"`
	Routing string    `json:"_routing"`
	Source  Source    `json:"_source"`
	Sort    []float64 `json:"sort"`
}

type FileMeta struct {
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Format   string `json:"format"`
	Duration int    `json:"duration"`
	EXIF     EXIF   `json:"EXIF"`
}

type EXIF struct {
	ModifyDate string `json:"ModifyDate"`
}

type FileSearchScrollPayload struct {
	ScrollId string `json:"scrollId"`
}

type FileSearchScrollResponse struct {
	Hits     []FileHit `json:"hits"`
	Total    int       `json:"total"`
	Relation string    `json:"_relation"`
	ScrollId string    `json:"scrollId"`
}

type FolderContents struct {
	Contents     []File `json:"contents"`
	Continuation string `json:"continuation"`
}

type CommonFileInfo struct {
	MTime       string                 `json:"mtime"`
	ContentType string                 `json:"contentType"`
	Size        int                    `json:"size"`
	IsDirectory bool                   `json:"isDirectory"`
	Meta        map[string]interface{} `json:"meta"`
}

type File struct {
	CommonFileInfo
	Filename Filename `json:"filename"`
}

type FileInfo struct {
	CommonFileInfo
	CTime string `json:"ctime"`
}

type Source struct {
	CommonFileInfo
	AccountId string   `json:"accountId"`
	Filename  Filename `json:"filename"`
	Dirname   string   `json:"dirname"`
	Basename  string   `json:"basename"`
	Extension string   `json:"extension"`
	Id        string   `json:"id"`
}

type FileDeletionResponse struct {
	HTTPStatusCode int `json:"httpStatusCode"`
}

type DeleteFilePayload struct {
	Filename Filename `json:"filename"`
}
