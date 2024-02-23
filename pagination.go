package pagination

import "gorm.io/gorm"

type Paginator struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type Option struct {
	DB      *gorm.DB
	Page    int
	Limit   int
	ShowSQL bool
}

func Paginate(o *Option, data interface{}) *Paginator {
	db := o.DB

	if o.ShowSQL {
		db = db.Debug()
	}

	o.Page, o.Limit = defaultPageInfo(o.Page, o.Limit)

	paginator := &Paginator{
		Page:  o.Page,
		Limit: o.Limit,
	}

	done := make(chan bool, 1)
	go count(db, data, &paginator.Total, done)
	db.Scopes(Page(o.Page, o.Limit)).Find(data)
	paginator.Data = data
	<-done

	return paginator
}

func Page(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, limit = defaultPageInfo(page, limit)

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func count(db *gorm.DB, model interface{}, total *int64, done chan bool) {
	db.Model(model).Count(total)
	done <- true
}

func defaultPageInfo(page, limit int) (int, int) {
	if page < 1 {
		page = 1
	}

	switch {
	case limit > 100:
		limit = 100
	case limit < 1:
		limit = 10
	}
	return page, limit
}
