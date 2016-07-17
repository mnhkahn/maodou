package models

import "time"

type Result struct {
	Id          uint64    `orm:"pk autoincr INT(11)" json:"id"`
	Title       string    `orm:"size(200)" json:"title"`
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)" json:"create_time"`
	Author      string    `orm:"size(45)" json:"author"`
	Detail      string    `orm:"type(text)" json:"detail"`
	Category    string    `orm:"size(45)" json:"category"`
	Tags        string    `orm:"size(45)" json:"tags"`
	Figure      string    `orm:"size(100)" json:"figure"`
	Description string    `orm:"type(text)" json:"description"`
	Link        string    `orm:"size(100)" json:"link"`
	Source      string    `orm:"size(45)" json:"source"`
	ParseDate   time.Time `orm:"auto_now_add;type(datetime)" json:"parse_date"`
}
