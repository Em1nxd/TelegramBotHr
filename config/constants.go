package config

const (
	DateTimeFormat  = "2006-01-02 15:04:05"
	DateFormat      = "2006-01-02"
	ConsumerGroupID = "api_gateway"

	MinioExcelBucketName   = "excel"
	MinioCatalogBucketName = "image"

	SystemUserTyperID = "1fe92aa8-2a61-4bf1-b907-182b497584ad"
	AdminUserTypeID   = "9fb3ada6-a73b-4b81-9295-5c1605e54552"

	SystemTypeID = "1fe92aa8-2a61-4bf1-b907-182b497584ad"
	AdminTypeID  = "9fb3ada6-a73b-4b81-9295-5c1605e54552"
)

var (
	ContentTypes = map[string]string{
		"png":  "image",
		"jpg":  "image",
		"jpeg": "image",
		"xls":  "excel",
		"xlsx": "excel",
		"csv":  "excel",
	}

	Languages = []string{"uz", "en", "ru"}
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)
