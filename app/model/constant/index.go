package constant

const (
	INDEX_PRODUCT_DISPLAY = "product-display"
	INDEX_SUGGESTIONS     = "search-suggestions"
	INDEX_AUTHOR_PREVIEW  = "author-preview"
)

var (
	FILTER_PRODUCT_DISPLAY_DEFAULT = []string{
		"product_meta_is_active=true",
		"product_is_active=true",
		"category_is_active=true",
		"warehouse_is_active=true",
		"(author_is_active=true AND vendor_is_active=true) OR (author_slug IS EMPTY AND vendor_is_active=true)",
	}
)
