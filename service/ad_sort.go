package service

import (
	"easygoadmin/constant"
	"easygoadmin/dto"
	"easygoadmin/model"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"easygoadmin/vo"
	"errors"
	"strconv"
	"strings"
	"time"
)

var AdSort = new(adSortService)

type adSortService struct{}

func (s *adSortService) GetList(req dto.AdSortPageReq) ([]vo.AdSortInfoVo, int64, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 广告位名称
	if req.Description != "" {
		query = query.Where("description like ?", "%"+req.Description+"%")
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 对象转换
	var list []model.AdSort
	count, err := query.FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}

	// 数据处理
	var result = make([]vo.AdSortInfoVo, 0)
	for _, v := range list {
		platformName, ok := constant.ADSORT_PLATFORM_LIST[v.Platform]
		item := vo.AdSortInfoVo{}
		item.AdSort = v
		if ok {
			item.PlatformName = platformName
		}
		// 站点名称
		if v.ItemId > 0 {
			info := &model.Item{Id: v.ItemId}
			has, err2 := info.Get()
			if err2 == nil && has {
				item.ItemName = info.Name
			}
		}

		// 栏目名称
		if v.CateId > 0 {
			cateName := ItemCate.GetCateName(v.CateId, ">>")
			item.CateName = cateName
		}

		// 加入数组
		result = append(result, item)
	}
	return result, count, nil
}

func (s *adSortService) Add(req dto.AdSortAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.AdSort
	entity.Description = req.Description
	entity.ItemId = req.ItemId
	entity.CateId = req.CateId
	entity.LocId = req.LocId
	entity.Platform = req.Platform
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	entity.Mark = 1

	// 插入数据
	return entity.Insert()
}

func (s *adSortService) Update(req dto.AdSortUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.AdSort{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置对象
	entity.Description = req.Description
	entity.ItemId = req.ItemId
	entity.CateId = req.CateId
	entity.LocId = req.LocId
	entity.Platform = req.Platform
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()

	// 更新记录
	return entity.Update()
}

func (s *adSortService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.AdSort{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			id, _ := strconv.Atoi(v)
			entity := &model.AdSort{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *adSortService) GetAdSortList() []vo.AdSortInfoVo {
	// 查询广告位列表
	var list []model.AdSort
	utils.XormDb.Where("mark=1").OrderBy("sort asc").Find(&list)
	// 数据处理
	result := make([]vo.AdSortInfoVo, 0)
	for _, v := range list {
		item := vo.AdSortInfoVo{}
		item.AdSort = v
		item.Description = v.Description + " >> " + gconv.String(v.LocId)
		result = append(result, item)
	}
	return result
}
