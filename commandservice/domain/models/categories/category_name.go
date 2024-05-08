package categories

import (
	"commandservice/errs"
	"unicode/utf8"
)

type CategoryName struct {
	value string
}

func (ins *CategoryName) Value() string {
	return ins.value
}

func (ins *CategoryName) Equals(other *CategoryName) bool {
	return ins.value == other.Value()
}

func NewCategoryName(value string) (*CategoryName, *errs.DomainError) {
	const MIN_LENGTH int = 5
	const MAX_LENGTH int = 30

	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(
			"カテゴリ名は5文字以上30文字以下でなければなりません",
		)
	}
	return &CategoryName{value: value}, nil
}
