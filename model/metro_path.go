package model

type MetroPath struct {
	Id          int64  `orm:"auto"`
	PathName    string `orm:"column(path_name)"`
	FromNode    string `orm:"column(from_node)"`
	ToNode      string `orm:"column(to_node)"`
	DetailCount int    `orm:"column(detail_count)"`
}

func (this *MetroPath) TableName() string {
	return "metro_path"
}

/**
 * 地铁路线详细信息
 */
type MetroPathDetail struct {
	Id            int64  `orm:"auto"`
	PathId        int    `orm:"column(path_id)"`
	FromStation   string `orm:"column(from_node)"`
	ToStation     string `orm:"column(to_node)"`
	Line          string `orm:"column(line)"`
	LineNum       int    `orm:"-"`
	Direction     string `orm:"column(udcode)"`
	PathIndex     int    `orm:"column(path_index)"`
	IsChangePoint int    `orm:"is_change_point"`
}

func (this *MetroPathDetail) TableName() string {
	return "metro_path_detail"
}
