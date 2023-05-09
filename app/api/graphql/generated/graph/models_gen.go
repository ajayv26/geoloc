// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"fmt"
	"geoloc/app/models"
	"io"
	"strconv"

	"github.com/volatiletech/null"
)

type BatchActionInput struct {
	ID       *null.Int64  `json:"id"`
	Str      *null.String `json:"str"`
	No       *null.Int64  `json:"no"`
	DateTime *null.Time   `json:"dateTime"`
	Bool     *null.Bool   `json:"bool"`
}

type PageInfo struct {
	StartCursor int64 `json:"startCursor"`
	EndCursor   int64 `json:"endCursor"`
}

type RequestToken struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SearchFilter struct {
	Search  *null.String  `json:"search"`
	Filter  *FilterOption `json:"filter"`
	SortBy  *SortByOption `json:"sortBy"`
	SortDir *SortDir      `json:"sortDir"`
	Offset  *null.Int     `json:"offset"`
	Limit   *null.Int     `json:"limit"`
}

type UpdateUser struct {
	FirstName *null.String `json:"firstName"`
	LastName  *null.String `json:"lastName"`
	Email     *null.String `json:"email"`
	Phone     *null.String `json:"phone"`
}

type UserResult struct {
	Users []models.User `json:"users"`
	Total int           `json:"total"`
}

type Action string

const (
	ActionArchive   Action = "Archive"
	ActionUnarchive Action = "Unarchive"
)

var AllAction = []Action{
	ActionArchive,
	ActionUnarchive,
}

func (e Action) IsValid() bool {
	switch e {
	case ActionArchive, ActionUnarchive:
		return true
	}
	return false
}

func (e Action) String() string {
	return string(e)
}

func (e *Action) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Action(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Action", str)
	}
	return nil
}

func (e Action) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilterOption string

const (
	FilterOptionAll      FilterOption = "All"
	FilterOptionActive   FilterOption = "Active"
	FilterOptionDraft    FilterOption = "Draft"
	FilterOptionAccepted FilterOption = "Accepted"
	FilterOptionArchived FilterOption = "Archived"
)

var AllFilterOption = []FilterOption{
	FilterOptionAll,
	FilterOptionActive,
	FilterOptionDraft,
	FilterOptionAccepted,
	FilterOptionArchived,
}

func (e FilterOption) IsValid() bool {
	switch e {
	case FilterOptionAll, FilterOptionActive, FilterOptionDraft, FilterOptionAccepted, FilterOptionArchived:
		return true
	}
	return false
}

func (e FilterOption) String() string {
	return string(e)
}

func (e *FilterOption) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterOption(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterOption", str)
	}
	return nil
}

func (e FilterOption) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortByOption string

const (
	SortByOptionDateCreated  SortByOption = "DateCreated"
	SortByOptionDateUpdated  SortByOption = "DateUpdated"
	SortByOptionAlphabetical SortByOption = "Alphabetical"
)

var AllSortByOption = []SortByOption{
	SortByOptionDateCreated,
	SortByOptionDateUpdated,
	SortByOptionAlphabetical,
}

func (e SortByOption) IsValid() bool {
	switch e {
	case SortByOptionDateCreated, SortByOptionDateUpdated, SortByOptionAlphabetical:
		return true
	}
	return false
}

func (e SortByOption) String() string {
	return string(e)
}

func (e *SortByOption) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortByOption(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortByOption", str)
	}
	return nil
}

func (e SortByOption) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortDir string

const (
	SortDirAscending  SortDir = "Ascending"
	SortDirDescending SortDir = "Descending"
)

var AllSortDir = []SortDir{
	SortDirAscending,
	SortDirDescending,
}

func (e SortDir) IsValid() bool {
	switch e {
	case SortDirAscending, SortDirDescending:
		return true
	}
	return false
}

func (e SortDir) String() string {
	return string(e)
}

func (e *SortDir) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDir(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDir", str)
	}
	return nil
}

func (e SortDir) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ViewOption string

const (
	ViewOptionAdmin ViewOption = "ADMIN"
)

var AllViewOption = []ViewOption{
	ViewOptionAdmin,
}

func (e ViewOption) IsValid() bool {
	switch e {
	case ViewOptionAdmin:
		return true
	}
	return false
}

func (e ViewOption) String() string {
	return string(e)
}

func (e *ViewOption) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ViewOption(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ViewOption", str)
	}
	return nil
}

func (e ViewOption) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}