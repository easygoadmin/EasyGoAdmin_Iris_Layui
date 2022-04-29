package service

import (
	"easygoadmin/dto"
	"easygoadmin/model"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"strconv"
	"strings"
	"time"
)

var MemberLevel = new(memberLevelService)

type memberLevelService struct{}

func (s *memberLevelService) GetList(req dto.MemberLevelPageReq) ([]model.MemberLevel, int64, error) {
	// 创建查询实例
	query := utils.XormDb.Where("mark=1")
	// 等级名称
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	// 排序
	query = query.OrderBy("sort asc")
	// 分页
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 对象转换
	var list []model.MemberLevel
	count, err := query.FindAndCount(&list)
	return list, count, err
}

func (s *memberLevelService) Add(req dto.MemberLevelAddReq, userId int) (int64, error) {
	// 实例化对象
	var entity model.MemberLevel
	entity.Name = req.Name
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now().Unix()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()
	entity.Mark = 1

	// 插入数据
	return entity.Insert()
}

func (s *memberLevelService) Update(req dto.MemberLevelUpdateReq, userId int) (int64, error) {
	// 查询记录
	entity := &model.MemberLevel{Id: req.Id}
	has, err := entity.Get()
	if err != nil || !has {
		return 0, err
	}

	// 设置参数
	entity.Name = req.Name
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now().Unix()

	// 更新记录
	return entity.Update()
}

func (s *memberLevelService) Delete(ids string) (int64, error) {
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &model.MemberLevel{Id: gconv.Int(ids)}
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
			entity := &model.MemberLevel{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}
