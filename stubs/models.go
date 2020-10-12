package stubs

// https://gohugo.io/variables/

type PageInfo struct {
	Aliases         []string     //aliases of this page
	Content         string       //the content itself, defined below the front matter.
	Data            SiteData     // the data specific to this type of page.
	Date            Time         //the date associated with the page
	Description     string       //the description for the page.
	Dir             string       //the path of the folder containing this content file. The path is relative to the content folder.
	Draft           bool         //a boolean, true if the content is marked as a draft in the front matter.
	ExpiryDate      Time         //the date on which the content is scheduled to expire
	File            FileInfo     //filesystem-related data for this content file
	FuzzyWordCount  int          //the approximate number of words in the content.
	Hugo            HugoInfo     //hugo Variables.
	IsHome          bool         //true in the context of the homepage.
	IsNode          bool         //always false for regular content pages.
	IsPage          bool         //always true for regular content pages.
	IsSection       bool         //true if .Kind is section
	IsTranslated    bool         //true if there are translations to display.
	Keywords        []string     //the meta keywords for the content.
	Kind            string       //the page’s kind. Possible return values are page, home, section, taxonomy, or taxonomyTerm. Note that there are also RSS, sitemap, robotsTXT, and 404 kinds, but these are only available during the rendering of each of these respective page’s kind and therefore not available in any of the Pages collections.
	Language        LanguageInfo //a language object that points to the language’s definition in the site config. .Language.Lang gives you the language code.
	Lastmod         Time         //the date the content was last modified
	LinkTitle       string       //access when creating links to the content. If set, Hugo will use the linktitle from the front matter before title.
	Next            *PageInfo    //Points up to the next regular page
	NextInSection   *PageInfo    //Points up to the next regular page below the same top level section
	OutputFormats   []string     //contains all formats, including the current format, for a given page
	Pages           []*PageInfo  //a collection of associated pages
	Permalink       string       //the Permanent link for this page
	Plain           string       //the Page content stripped of HTML tags and presented as a string.
	PlainWords      []string     //the Page content stripped of HTML as a []string using Go’s strings.Fields to split .Plain into a slice.the Page content stripped of HTML as a []string using Go’s strings.Fields to split .Plain into a slice.
	Prev            *PageInfo    //Points down to the previous regular page
	PrevInSection   *PageInfo    //Points down to the previous regular page below the same top level section
	PublishDate     Time         //the date on which the content was or will be published
	RSSLink         string       //link to the page’s RSS feed. This is deprecated
	RawContent      string       //raw markdown content without the front matter
	ReadingTime     string       //the estimated time, in minutes, it takes to read the content.
	Resources       []string     //resources such as images and CSS that are associated with this page
	Ref             string       //returns the permalink for a given reference (e.g., .Ref "sample.md")
	RelPermalink    string       //the relative permanent link for this page.
	RelRef          string       //returns the relative permalink for a given reference (e.g., RelRef "sample.md")
	Site            SiteInfo     //site Variables
	Sites           SitesInfo    //returns all sites (languages). A typical use case would be to link back to the main language:
	Summary         string       //a generated summary of the content for easily showing a snippet in a summary view.
	TableOfContents string       //the rendered table of contents for the page
	Title           string       //the title for this page.
	Translations    []string     //a list of translated versions of the current page. See Multilingual Mode for more information.
	TranslationKey  string       //the key used to map language translations of the current page
	Truncated       bool         //a boolean, true if the .Summary is truncated. Useful for showing a “Read more…” link only when necessary
	Type            string       //the content type of the content (e.g., posts)
	UniqueID        string       //the MD5-checksum of the content file’s path. This variable is deprecated and will be removed, use .File.UniqueID instead.
	Weight          int          //assigned weight (in the front matter) to this content, used in sorting
	WordCount       int          //the number of words in the content.
	GitInfo         GitInfo      // git info
	Params          PageParams   //Page-level params
	Paginator       *Paginator   //paginator
}

// The view is an alternative layout and should be a file name that points to a template in one of the locations specified in the documentation for Content Views.
func (r *PageInfo) Render(layout string) string {
	return ""
}

// .RenderString is a method on Page that renders some markup to HTML using the content renderer defined for that page (if not set in the options).
func (r *PageInfo) RenderString(markup string) string {
	return ""
}

//.GetPage returns a page of a given path. Both Site and Page implements this method
func (r *PageInfo) GetPage(path string) PageInfo {
	return PageInfo{}
}

//Acts as a “scratchpad” to allow for writable page- or shortcode-scoped variables.
func (r *PageInfo) Scratch() Scratch {
	return Scratch{}
}

//.HasMenuCurrent is a method in Page object returning a boolean value.
func (r *PageInfo) HasMenuCurrent(menu string, menuEntry interface{}) bool {
	return true
}

//.IsMenuCurrent is a method in Page object returning a boolean value
func (r *PageInfo) IsMenuCurrent(menu string, menuEntry interface{}) bool {
	return true
}

type SitesInfo struct {
	First SiteInfo //returns the site for the first language. If this is not a multilingual setup, it will return itself.
}

type SiteInfo struct {
	AllPages        []PageInfo             //array of all pages, regardless of their translation
	Author          map[string]interface{} //a map of the authors as defined in the site configuration.
	BaseURL         string                 //the base URL for the site as defined in the site configuration
	BuildDrafts     bool                   // a boolean (default: false) to indicate whether to build drafts as defined in the site configuration boolean (default: false) to indicate whether to build drafts as defined in the site configuration
	Copyright       string                 //a string representing the copyright of your website as defined in the site configuration.
	Data            SiteData               //custom data
	DisqusShortname string                 //a string representing the shortname of the Disqus shortcode as defined in the site configuration.
	GoogleAnalytics string                 //a string representing your tracking code for Google Analytics as defined in the site configuration.
	Home            *PageInfo              //reference to the homepage’s page object
	IsMultiLingual  bool                   //whether there are more than one language in this site.
	IsServer        bool                   //a boolean to indicate if the site is being served with Hugo’s built-in server
	Language        LanguageInfo           //indicates the language currently being used to render the website
	LanguageCode    string                 //a string representing the language as defined in the site configuration. This is mostly used to populate the RSS feeds with the right language code.
	LanguagePrefix  string                 //this can be used to prefix URLs to point to the correct language
	Languages       []LanguageInfo         //an ordered list (ordered by defined weight) of languages
	LastChange      string                 //a string representing the date/time of the most recent change to your site
	Menus           []MenuEntry            //all of the menus in the site.
	Pages           []PageInfo             //array of all content ordered by Date with the newest first
	RegularPages    []PageInfo             //a shortcut to the regular page collection
	Sections        []string               //top-level directories of the site.
	Title           string                 //a string representing the title of the site
	Taxonomies      map[string]string      //taxonomies
	Params          SiteParams             //a container holding the values from the params section of your site configuration.
}

//.GetPage returns a page of a given path. Both Site and Page implements this method
func (r *SiteInfo) GetPage(path string) PageInfo {
	return PageInfo{}
}

type Paginator struct {
	PageNumber            int        //The current page’s number in the pager sequence
	URL                   string     //The relative URL to the current pager
	Pages                 []PageInfo //The pages in the current pager
	NumberOfElements      int        //The number of elements on this page
	HasPrev               bool       //Whether there are page(s) before the current
	Prev                  PageInfo   //The pager for the previous page
	HasNext               bool       //Whether there are page(s) after the current
	Next                  PageInfo   //The pager for the next page
	First                 PageInfo   //The pager for the first page
	Last                  PageInfo   //The pager for the last page
	Pagers                []PageInfo
	TotalPages            int //The number of pages in the paginator
	TotalNumberOfElements int //The number of elements on all pages in this paginator
}

type LanguageInfo struct {
	Lang         string // the language code of the current locale (e.g., en).
	LanguageName string //the full language name (e.g. English).
	Weight       string //the weight that defines the order in the .Site.Languages list.
}

type FileInfo struct {
	Path            string //the original relative path of the page, relative to the content dir (e.g., posts/foo.en.md)
	LogicalName     string //the name of the content file that represents a page (e.g., foo.en.md)
	ContentBaseName string //the filename without extension or optional language identifier (e.g., foo)
	BaseFileName    string //is a either TranslationBaseName or name of containing folder if file is a leaf bundle.
	Ext             string //the file extension of the content file (e.g., md)
	Lang            string //the language associated with the given file if Hugo’s Multilingual features are enabled (e.g., en)
	Dir             string //given the path content/posts/dir1/dir2/, the relative directory path of the content file will be returned (e.g., posts/dir1/dir2/)
	UniqueID        string //the MD5-checksum of the content file’s path.
}

type MenuEntry struct {
	Menu       string      //string Name of the menu that contains this menu entry.
	URL        string      //string URL that the menu entry points to
	Page       PageInfo    //Page Reference to the page object associated with the menu entry
	Name       string      //string Name of the menu entry
	Identifier string      //string Value of the identifier key if set for the menu entry
	Pre        string      //template.HTML Value of the pre key if set for the menu entry. This value typically contains a string representing HTML.
	Post       string      //template.HTML Value of the post key if set for the menu entry. This value typically contains a string representing HTML.
	Weight     int         //int Value of the weight key if set for the menu entry
	Parent     string      //string Name (or Identifier if present) of this menu entry’s parent menu entry
	Children   []MenuEntry //It is a collection of children menu entries, if any, under the current menu entry
	//functions
	HasChildren bool   //boolean Returns true if .Children is non-nil.
	KeyName     string //string Returns the .Identifier if present, else returns the .Name.
	Title       string //string Link title, meant to be used in the title attribute of a menu entry’s <a>-tags
}

type HugoInfo struct {
	Generator   string //<meta> tag for the version of Hugo that generated the site. .Hugo.Generator outputs a complete HTML tag; e.g. <meta name="generator" content="Hugo 0.18" />
	Version     string //the current version of the Hugo binary you are using e.g. 0.13-DEV
	Environment string //the current running environment as defined through the --environment cli tag.
	BuildDate   string //the compile date of the current Hugo binary formatted with RFC 3339 e.g. 2002-10-02T10:00:00-05:00
	CommitHash  string //the git commit hash of the current Hugo binary e.g. 0e8bed9ccffba0df554728b46c5bbf6d78ae5247
}

type GitInfo struct {
	AbbreviatedHash string //the abbreviated commit hash (e.g., 866cbcc)
	AuthorName      string //the author’s name, respecting .mailmap
	AuthorEmail     string //the author’s email address, respecting .mailmap
	AuthorDate      Time   //the author date
	Hash            string //the commit hash (e.g., 866cbccdab588b9908887ffd3b4f2667e94090c3)
	Subject         string //commit message subject (e.g., tpl: Add custom index function)
}

type Time struct {
}

// format {{ .Date.Format "Jan 2nd 2006"}}
func (r *Time) Format(pattern string) string {
	return ""
}

//.Unix returns the local Time corresponding to the given Unix time, sec seconds and nsec nanoseconds since January 1, 1970 UTC.
func (r *Time) Unix() int {
	return 0
}

type Scratch struct {
}

// Set the given value to a given key
func (r *Scratch) Set(key string, value string) {
}

// Get the value of a given key
func (r *Scratch) Get(key string) interface{} {
	return ""
}

// Will add a given value to existing value of the given key.
// For single values, Add accepts values that support Go’s + operator. If the first Add for a key is an array or slice, the following adds will be appended to that list.
func (r *Scratch) Add(key string, value string) {
}

// Takes a key, mapKey and value and add a map of mapKey and value to the given key.
func (r *Scratch) SetInMap(key string, mapKey string, value string) {
}

// Returns array of values from key sorted by mapKey
func (r *Scratch) GetSortedMapValues(key string) []string {
	return []string{}
}

// Removes the given key
func (r *Scratch) Delete(key string) {
}

// ========== customized section ============

//customized page params
type PageParams struct {
	tags       []string
	categories []string
	imgs       []string
	justify    bool
}

//customized site params
type SiteParams struct {
	logo    string
	slogan  string
	license string
	style   string
	syntax  Syntax
	nav     Nav
	math    Math
}

// customized data structs
type SiteData struct {
	//custom defined data struct
	items Items
}

type Items struct {
	servicecall []Item
	traffic     []Item
	security    []Item
	integration []Item
}

type Item struct {
	title string
	thumb string
	url   string
}

type Syntax struct {
	use       string
	theme     string
	darkTheme string
	webFonts  bool
}

type Nav struct {
	showCategories bool
	showTags       bool
}

type Math struct {
	enable bool
	use    string
}
