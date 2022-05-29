package notion

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

type Block struct {
	Object           string            `json:"object"`
	ID               string            `json:"id"`
	CreatedTime      string            `json:"created_time"`
	CreatedBy        RelatedUser       `json:"created_by"`
	LastEditedTime   string            `json:"last_edited_time"`
	LastEditedBy     RelatedUser       `json:"last_edited_by"`
	HasChildren      bool              `json:"has_children"`
	Type             string            `json:"type"`
	Archived         bool              `json:"archived"`
	Paragraph        *Paragraph        `json:"paragraph,omitempty"`
	HeadingOne       *HeadingOne       `json:"heading_one,omitempty"`
	HeadingTwo       *HeadingTwo       `json:"heading_two,omitempty"`
	HeadingThree     *HeadingThree     `json:"heading_three,omitempty"`
	Callout          *Callout          `json:"callout,omitempty"`
	Quote            *Quote            `json:"quoto,omitempty"`
	NumberedListItem *NumberedListItem `json:"numbered_list_item,omitempty"`
	BulletedListItem *BulletedListItem `json:"bulleted_list_item,omitempty"`
	ToDo             *ToDo             `json:"to_do,omitempty"`
	Toggle           *Toggle           `json:"toggle,omitempty"`
	Code             *Code             `json:"code,omitempty"`
	ChildPage        *ChildPage        `json:"child_page,omitempty"`
	ChildDatabase    *ChildDatabase    `json:"child_database,omitempty"`
	Embed            *Embed            `json:"embed,omitempty"`
	Image            *Image            `json:"image,omitempty"`
	Video            *Video            `json:"video,omitempty"`
	File             *File             `json:"file,omitempty"`
	PDF              *PDF              `json:"pdf,omitempty"`
	Bookmark         *Bookmark         `json:"bookmark,omitempty"`
	Equation         *Equation         `json:"equation,omitempty"`
	TableOfContent   *TableOfContent   `json:"table_of_content,omitempty"`
	Breadcrumb       *Breadcrumb       `json:"breadcrumb,omitempty"`
	ColumnList       *ColumnList       `json:"column_list,omitempty"`
	Column           *Column           `json:"column,omitempty"`
	LinkPreview      *LinkPreview      `json:"linked_preview,omitempty"`
	Template         *Template         `json:"template,omitempty"`
	LinkToPage       *LinkToPage       `json:"link_to_page,omitempty"`
	SyncedBlock      *SyncedBlock      `json:"synced_block,omitempty"`
	SyncedForm       *SyncedForm       `json:"synced_form,omitempty"`
	Table            *Table            `json:"table,omitempty"`
	TableRow         *TableRow         `json:"table_row,omitempty"`
}

// block type: https://developers.notion.com/reference/block

type Paragraph struct {
	RichText RichText    `json:"rich_text"`
	Color    string      `json:"color"`
	Children interface{} `json:"children"`
}

type Heading struct {
	RichText RichText `json:"rich_text"`
	Color    string   `json:"color"`
}

type HeadingOne struct {
	Heading
}

type HeadingTwo struct {
	Heading
}

type HeadingThree struct {
	Heading
}

type Callout struct {
	RichText RichText    `json:"rich_text"`
	Icon     string      `json:"icon"`
	Color    string      `json:"color"`
	Children interface{} `json:"children"`
}

type Quote struct {
	RichText RichText `json:"rich_text"`
	Color    string   `json:"color"`
}

type BulletedListItem struct {
	RichText RichText    `json:"rich_text"`
	Color    string      `json:"color"`
	Children interface{} `json:"children"`
}

type NumberedListItem struct {
	RichText RichText    `json:"rich_text"`
	Color    string      `json:"color"`
	Children interface{} `json:"children"`
}

type ToDo struct {
	RichText RichText    `json:"rich_text"`
	Checked  bool        `json:"checked"`
	Color    string      `json:"color"`
	Children interface{} `json:"children"`
}

type Toggle struct {
	RichText RichText    `json:"rich_text"`
	Color    string      `json:"color"`
	Children interface{} `json:"children"`
}

type Code struct {
	RichText RichText `json:"rich_text"`
	Caption  RichText `json:"caption"`
	Language string   `json:"language"`
}

type ChildPage struct {
	Title string `json:"title"`
}

type ChildDatabase struct {
	Title string `json:"title"`
}

type Embed struct {
	URL string `json:"url"`
}

type Image struct {
	FileObject
}

type Video struct {
	FileObject
}

type File struct {
	FileObject
	Caption RichText `json:"caption"`
}

type PDF struct {
	FileObject
}

type Bookmark struct {
	URL     string   `json:"url"`
	Caption RichText `json:"caption"`
}

type Equation struct {
	Expression string `json:"expression"`
}

type TableOfContent struct {
	Color string `json:"color"`
}

type Breadcrumb struct {
}

type ColumnList struct {
	Children interface{} `json:"children"`
}

type Column struct {
	Children interface{} `json:"children"`
}

type LinkPreview struct {
	URL string `json:"url"`
}

type Template struct {
	RichText RichText    `json:"rich_text"`
	Children interface{} `json:"children"`
}

type LinkToPage struct {
	Type       string `json:"type"`
	PageID     string `json:"page_id"`
	DatabaseID string `json:"databse_id"`
}

type SyncedBlock struct {
	SyncedFrom *SyncedForm `json:"synced_from"`
	Children   interface{} `json:"children"`
}

type SyncedForm struct {
	Type    string `json:"type"`
	BlockID string `json:"block_id"`
}

type Table struct {
	TableWidth      int         `json:"table_width"`
	HasColumnHeader bool        `json:"has_column_header"`
	HasRowHeader    bool        `json:"has_row_header"`
	Children        interface{} `json:"children"`
}

type TableRow struct {
	Cells []RichText `json:"cells"`
}

type FileObject struct {
	Type     string `json:"type"`
	External string `json:"external"`
}

type RelatedUser struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type RichText []struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        string      `json:"href"`
}

type Text struct {
	Content string      `json:"content"`
	Link    interface{} `json:"link"`
}

func (c *Client) RetrieveBlock(ctx context.Context, blockID string) (*Block, error) {
	reqURL := *c.URL

	reqURL.Path = path.Join(reqURL.Path, "blocks", blockID)

	url := path.Join("blocks", blockID)

	req, err := c.ConstructReq(ctx, url, http.MethodGet)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Notion-Version", "2021-08-16")
	req = req.WithContext(ctx)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	log.Println(res.StatusCode)

	switch res.StatusCode {
	case http.StatusOK:
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var block *Block
		if err := json.Unmarshal(bodyBytes, &block); err != nil {
			return nil, err
		}

		return block, nil

	default:
		return nil, errors.New("unexpected error")
	}
}
